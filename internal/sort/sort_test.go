package sort_test

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/prichrd/concepts/internal/sort"
	"github.com/stretchr/testify/require"
)

var sortFunctions = []sortFunction[int]{
	sort.Selection[int],
	sort.Bubble[int],
	sort.Insertion[int],
}

var testCases = []testCase[int]{
	{[]int{1, 2, 3, 4, 5}},
	{[]int{5, 4, 3, 2, 1}},
	{[]int{1, 3, 4, 5, 2}},
	{[]int{1, 2, 1, 2, 1}},
	{[]int{5, 5, 5, 1, 1}},
	{[]int{1, 1, 5, 5, 5}},
	{[]int{}},
	{[]int{1}},
	{[]int{1, 2}},
	{[]int{2, 1}},
}

func TestSort(t *testing.T) {
	for _, fx := range sortFunctions {
		t.Run(functionName(fx), func(t *testing.T) {
			for _, tc := range testCases {
				for _, asc := range []bool{false, true} {
					t.Run(testCaseName(tc.unsorted, asc), func(t *testing.T) {
						sorted := make([]int, len(tc.unsorted))
						copy(sorted, tc.unsorted)
						fx(sorted, asc)
						require.ElementsMatch(t, sorted, tc.unsorted)
						assertOrdered(t, sorted, asc)
					})
				}
			}
		})
	}
}

func assertOrdered(t *testing.T, s []int, asc bool) {
	failed := false
	for i := 1; i < len(s); i++ {
		if asc {
			if s[i-1] > s[i] {
				failed = true
			}
		} else {
			if s[i-1] < s[i] {
				failed = true
			}
		}
	}
	if failed {
		t.Errorf("got %v", s)
	}
}

type sortFunction[V comparable] func([]V, bool)

type testCase[V comparable] struct {
	unsorted []V
}

func functionName[V comparable](f sortFunction[V]) string {
	rfp := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	rfp = rfp[strings.LastIndex(rfp, "/")+1:]
	rfp = strings.ReplaceAll(rfp, "[...]", "")
	return rfp
}

func testCaseName[V comparable](s []V, asc bool) string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, v := range s {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf("%v", v))
	}
	sb.WriteString("]")
	if !asc {
		sb.WriteString("(reverse)")
	}
	return sb.String()
}
