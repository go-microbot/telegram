package query

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AsMap(t *testing.T) {
	testStruct := struct {
		// int params.
		RequiredIntValue      ParamInt `query:"int_required"`
		RequiredIntEmptyValue ParamInt `query:"int_required_empty"`
		IntIgnoreValue        ParamInt `query:"int_ignore,omitempty"`
		IntValue              ParamInt
		// string slice params.
		RequiredStringSliceValue      ParamStringSlice `query:"string_slice_required"`
		RequiredStringSliceEmptyValue ParamStringSlice `query:"string_slice_required_empty"`
		StringSliceIgnoreValue        ParamStringSlice `query:"string_slice_ignore,omitempty"`
		StringSliceValue              ParamStringSlice
		// unparsed params.
		Hello string `query:"hello"`
		World int
	}{
		RequiredIntValue:         NewParamInt(100),
		IntValue:                 NewParamInt(200),
		RequiredStringSliceValue: NewParamStringSlice([]string{"1", "2", "3"}),
		StringSliceValue:         NewParamStringSlice([]string{"hello", "world"}),
	}

	expMap := map[string]string{
		"int_required":                "100",
		"IntValue":                    "200",
		"int_required_empty":          "",
		"string_slice_required":       "[\"1\",\"2\",\"3\"]",
		"StringSliceValue":            "[\"hello\",\"world\"]",
		"string_slice_required_empty": "[]",
	}
	result := AsMap(testStruct)

	require.Equal(t, expMap, result)
}
