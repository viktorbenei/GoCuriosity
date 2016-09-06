package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestSubSerModel struct {
	TestStringTwo string `json:"test_string_two,omitempty" yaml:"test_string_two,omitempty"`
}

type TestSerModel struct {
	TestString    string            `json:"test_string,omitempty" yaml:"test_string,omitempty"`
	TestBool      bool              `json:"test_bool,omitempty" yaml:"test_bool,omitempty"`
	SubMod        TestSubSerModel   `json:"sub_mod,omitempty" yaml:"sub_mod,omitempty"`
	PointerSubMod *TestSubSerModel  `json:"pointer_sub_mod,omitempty" yaml:"pointer_sub_mod,omitempty"`
	TestMap       map[string]string `json:"test_map,omitempty" yaml:"test_map,omitempty"`
	TestSlice     []string          `json:"test_slice,omitempty" yaml:"test_slice,omitempty"`
}

func Test_serialize(t *testing.T) {
	t.Log("JSON - serialize - empty struct")
	{
		testObj := TestSerModel{}
		bytes, err := json.Marshal(testObj)
		require.NoError(t, err)
		// default values will emit everything *except* the sub struct
		require.Equal(t, `{"sub_mod":{}}`, string(bytes))
	}

	t.Log("JSON - deserialize - empty json to struct")
	{
		var testObj TestSerModel
		err := json.Unmarshal([]byte(`{}`), &testObj)
		require.NoError(t, err)
		require.Equal(t, "", testObj.TestString)
		require.Equal(t, false, testObj.TestBool)
		// a sub struct is not nil
		require.Equal(t, TestSubSerModel{}, testObj.SubMod)
		// unless you use a pointer (*)
		require.Equal(t, (*TestSubSerModel)(nil), testObj.PointerSubMod)
		require.NotEqual(t, nil, testObj.PointerSubMod)
		// but can be compared to a `nil` with `==`
		require.Equal(t, true, testObj.PointerSubMod == nil)
		// but maps and slices are!
		require.Equal(t, map[string]string(nil), testObj.TestMap)
		require.Equal(t, []string(nil), testObj.TestSlice)
	}
}
