package sort_test

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/prichrd/concepts/pkg/sort"
	"github.com/stretchr/testify/require"
)

var sortFunctions = []sortFunction[int]{
	sort.Selection[int],
	sort.Bubble[int],
	sort.Insertion[int],
	sort.Merge[int],
	sort.Quick[int],
	sort.Count[int],
}

var testCases = []testCase[int]{
	{[]int{1, 2, 3, 4, 5}},
	{[]int{5, 4, 3, 2, 1}},
	{[]int{1, 3, 4, 5, 2}},
	{[]int{1, 2, 1, 2, 1}},
	{[]int{5, 5, 5, 1, 1}},
	{[]int{1, 1, 5, 5, 5}},
	{[]int{1, 1, 1, 1, 1}},
	{[]int{}},
	{[]int{1}},
	{[]int{1, 2}},
	{[]int{2, 1}},
	{[]int{8, 7, 6, 1, 0, 9, 2}},
	{[]int{4, 2, 2, 8, 3, 3, 1}},
	{[]int{121, 432, 564, 23, 1, 45, 788}},
}

func TestSort(t *testing.T) {
	for _, fx := range sortFunctions {
		sort := fx
		t.Run(functionName(fx), func(t *testing.T) {
			for _, tc := range testCases {
				for _, asc := range []bool{true, false} {
					t.Run(testCaseName(tc.unsorted, asc), func(t *testing.T) {
						t.Parallel()
						sorted := make([]int, len(tc.unsorted))
						copy(sorted, tc.unsorted)
						sort(sorted, asc)
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
