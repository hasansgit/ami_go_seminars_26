package censor

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestWrite(t *testing.T) {
	testCases := []struct {
		name   string
		banned []string
		input  []byte
		want   string
	}{
		{
			name:   "одно вхождение",
			banned: []string{"spam"},
			input:  []byte("buy spam now"),
			want:   "buy **** now",
		},
		{
			name:   "несколько вхождений одного слова",
			banned: []string{"aba"},
			input:  []byte("aba aba"),
			want:   "*** ***",
		},
		{
			name:   "несколько запрещённых слов",
			banned: []string{"cat", "dog"},
			input:  []byte("cat and dog"),
			want:   "*** and ***",
		},
		{
			name:   "слово внутри слова тоже заменяется",
			banned: []string{"cat"},
			input:  []byte("concatenate"),
			want:   "con***enate",
		},
		{
			name:   "регистр учитывается",
			banned: []string{"spam"},
			input:  []byte("Spam spam SPAM"),
			want:   "Spam **** SPAM",
		},
		{
			name:   "текст без запрещённых слов не меняется",
			banned: []string{"xxx"},
			input:  []byte("clean text"),
			want:   "clean text",
		},
		{
			name:   "замена идёт слева направо без перекрытий",
			banned: []string{"aa"},
			input:  []byte("aaaaa"),
			want:   "****a",
		},
		{
			name:   "пустые слова в banned игнорируются",
			banned: []string{"", "ab"},
			input:  []byte("ab ab"),
			want:   "** **",
		},
		{
			name:   "nil-список banned пропускает всё",
			banned: nil,
			input:  []byte("anything goes"),
			want:   "anything goes",
		},
		{
			name:   "пустой вход",
			banned: []string{"spam"},
			input:  []byte{},
			want:   "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			w := New(&buf, tc.banned)

			n, err := w.Write(tc.input)
			if err != nil {
				t.Fatalf("Write(%q) вернул ошибку %v, ожидалась nil", tc.input, err)
			}
			if n != len(tc.input) {
				t.Errorf("Write(%q) = %d байт, ожидалось %d (len входа)",
					tc.input, n, len(tc.input))
			}
			if got := buf.String(); got != tc.want {
				t.Errorf("после Write(%q) в writer ушло %q, ожидалось %q",
					tc.input, got, tc.want)
			}
		})
	}
}

// Writer обязан удовлетворять io.Writer — иначе его нельзя передать
// в Fprint, io.Copy и любую другую функцию стандартной библиотеки.
func TestWriterSatisfiesIoWriter(t *testing.T) {
	var buf bytes.Buffer
	var _ io.Writer = New(&buf, nil)
}

// Стандартная библиотека должна уметь писать в Writer напрямую.
func TestWriterWorksWithFmtFprint(t *testing.T) {
	var buf bytes.Buffer
	w := New(&buf, []string{"secret"})

	_, err := io.WriteString(w, "this is secret data")
	if err != nil {
		t.Fatalf("io.WriteString вернул ошибку %v, ожидалась nil", err)
	}

	if got, want := buf.String(), "this is ****** data"; got != want {
		t.Errorf("после io.WriteString в writer ушло %q, ожидалось %q", got, want)
	}
}

// Замена сохраняет длину: каждое вхождение заменяется на равное число
// звёздочек. Это свойство проверяется на цепочке записей.
func TestOutputLengthEqualsInputLength(t *testing.T) {
	var buf bytes.Buffer
	w := New(&buf, []string{"ban"})

	chunks := []string{"ban", "ana", "banban", "?"}
	total := 0
	for _, chunk := range chunks {
		n, err := w.Write([]byte(chunk))
		if err != nil {
			t.Fatalf("Write(%q) вернул ошибку %v, ожидалась nil", chunk, err)
		}
		total += n
	}

	if buf.Len() != total {
		t.Errorf("записано %d байт, а в writer ушло %d — замена обязана сохранять длину",
			total, buf.Len())
	}
}

var errBoom = errors.New("boom")

// failingWriter принимает не больше limit байт, дальше возвращает ошибку.
type failingWriter struct {
	limit   int
	written int
}

func (f *failingWriter) Write(p []byte) (int, error) {
	room := f.limit - f.written
	if room <= 0 {
		return 0, errBoom
	}
	if room >= len(p) {
		f.written += len(p)
		return len(p), nil
	}
	f.written += room
	return room, errBoom
}

// Ошибка обёрнутого writer'а обязана доходить до вызывающего.
func TestWritePropagatesError(t *testing.T) {
	w := New(&failingWriter{limit: 4}, nil)

	n, err := w.Write([]byte("0123456789"))
	if !errors.Is(err, errBoom) {
		t.Fatalf("Write ошибка = %v, ожидалось errors.Is(err, errBoom) == true", err)
	}
	if n != 4 {
		t.Errorf("Write = %d байт, ожидалось 4 — столько принял обёрнутый writer", n)
	}
}
