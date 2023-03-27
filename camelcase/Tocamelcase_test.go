package camelcase_test

import (
	"testing"

	"github.com/g91TeJl/Algorithm/camelcase"
	"github.com/stretchr/testify/assert"
)

func TestToCamelCase(t *testing.T) {
	testCases := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "len == 0",
			str:  "",
			want: "",
		},
		{
			name: "len == 1",
			str:  "a",
			want: "a",
		},
		{
			name: "last -",
			str:  "last-",
			want: "last",
		},
		{
			name: "last _",
			str:  "last_",
			want: "last",
		},
		{
			name: "Street_journal",
			str:  "Street_journal",
			want: "StreetJournal",
		},
		{
			name: "Street_-journal",
			str:  "Street_-journal",
			want: "StreetJournal",
		},
		{
			name: "Street_journal_--_",
			str:  "Street_journal_--_",
			want: "StreetJournal",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, camelcase.ToCamelCase(tc.str))
		})
	}
}
