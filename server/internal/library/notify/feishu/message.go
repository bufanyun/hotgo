package feishu

import (
	"reflect"
	"strings"
)

type MsgType string

// MsgType
const (
	MsgTypeText        MsgType = "text"
	MsgTypePost        MsgType = "post"
	MsgTypeImage       MsgType = "image"
	MsgTypeShareChat   MsgType = "share_chat"
	MsgTypeInteractive MsgType = "interactive"
)

// Message interface
type Message interface {
	Body() map[string]interface{}
}

func structToMap(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")

		// remove omitEmpty
		omitEmpty := false
		if strings.HasSuffix(tag, "omitempty") {
			omitEmpty = true
			idx := strings.Index(tag, ",")
			if idx > 0 {
				tag = tag[:idx]
			} else {
				tag = ""
			}
		}

		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if omitEmpty && reflectValue.Field(i).IsZero() {
				continue
			}

			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = structToMap(field)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}
