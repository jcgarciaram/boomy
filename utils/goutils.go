package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"gopkg.in/guregu/null.v3"
)

// GetType is used to get the name of a struct
func GetType(myvar interface{}) string {
	t := reflect.TypeOf(myvar)

	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}

	return t.Name()
}

// GetTypePointer is used to get the name of a struct
func GetTypePointer(myvar interface{}) (string, bool) {
	t := reflect.TypeOf(myvar)

	if t.Kind() == reflect.Ptr {
		return t.Elem().Name(), true
	}

	return t.Name(), false
}

// InterfaceToInt performs type conversion to convert an interface into an int.
func InterfaceToInt(v interface{}) (int, error) {

	if b, ok := v.(int); ok {
		return b, nil
	} else if b, ok := v.(int8); ok {
		return int(b), nil
	} else if b, ok := v.(int16); ok {
		return int(b), nil
	} else if b, ok := v.(int32); ok {
		return int(b), nil
	} else if b, ok := v.(int64); ok {
		return int(b), nil
	} else if b, ok := v.(uint8); ok {
		return int(b), nil
	} else if b, ok := v.(uint16); ok {
		return int(b), nil
	} else if b, ok := v.(uint32); ok {
		return int(b), nil
	} else if b, ok := v.(uint64); ok {
		return int(b), nil
	} else if b, ok := v.(float64); ok {
		return int(b), nil
	} else if b, ok := v.(float32); ok {
		return int(b), nil
	} else if v == nil {
		return 0, nil
	}

	return 0, errors.New("Value passed is not numeric")
}

// InterfaceToFloat64 performs type conversion to convert an interface into a float64.
func InterfaceToFloat64(v interface{}) (float64, error) {

	if b, ok := v.(float64); ok {
		return b, nil
	} else if b, ok := v.(float32); ok {
		return float64(b), nil
	} else if b, ok := v.(int); ok {
		return float64(b), nil
	} else if b, ok := v.(int8); ok {
		return float64(b), nil
	} else if b, ok := v.(int16); ok {
		return float64(b), nil
	} else if b, ok := v.(int32); ok {
		return float64(b), nil
	} else if b, ok := v.(int64); ok {
		return float64(b), nil
	} else if b, ok := v.(uint8); ok {
		return float64(b), nil
	} else if b, ok := v.(uint16); ok {
		return float64(b), nil
	} else if b, ok := v.(uint32); ok {
		return float64(b), nil
	} else if b, ok := v.(uint64); ok {
		return float64(b), nil
	} else if v == nil {
		return float64(0), nil
	}

	return 0, errors.New("Value passed is not numeric")
}

// InterfaceToString performs type conversion to convert an interface into a string.
func InterfaceToString(v interface{}) (string, error) {

	if b, ok := v.(string); ok {
		return b, nil
	} else if v == nil {
		return "", nil
	}

	return "", errors.New("Value passed is not string")
}

// InterfaceToNullableInt performs type conversion to convert an interface into an int.
func InterfaceToNullableInt(v interface{}) (null.Int, error) {

	if b, ok := v.(int); ok {
		return null.NewInt(int64(b), true), nil
	} else if b, ok := v.(int8); ok {
		return null.NewInt(int64(b), true), nil
	} else if b, ok := v.(int16); ok {
		return null.NewInt(int64(b), true), nil
	} else if b, ok := v.(int32); ok {
		return null.NewInt(int64(b), true), nil
	} else if b, ok := v.(int64); ok {
		return null.NewInt(b, true), nil
	} else if b, ok := v.(uint8); ok {
		return null.NewInt(int64(b), true), nil
	} else if b, ok := v.(uint16); ok {
		return null.NewInt(int64(b), true), nil
	} else if b, ok := v.(uint32); ok {
		return null.NewInt(int64(b), true), nil
	} else if b, ok := v.(uint64); ok {
		return null.NewInt(int64(b), true), nil
	} else if b, ok := v.(float64); ok {
		return null.NewInt(int64(b), true), nil
	} else if b, ok := v.(float32); ok {
		return null.NewInt(int64(b), true), nil
	} else if v == nil {
		return null.NewInt(int64(b), false), nil
	}

	return null.NewInt(int64(0), false), errors.New("Value passed is not numeric")
}

// InterfaceToNullableFloat performs type conversion to convert an interface into a float64.
func InterfaceToNullableFloat(v interface{}) (null.Float, error) {

	if b, ok := v.(float64); ok {
		return null.NewFloat(b, true), nil
	} else if b, ok := v.(float32); ok {
		return null.NewFloat(float64(b), true), nil
	} else if b, ok := v.(int); ok {
		return null.NewFloat(float64(b), true), nil
	} else if b, ok := v.(int8); ok {
		return null.NewFloat(float64(b), true), nil
	} else if b, ok := v.(int16); ok {
		return null.NewFloat(float64(b), true), nil
	} else if b, ok := v.(int32); ok {
		return null.NewFloat(float64(b), true), nil
	} else if b, ok := v.(int64); ok {
		return null.NewFloat(float64(b), true), nil
	} else if b, ok := v.(uint8); ok {
		return null.NewFloat(float64(b), true), nil
	} else if b, ok := v.(uint16); ok {
		return null.NewFloat(float64(b), true), nil
	} else if b, ok := v.(uint32); ok {
		return null.NewFloat(float64(b), true), nil
	} else if b, ok := v.(uint64); ok {
		return null.NewFloat(float64(b), true), nil
	} else if v == nil {
		return null.NewFloat(float64(0), false), nil
	}

	return null.NewFloat(float64(0), false), errors.New("Value passed is not numeric")
}

// InterfaceToNullableString performs type conversion to convert an interface into a string.
func InterfaceToNullableString(v interface{}) (null.String, error) {

	if b, ok := v.(string); ok {
		return null.NewString(b, true), nil
	} else if v == nil {
		return null.NewString("", false), nil
	}

	return null.NewString("", false), errors.New("Value passed is not a string")
}

// RowToStruct marshals results from MySQL utils which are returned in a ma[string]interface{} format into a structute using JSON as an intermediary
func RowToStruct(row map[string]interface{}, s interface{}) error {

	tempJSON, err := json.Marshal(row)
	if err != nil {
		return errors.New("Error marshaling to temporary JSON")
	}

	if err := json.Unmarshal(tempJSON, s); err != nil {
		return errors.New("Error unmarshaling temporary JSON to structure")
	}

	return nil
}

// StringToInterface performs type conversion to convert String to an Interface. Return nil if string is empty
func StringToInterface(s string) (ret interface{}) {

	if s == "" {
		return nil
	}

	ret = s

	return ret
}

// StructToMap converts a struct to a map[string]interface{} using the struct's tags
func StructToMap(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("StructToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			out[tagv] = v.Field(i).Interface()
		}
	}
	return out, nil
}

// RandomSecret generates a random n-character long string
func RandomSecret(n int) (string, error) {
	byteSlice := make([]byte, n)
	if _, err := rand.Read(byteSlice); err != nil {
		return "", err
	}

	return hex.EncodeToString(byteSlice), nil
}
