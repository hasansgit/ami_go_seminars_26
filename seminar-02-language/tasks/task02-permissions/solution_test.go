package permissions

import "testing"

func TestGrant(t *testing.T) {
	current := Read
	got := Grant(current, Write|Execute)
	want := Read | Write | Execute

	if got != want {
		t.Fatalf("Grant(%03b, %03b) = %03b, ожидалось %03b", current, Write|Execute, got, want)
	}

	if current != Read {
		t.Fatalf("Grant изменил свой аргумент: current = %03b, ожидалось %03b", current, Read)
	}
}

func TestRevoke(t *testing.T) {
	all := Read | Write | Execute
	if got, want := Revoke(all, Write), Read|Execute; got != want {
		t.Fatalf("Revoke(%03b, %03b) = %03b, ожидалось %03b", all, Write, got, want)
	}

	if got := Revoke(Read, Write); got != Read {
		t.Fatalf("отзыв отсутствующего права: got %03b, ожидалось %03b", got, Read)
	}
}

func TestHas(t *testing.T) {
	current := Read | Execute
	testCases := []struct {
		name     string
		required Permission
		want     bool
	}{
		{name: "одно право присутствует", required: Read, want: true},
		{name: "одно право отсутствует", required: Write, want: false},
		{name: "все права присутствуют", required: Read | Execute, want: true},
		{name: "только часть прав присутствует", required: Read | Write, want: false},
		{name: "пустой набор прав", required: 0, want: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Has(current, tc.required); got != tc.want {
				t.Fatalf("Has(%03b, %03b) = %t, ожидалось %t", current, tc.required, got, tc.want)
			}
		})
	}
}
