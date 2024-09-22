package differ

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	data := map[string]any{
		"3": 3232,
		"5": 234234,
		"6": map[string]any{
			"5": 123,
			"7": "hello",
		},
		//"blah": "bleh",
	}
	jbytes, err := json.Marshal(data)
	assert.Nil(t, err)
	fmt.Println(string(jbytes))

	var val any
	val = map[string]any{}
	err = json.Unmarshal(jbytes, &val)
	fmt.Println(err)
	fmt.Println(val)
	fmt.Printf("%T \n", val.(map[string]any)["6"])
}
