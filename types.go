package cryptconverter

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/descriptor"
	pb "github.com/temporalio/samples-go/encrypted-payloads/helloworld"
	"google.golang.org/protobuf/proto"
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
	cacheProtectMe      = make(map[reflect.Type]bool)
	rtDescripterMessage = reflect.TypeOf((*descriptor.Message)(nil)).Elem()
)

func HasProtectMeOption(obj interface{}) bool {
	var msg descriptor.Message

	// need the value because we can only get the interface from it
	va := reflect.ValueOf(obj)
	// the type of the current value
	t := va.Type()

	for {
		if t.Kind() == reflect.Ptr {
			if t.Implements(rtDescripterMessage) {
				obj = va.Interface()
				msg = obj.(descriptor.Message)
				break
			}
			va = va.Elem()
			t = va.Type()
		} else {
			break
		}
	}

	if msg == nil {
		return false
	}
	_, md := descriptor.ForMessage(msg)
	options := md.GetOptions()
	fmt.Println(options)
	a := proto.GetExtension(options, pb.E_ProtectMe)
	result, _ := a.(bool)
	fmt.Println(result)
	return result
}

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
