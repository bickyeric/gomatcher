package gomatcher_test

import (
	"testing"

	"github.com/bickyeric/gomatcher"
	"github.com/bickyeric/gomatcher/internal/usecase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestStructIncludes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testcase struct {
		Matcher  gomatcher.StructIncludes
		Expected bool
	}

	actualValue := &usecase.User{
		ID:   90,
		Name: "John Doe",
	}
	testcases := []testcase{
		{gomatcher.StructIncludes{"ID": 90}, true},
		{gomatcher.StructIncludes{"ID": 81}, false},
		{gomatcher.StructIncludes{"id": 90}, false},
		{gomatcher.StructIncludes{"Name": "foo"}, false},
		{gomatcher.StructIncludes{"Name": "John Doe"}, true},
	}

	for index, tc := range testcases {
		assert.Equal(t, tc.Expected, tc.Matcher.Matches(actualValue), "testcase: %d", index)
	}
}
