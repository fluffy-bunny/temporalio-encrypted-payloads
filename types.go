package cryptconverter

import (
	"reflect"
)

type ProtectMe struct {
}
type SensitivePayload struct {
	ProtectMe *ProtectMe
	Secret    string
}
type OpenPayload struct {
	Quote string
}

var (
	cacheProtectMe = make(map[reflect.Type]bool)
)

func HasProtectMe(obj interface{}) bool {

	t := reflect.TypeOf(obj)

	for {
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		} else {
			break
		}
	}

	has, ok := cacheProtectMe[t]
	if ok {
		return has
	}

	if t.Kind() != reflect.Struct {
		cacheProtectMe[t] = false
		return false
	}

	_, ok = t.FieldByName("ProtectMe")
	cacheProtectMe[t] = ok

	return ok

}
