package request

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testRunner struct {
	id       int
	name     string
	scenario func(*testing.T)
}

var tests = []testRunner{
	{
		id:   1,
		name: "test",
		scenario: func(t *testing.T) {
			mj := MyJSON{}
			if err := mj.New([]byte(`{"hoge":null, "fuga":"0", "piyo":3, "foo": [0,4], "bar": "2a"}`)); err != nil {
				fmt.Println(err)
			}
			mjFields := mj.Fields()
			fmt.Printf("%#v\n", mjFields)

			expected := map[string]interface{}{
				"bar":  "2a",
				"foo":  []interface{}{float64(0), float64(4)},
				"fuga": "0",
				"hoge": interface{}(nil),
				"piyo": float64(3),
			}
			assert.Equal(t, expected, mjFields)
			//vMap := map[string][]validation.Rule{
			//	"hoge": {validation.Nil},
			//	"fuga": {validation.Required, is.Alphanumeric},
			//	"piyo": {validation.Required, is.Int},
			//	"foo":  {validation.Required},
			//}
			//mj.Validate(vMap)
		},
	},
}

func TestMyJSON_Fields(t *testing.T) {
	tests := tests

	for _, tt := range tests {
		t.Run(tt.name, tt.scenario)
	}
}
