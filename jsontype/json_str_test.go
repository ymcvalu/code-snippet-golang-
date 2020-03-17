package jsontype

import (
	"encoding/json"
	"testing"
)

func TestJsonStr(t *testing.T) {
	type Obj struct {
		Js JsonStr `json:"json_str"`
	}

	obj := Obj{
		Js: `{"code":"0", "msg":"success"}`,
	}

	// marshal
	byts, _ := json.Marshal(obj)
	t.Log(string(byts))

	// unmarshal
	nobj := Obj{}
	json.Unmarshal(byts, &nobj)
	t.Log(nobj.Js)

}
