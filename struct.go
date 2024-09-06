package gomatcher

import (
	"fmt"
	"reflect"

	"go.uber.org/mock/gomock"
)

type StructIncludes map[string]any

func (ex StructIncludes) Matches(actual any) bool {
	actualrv := reflect.ValueOf(actual)
	if actualrv.Kind() == reflect.Pointer {
		actualrv = actualrv.Elem()
	}

	for exKey, exValue := range ex {
		actualFieldrv := actualrv.FieldByName(exKey)
		if !actualFieldrv.IsValid() {
			return false
		}

		exrv := reflect.ValueOf(exValue)
		if exrv.CanConvert(actualFieldrv.Type()) {
			exrv = exrv.Convert(actualFieldrv.Type())
		}

		if !gomock.Eq(exrv.Interface()).Matches(actualFieldrv.Interface()) {
			return false
		}
	}

	return true
}

func (m StructIncludes) String() string {
	return fmt.Sprint(map[string]any(m))
}

func (m StructIncludes) Got(got any) string {
	return fmt.Sprintf("%+v", got)
}
