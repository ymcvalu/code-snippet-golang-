package jsontype

import (
	"encoding/gob"
)

func init() {
	gob.Register(JsonStr(""))
}

type JsonStr string

func (t JsonStr) MarshalJSON() ([]byte, error) {
	return []byte(t), nil
}

func (t *JsonStr) UnmarshalJSON(byts []byte) error {
	*t = JsonStr(string(byts))

	return nil
}
