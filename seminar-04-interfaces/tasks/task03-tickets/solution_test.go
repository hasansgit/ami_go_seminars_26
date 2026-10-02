package tickets

import (
	"reflect"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		name string
		list []Ticket
		want []Ticket
	}{
		{
			name: "разные приоритеты",
			list: []Ticket{
				{ID: 1, Priority: 3, Arrived: 1},
				{ID: 2, Priority: 1, Arrived: 2},
				{ID: 3, Priority: 2, Arrived: 3},
			},
			want: []Ticket{
				{ID: 2, Priority: 1, Arrived: 2},
				{ID: 3, Priority: 2, Arrived: 3},
				{ID: 1, Priority: 3, Arrived: 1},
			},
		},
		{
			name: "равный приоритет решает порядок поступления",
			list: []Ticket{
				{ID: 10, Priority: 1, Arrived: 30},
				{ID: 11, Priority: 1, Arrived: 10},
				{ID: 12, Priority: 1, Arrived: 20},
			},
			want: []Ticket{
				{ID: 11, Priority: 1, Arrived: 10},
				{ID: 12, Priority: 1, Arrived: 20},
				{ID: 10, Priority: 1, Arrived: 30},
			},
		},
		{
			name: "смешанный случай",
			list: []Ticket{
				{ID: 1, Priority: 2, Arrived: 1},
				{ID: 2, Priority: 1, Arrived: 5},
				{ID: 3, Priority: 1, Arrived: 2},
				{ID: 4, Priority: 2, Arrived: 0},
			},
			want: []Ticket{
				{ID: 3, Priority: 1, Arrived: 2},
				{ID: 2, Priority: 1, Arrived: 5},
				{ID: 4, Priority: 2, Arrived: 0},
				{ID: 1, Priority: 2, Arrived: 1},
			},
		},
		{
			name: "уже отсортирован",
			list: []Ticket{
				{ID: 1, Priority: 1, Arrived: 1},
				{ID: 2, Priority: 2, Arrived: 2},
			},
			want: []Ticket{
				{ID: 1, Priority: 1, Arrived: 1},
				{ID: 2, Priority: 2, Arrived: 2},
			},
		},
		{name: "пустой слайс", list: []Ticket{}, want: []Ticket{}},
		{name: "nil-слайс", list: nil, want: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list := make([]Ticket, len(tc.list))
			copy(list, tc.list)

			Sort(list)

			if len(tc.want) == 0 && len(list) == 0 {
				return
			}
			if !reflect.DeepEqual(list, tc.want) {
				t.Errorf("после Sort(%+v) = %+v, ожидалось %+v", tc.list, list, tc.want)
			}
		})
	}
}

// ByUrgency обязан удовлетворять sort.Interface — тогда заявки сортируются
// и через sort.Sort, без всякого Sort.
func TestByUrgencySatisfiesSortInterface(t *testing.T) {
	var _ sort.Interface = ByUrgency{}
}

func TestByUrgencyWorksWithSortPackage(t *testing.T) {
	list := []Ticket{
		{ID: 1, Priority: 5, Arrived: 1},
		{ID: 2, Priority: 1, Arrived: 9},
		{ID: 3, Priority: 1, Arrived: 2},
	}

	sort.Sort(ByUrgency(list))

	wantIDs := []int{3, 2, 1}
	for i, wantID := range wantIDs {
		if list[i].ID != wantID {
			t.Errorf("после sort.Sort на позиции %d заявка ID = %d, ожидалось %d (весь список: %+v)",
				i, list[i].ID, wantID, list)
		}
	}
}

// Less обязан задавать строгий порядок: Less(i, j) и Less(j, i) не могут
// быть истинны одновременно. Иначе поведение sort.Sort не определено.
func TestLessIsStrictOrder(t *testing.T) {
	list := []Ticket{
		{ID: 1, Priority: 1, Arrived: 1},
		{ID: 2, Priority: 1, Arrived: 2},
		{ID: 3, Priority: 2, Arrived: 0},
	}
	by := ByUrgency(list)

	for i := range list {
		if by.Less(i, i) {
			t.Errorf("Less(%d, %d) = true для одного и того же элемента %+v", i, i, list[i])
		}
		for j := range list {
			if by.Less(i, j) && by.Less(j, i) {
				t.Errorf("Less(%d, %d) и Less(%d, %d) оба true для %+v и %+v",
					i, j, j, i, list[i], list[j])
			}
		}
	}
}

// Swap обязан менять элементы местами, а не что-то ещё.
func TestSwapExchangesElements(t *testing.T) {
	list := []Ticket{
		{ID: 1, Priority: 1, Arrived: 1},
		{ID: 2, Priority: 2, Arrived: 2},
	}

	ByUrgency(list).Swap(0, 1)

	if list[0].ID != 2 || list[1].ID != 1 {
		t.Errorf("после Swap(0, 1) = %+v, ожидалось, что заявки поменяются местами", list)
	}
}
