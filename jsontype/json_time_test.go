package jsontype

import (
	"encoding/json"
	"testing"
)

func TestJsonTime(t *testing.T) {
	tm := Now()
	// marshal
	byts, _ := json.Marshal(tm)
	t.Log(string(byts))

	// unmarshal
	newTm := new(JsonTime)
	json.Unmarshal(byts, newTm)
	t.Log(newTm.String())
}
