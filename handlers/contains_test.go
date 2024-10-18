package handlers

import "testing"

var Tests = []struct {
	Name   string
	Id     int
	Ids    []int
	Result bool
}{
	{"True test", 4, []int{1, 2, 3, 4, 5, 6, 7}, true},
	{"False test", 34, []int{1, 2, 3, 4, 5, 6, 7}, false},
}

func TestContains(t *testing.T) {
	for _, tc := range Tests {
		t.Run(tc.Name, func(t *testing.T) {
			result := Contains(tc.Id, tc.Ids)
			if result != tc.Result {
				t.Errorf("Expected %v, got %v", tc.Result, result)
			}
		})
	}
}
