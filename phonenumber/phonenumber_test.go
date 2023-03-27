package phonenumber_test

import (
	"testing"

	"github.com/g91TeJl/Algorithm/phonenumber"
	"github.com/stretchr/testify/assert"
)

func TestCreatePhoneNumber(t *testing.T) {
	testCase := []struct {
		name    string
		numbers [10]uint
		want    string
	}{
		{
			name:    "normal numbers",
			numbers: [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			want:    "(123) 456-7890",
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, phonenumber.CreatePhoneNumber(tc.numbers))
		})
	}
}
