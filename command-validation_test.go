package elo

import "testing"
import "github.com/stretchr/testify/assert"

var registerCommandTests = []struct {
	in  []string
	out bool
}{
	{[]string{"register", "bruce"}, true},
	{[]string{"register", "a"}, true},
	{[]string{"register", "Diana"}, true},
	{[]string{"register", ""}, false},
	{[]string{"register", "   "}, false},
	{[]string{"register"}, false},
}

func TestIsValidRegisterCommand(t *testing.T) {
	for _, tt := range registerCommandTests {
		actual := IsValidRegisterCommand(tt.in)
		assert.Equal(t, tt.out, actual)
	}
}

func IsValidRegisterCommand(args []string) bool {
	return len(args) == 2 && len(args[1]) > 0
}
