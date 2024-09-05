package gomatcher_test

import (
	"testing"

	"github.com/bickyeric/gomatcher"
	"github.com/stretchr/testify/assert"
)

func TestStructIncludes(t *testing.T) {
	dummyStruct := struct {
		String string
		Bool   bool
	}{
		String: "golang gomatcher",
		Bool:   true,
	}

	testcases := []struct {
		Matcher  gomatcher.StructIncludes
		Expected bool
	}{
		// {gomatcher.StructIncludes{}, true},
		{gomatcher.StructIncludes{"String": "golang gomatcher"}, true},
		{gomatcher.StructIncludes{"String": "apa nih"}, false},
		{gomatcher.StructIncludes{"Bool": true}, true},
		{gomatcher.StructIncludes{"String": 123}, false},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.Expected, tc.Matcher.Matches(dummyStruct))
	}
}
