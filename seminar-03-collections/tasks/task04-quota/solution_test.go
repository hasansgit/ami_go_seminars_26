package quota

import (
	"errors"
	"reflect"
	"testing"
)

func TestApplySuccess(t *testing.T) {
	testCases := []struct {
		name   string
		limits map[string]int
		ops    []Op
		want   []Usage
	}{
		{
			name:   "несколько пользователей и несколько операций на каждого",
			limits: map[string]int{"alice": 100, "bob": 80, "carol": 50},
			ops: []Op{
				{User: "bob", Amount: 20},
				{User: "alice", Amount: 30},
				{User: "bob", Amount: 10},
				{User: "alice", Amount: 20},
				{User: "carol", Amount: 50},
			},
			want: []Usage{
				{User: "alice", Total: 50},
				{User: "bob", Total: 30},
				{User: "carol", Total: 50},
			},
		},
		{
			name:   "результат отсортирован по User, хотя операции идут в обратном алфавитном порядке",
			limits: map[string]int{"zoe": 10, "amy": 10},
			ops: []Op{
				{User: "zoe", Amount: 5},
				{User: "amy", Amount: 5},
			},
			want: []Usage{
				{User: "amy", Total: 5},
				{User: "zoe", Total: 5},
			},
		},
		{
			name:   "расход ровно в лимит допустим",
			limits: map[string]int{"alice": 10},
			ops: []Op{
				{User: "alice", Amount: 10},
			},
			want: []Usage{
				{User: "alice", Total: 10},
			},
		},
		{
			name:   "ops == nil даёт пустой результат",
			limits: map[string]int{"alice": 10},
			ops:    nil,
			want:   []Usage{},
		},
		{
			name:   "ops — пустой слайс даёт пустой результат",
			limits: map[string]int{"alice": 10},
			ops:    []Op{},
			want:   []Usage{},
		},
		{
			name:   "limits == nil и ops == nil даёт пустой результат",
			limits: nil,
			ops:    nil,
			want:   []Usage{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Apply(tc.limits, tc.ops)
			if err != nil {
				t.Fatalf("Apply(%v, %v) вернул ошибку %v, ожидалась nil", tc.limits, tc.ops, err)
			}

			if got == nil {
				t.Fatalf("Apply(%v, %v) = nil, ожидался не-nil слайс", tc.limits, tc.ops)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Apply(%v, %v) = %v, ожидалось %v", tc.limits, tc.ops, got, tc.want)
			}
		})
	}
}

func TestApplyErrors(t *testing.T) {
	testCases := []struct {
		name    string
		limits  map[string]int
		ops     []Op
		wantErr error
	}{
		{
			name:   "превышение лимита",
			limits: map[string]int{"alice": 10},
			ops: []Op{
				{User: "alice", Amount: 5},
				{User: "alice", Amount: 10},
			},
			wantErr: ErrOverQuota,
		},
		{
			name:   "неизвестный пользователь",
			limits: map[string]int{"alice": 10},
			ops: []Op{
				{User: "bob", Amount: 5},
			},
			wantErr: ErrUnknownUser,
		},
		{
			name:   "неизвестный пользователь при nil limits",
			limits: nil,
			ops: []Op{
				{User: "alice", Amount: 5},
			},
			wantErr: ErrUnknownUser,
		},
		{
			name:   "нулевой Amount недопустим",
			limits: map[string]int{"alice": 10},
			ops: []Op{
				{User: "alice", Amount: 0},
			},
			wantErr: ErrInvalidAmount,
		},
		{
			name:   "отрицательный Amount недопустим",
			limits: map[string]int{"alice": 10},
			ops: []Op{
				{User: "alice", Amount: -5},
			},
			wantErr: ErrInvalidAmount,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Apply(tc.limits, tc.ops)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Apply(%v, %v) ошибка = %v, ожидалось errors.Is(err, %v) == true", tc.limits, tc.ops, err, tc.wantErr)
			}

			if got != nil {
				t.Fatalf("Apply(%v, %v) = %v, ожидался nil результат при ошибке", tc.limits, tc.ops, got)
			}
		})
	}
}

func TestApplyDoesNotModifyLimits(t *testing.T) {
	testCases := []struct {
		name   string
		limits map[string]int
		ops    []Op
	}{
		{
			name:   "успешное применение",
			limits: map[string]int{"alice": 10, "bob": 5},
			ops: []Op{
				{User: "alice", Amount: 10},
				{User: "bob", Amount: 3},
			},
		},
		{
			name:   "применение, завершившееся ошибкой превышения лимита",
			limits: map[string]int{"alice": 10},
			ops: []Op{
				{User: "alice", Amount: 5},
				{User: "alice", Amount: 10},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			limitsCopy := make(map[string]int, len(tc.limits))
			for k, v := range tc.limits {
				limitsCopy[k] = v
			}

			Apply(tc.limits, tc.ops)

			if !reflect.DeepEqual(tc.limits, limitsCopy) {
				t.Fatalf("после Apply limits = %v, ожидалось без изменений %v", tc.limits, limitsCopy)
			}
		})
	}
}

func TestApplyDoesNotModifyOps(t *testing.T) {
	limits := map[string]int{"alice": 100, "bob": 100}
	ops := []Op{
		{User: "alice", Amount: 5},
		{User: "bob", Amount: 3},
		{User: "alice", Amount: 2},
	}
	opsCopy := make([]Op, len(ops))
	copy(opsCopy, ops)

	Apply(limits, ops)

	if !reflect.DeepEqual(ops, opsCopy) {
		t.Fatalf("после Apply ops = %v, ожидалось без изменений %v", ops, opsCopy)
	}
}
