package strutils

import (
	"reflect"
	"unsafe"
)

func B2S(byts []byte) string {
	return *(*string)(unsafe.Pointer(&byts))
}

func S2B(str string) []byte {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&str))
	slice := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&slice))
}
