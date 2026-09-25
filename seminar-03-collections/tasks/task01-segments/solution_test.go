package segments

import (
	"errors"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	cases := []struct {
		name    string
		data    []int
		size    int
		want    [][]int
		wantErr error
	}{
		{
			name: "деление без остатка",
			data: []int{1, 2, 3, 4},
			size: 2,
			want: [][]int{{1, 2}, {3, 4}},
		},
		{
			name: "деление с остатком",
			data: []int{1, 2, 3, 4, 5},
			size: 2,
			want: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name: "size больше длины data",
			data: []int{1, 2, 3},
			size: 10,
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "size равен 1",
			data: []int{1, 2, 3},
			size: 1,
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name:    "size равен нулю даёт ErrInvalidSize",
			data:    []int{1, 2, 3},
			size:    0,
			want:    nil,
			wantErr: ErrInvalidSize,
		},
		{
			name:    "отрицательный size даёт ErrInvalidSize",
			data:    []int{1, 2, 3},
			size:    -1,
			want:    nil,
			wantErr: ErrInvalidSize,
		},
		{
			name: "nil data даёт nil",
			data: nil,
			size: 2,
			want: nil,
		},
		{
			name: "не-nil пустой data даёт не-nil пустой результат",
			data: []int{},
			size: 2,
			want: [][]int{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Split(tc.data, tc.size)

			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					t.Fatalf("Split(%v, %d) вернул ошибку %v, ожидалось %v", tc.data, tc.size, err, tc.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("Split(%v, %d) вернул ошибку %v, ожидалось nil", tc.data, tc.size, err)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Split(%v, %d) = %#v, ожидалось %#v", tc.data, tc.size, got, tc.want)
			}
		})
	}
}

func TestSplitSharesBackingArray(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	chunks, err := Split(data, 2)
	if err != nil {
		t.Fatalf("Split(%v, 2) вернул ошибку %v, ожидалось nil", data, err)
	}

	chunks[0][0] = 100
	if data[0] != 100 {
		t.Fatalf("запись в chunks[0][0] не видна в data: data = %v, ожидалось data[0] == 100", data)
	}

	data[2] = 200
	if chunks[1][0] != 200 {
		t.Fatalf("запись в data[2] не видна в chunks[1][0]: chunks = %v, ожидалось chunks[1][0] == 200", chunks)
	}
}

func TestSplitChunkCapacityEqualsLength(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	chunks, err := Split(data, 2)
	if err != nil {
		t.Fatalf("Split(%v, 2) вернул ошибку %v, ожидалось nil", data, err)
	}

	for i, chunk := range chunks {
		if cap(chunk) != len(chunk) {
			t.Fatalf("Split(%v, 2): cap(chunks[%d]) = %d, ожидалось %d (равно len)", data, i, cap(chunk), len(chunk))
		}
	}
}

func TestSplitDoesNotChangeData(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	before := append([]int(nil), data...)

	_, err := Split(data, 2)
	if err != nil {
		t.Fatalf("Split(%v, 2) вернул ошибку %v, ожидалось nil", before, err)
	}

	if !reflect.DeepEqual(data, before) {
		t.Fatalf("после Split data = %v, ожидалось %v (без изменений)", data, before)
	}
}

func TestCut(t *testing.T) {
	cases := []struct {
		name          string
		data          []int
		from, to      int
		want          []int
		wantErr       error
		ignoreNilness bool
	}{
		{
			name: "вырезание из середины",
			data: []int{1, 2, 3, 4, 5},
			from: 1, to: 3,
			want: []int{1, 4, 5},
		},
		{
			name: "вырезание с начала",
			data: []int{1, 2, 3, 4, 5},
			from: 0, to: 2,
			want: []int{3, 4, 5},
		},
		{
			name: "вырезание с конца",
			data: []int{1, 2, 3, 4, 5},
			from: 3, to: 5,
			want: []int{1, 2, 3},
		},
		{
			name: "from равен to не меняет содержимое",
			data: []int{1, 2, 3},
			from: 1, to: 1,
			want: []int{1, 2, 3},
		},
		{
			name: "вырезание всего слайса",
			data: []int{1, 2, 3},
			from: 0, to: 3,
			want: []int{},
		},
		{
			name: "to больше длины data даёт ErrRange",
			data: []int{1, 2, 3},
			from: 1, to: 5,
			wantErr: ErrRange,
		},
		{
			name: "отрицательный from даёт ErrRange",
			data: []int{1, 2, 3},
			from: -1, to: 2,
			wantErr: ErrRange,
		},
		{
			name: "from больше to даёт ErrRange",
			data: []int{1, 2, 3},
			from: 3, to: 1,
			wantErr: ErrRange,
		},
		{
			name: "nil data с нулевым диапазоном",
			data: nil,
			from: 0, to: 0,
			want: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			before := tc.data
			got, err := Cut(tc.data, tc.from, tc.to)

			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					t.Fatalf("Cut(%v, %d, %d) вернул ошибку %v, ожидалось %v", before, tc.from, tc.to, err, tc.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("Cut(%v, %d, %d) вернул ошибку %v, ожидалось nil", before, tc.from, tc.to, err)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Cut(%v, %d, %d) = %v, ожидалось %v", before, tc.from, tc.to, got, tc.want)
			}
		})
	}
}

func TestCutZeroesTail(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	originalLen := len(data)
	originalCap := cap(data)

	got, err := Cut(data, 1, 3)
	if err != nil {
		t.Fatalf("Cut(%v, 1, 3) вернул ошибку %v, ожидалось nil", data, err)
	}

	newLen := len(got)
	full := data[:originalCap]

	for i := newLen; i < originalLen; i++ {
		if full[i] != 0 {
			t.Fatalf("после Cut элемент data[:cap(data)][%d] = %d, ожидался 0 (хвост за новой длиной)", i, full[i])
		}
	}
}
