package binary_search_test

import (
	"testing"

	"github.com/g91TeJl/Algorithm/binary_search"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	testCase := []struct {
		name string
		list []int
		item int
		want interface{}
	}{
		{
			name: "len == 0",
			list: []int{},
			item: 1,
			want: -1,
		},
		{
			name: "len == 1",
			list: []int{1},
			item: 1,
			want: 0,
		},
		{
			name: "unsort",
			list: []int{3, 20, 4, 50, 1, 6},
			item: 3,
			want: 0,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, binary_search.Binary_search(tc.list, tc.item))
		})
	}
}
