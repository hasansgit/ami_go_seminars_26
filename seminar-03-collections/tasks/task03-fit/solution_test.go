package fit

import (
	"testing"
	"unicode/utf8"
)

func TestFit(t *testing.T) {
	testCases := []struct {
		name       string
		s          string
		limit      int
		wantFitted string
		wantTrunc  bool
	}{
		{
			name:       "строка помещается целиком",
			s:          "hello",
			limit:      10,
			wantFitted: "hello",
			wantTrunc:  false,
		},
		{
			name:       "limit ровно равен длине строки",
			s:          "hello",
			limit:      5,
			wantFitted: "hello",
			wantTrunc:  false,
		},
		{
			name:       "ascii-строка с обрезкой",
			s:          "hello world",
			limit:      8,
			wantFitted: "hello…",
			wantTrunc:  true,
		},
		{
			name:       "кириллица, limit попадает в середину руны",
			s:          "привет",
			limit:      6,
			wantFitted: "п…",
			wantTrunc:  true,
		},
		{
			name:       "эмодзи, limit не позволяет вместить ни одной руны из префикса",
			s:          "😀😀😀",
			limit:      6,
			wantFitted: "…",
			wantTrunc:  true,
		},
		{
			name:       "эмодзи, в префикс помещается одна руна",
			s:          "😀😀😀",
			limit:      9,
			wantFitted: "😀…",
			wantTrunc:  true,
		},
		{
			name:       "limit равен 3",
			s:          "hello",
			limit:      3,
			wantFitted: "…",
			wantTrunc:  true,
		},
		{
			name:       "limit равен 2",
			s:          "hello",
			limit:      2,
			wantFitted: "",
			wantTrunc:  true,
		},
		{
			name:       "limit равен 0",
			s:          "hello",
			limit:      0,
			wantFitted: "",
			wantTrunc:  true,
		},
		{
			name:       "отрицательный limit",
			s:          "hello",
			limit:      -5,
			wantFitted: "",
			wantTrunc:  true,
		},
		{
			name:       "пустая строка при положительном limit",
			s:          "",
			limit:      5,
			wantFitted: "",
			wantTrunc:  false,
		},
		{
			name:       "пустая строка при limit 0",
			s:          "",
			limit:      0,
			wantFitted: "",
			wantTrunc:  false,
		},
		{
			name:       "пустая строка при отрицательном limit",
			s:          "",
			limit:      -3,
			wantFitted: "",
			wantTrunc:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotFitted, gotTrunc := Fit(tc.s, tc.limit)
			if gotFitted != tc.wantFitted || gotTrunc != tc.wantTrunc {
				t.Fatalf("Fit(%q, %d) = (%q, %t), ожидалось (%q, %t)",
					tc.s, tc.limit, gotFitted, gotTrunc, tc.wantFitted, tc.wantTrunc)
			}
		})
	}
}

func TestFitResultWithinByteLimit(t *testing.T) {
	cases := []struct {
		s     string
		limit int
	}{
		{"hello world", 8},
		{"привет", 6},
		{"привет", 7},
		{"😀😀😀", 6},
		{"😀😀😀", 9},
		{"hello", 3},
		{"hello", 100},
	}

	for _, tc := range cases {
		fitted, _ := Fit(tc.s, tc.limit)
		if len(fitted) > tc.limit {
			t.Fatalf("Fit(%q, %d): длина результата %q равна %d байт, что больше limit %d",
				tc.s, tc.limit, fitted, len(fitted), tc.limit)
		}
	}
}

func TestFitResultIsValidUTF8(t *testing.T) {
	cases := []struct {
		s     string
		limit int
	}{
		{"hello world", 8},
		{"привет", 6},
		{"привет", 7},
		{"😀😀😀", 6},
		{"😀😀😀", 9},
		{"hello", 3},
		{"hello", 2},
		{"hello", 0},
		{"hello", -5},
		{"", 5},
	}

	for _, tc := range cases {
		fitted, _ := Fit(tc.s, tc.limit)
		if !utf8.ValidString(fitted) {
			t.Fatalf("Fit(%q, %d) = %q — результат не является корректным UTF-8", tc.s, tc.limit, fitted)
		}
	}
}
