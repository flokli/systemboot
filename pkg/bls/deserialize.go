package bls

import (
	"fmt"
	"reflect"
)

func Deserialize(kvMap KVMap, Iface interface{}) (err error) {
	ValueIface := reflect.ValueOf(Iface)
	T := ValueIface.Elem().Type()

	if k := T.Kind(); k != reflect.Struct {
		return fmt.Errorf("result is no struct but %s", k)
	}

	// iterate over fields of struct
	for i := 0; i < ValueIface.Elem().NumField(); i++ {
		FieldType := T.Field(i) // reflect.Type
		FieldValue := ValueIface.Elem().Field(i)

		keyTag, ok := FieldType.Tag.Lookup("key")
		if !ok {
			// field has no field tag, ignore
			continue
		}
		v, ok := kvMap[keyTag]
		if !ok {
			// kvMap does not contain the key specified in field tag, ignore
			continue
		}
		if FieldType.Type.Kind() == reflect.String {
			if len(v) == 1 {
				FieldValue.SetString(v[0])
			} else {
				return fmt.Errorf("wanted to set %s (should contain string) to %v (list)!", keyTag, v)
			}
		} else if FieldType.Type == reflect.TypeOf([]string{}) {
			FieldValue.Set(reflect.ValueOf(v))
		} else {
			// unimplemented struct type, ignore
			continue
		}
	}
	return err
}
