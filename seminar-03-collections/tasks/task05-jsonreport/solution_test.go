package jsonreport

import (
	"testing"
)

func TestAggregate(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "несколько проектов с несколькими сборками",
			input: `[
				{"project":"a","status":"ok","seconds":10},
				{"project":"b","status":"failed","seconds":5},
				{"project":"a","status":"failed","seconds":20},
				{"project":"b","status":"ok","seconds":15}
			]`,
			want: `[{"project":"a","total":2,"failed":1,"seconds":30},{"project":"b","total":2,"failed":1,"seconds":20}]`,
		},
		{
			name:  "результат отсортирован по Project, хотя во входных данных порядок обратный",
			input: `[{"project":"zeta","status":"ok","seconds":1},{"project":"alpha","status":"ok","seconds":2}]`,
			want:  `[{"project":"alpha","total":1,"failed":0,"seconds":2},{"project":"zeta","total":1,"failed":0,"seconds":1}]`,
		},
		{
			name: "подсчёт Failed при смеси статусов",
			input: `[
				{"project":"a","status":"failed","seconds":1},
				{"project":"a","status":"failed","seconds":1},
				{"project":"a","status":"ok","seconds":1}
			]`,
			want: `[{"project":"a","total":3,"failed":2,"seconds":3}]`,
		},
		{
			name:  "пустой массив на входе даёт пустой массив на выходе",
			input: `[]`,
			want:  `[]`,
		},
		{
			name:  "null на входе даёт пустой массив на выходе",
			input: `null`,
			want:  `[]`,
		},
		{
			name:  "неизвестные поля входных объектов игнорируются",
			input: `[{"project":"a","status":"ok","seconds":5,"extra":"ignore me","nested":{"x":1}}]`,
			want:  `[{"project":"a","total":1,"failed":0,"seconds":5}]`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Aggregate([]byte(tc.input))
			if err != nil {
				t.Fatalf("Aggregate(%s) вернул ошибку %v, ожидалась nil", tc.input, err)
			}

			if string(got) != tc.want {
				t.Fatalf("Aggregate(%s) = %s, ожидалось %s", tc.input, got, tc.want)
			}
		})
	}
}

func TestAggregateErrors(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{
			name:  "битый JSON",
			input: `[{"project":"a","status":"ok","seconds":1},]`,
		},
		{
			name:  "не JSON вовсе",
			input: `это не json`,
		},
		{
			name:  "валидный JSON, но не массив (объект)",
			input: `{"project":"a","status":"ok","seconds":1}`,
		},
		{
			name:  "валидный JSON, но не той формы (массив чисел)",
			input: `[1, 2, 3]`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Aggregate([]byte(tc.input))
			if err == nil {
				t.Fatalf("Aggregate(%s) ошибка = nil, ожидалась не-nil ошибка", tc.input)
			}

			if got != nil {
				t.Fatalf("Aggregate(%s) = %s, ожидался nil результат при ошибке", tc.input, got)
			}
		})
	}
}
