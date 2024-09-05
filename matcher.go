package gomatcher

import (
	"reflect"

	"go.uber.org/mock/gomock"
)

type StructIncludes map[string]any

func (m StructIncludes) Matches(x any) bool {
	xrv := reflect.ValueOf(x)

	for key, value := range m {
		xFieldrv := xrv.FieldByName(key)

		if !gomock.Eq(xFieldrv.Interface()).Matches(value) {
			return false
		}
	}

	return true
}

func (m StructIncludes) String() string {
	return "apanih"
}
