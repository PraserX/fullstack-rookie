package helpers

import (
	"html"
	"reflect"
	"regexp"
)

// ContainsForbiddenChars checking string for forbidden characters. So if there
// is some forbidden character at string, this function returns true and issued
// string, otherwise it will return false and empty string.
//
// Be aware! Function only check string or struct with its top level strings.
// So if struct contains some arrays of structures, then it will not be checked.
func ContainsForbiddenChars(inspect interface{}) (bool, string) {
	var val reflect.Value
	var notValid = regexp.MustCompile(`^.*('|"|<|>|&|;).*$`).MatchString

	if reflect.TypeOf(inspect).Kind() != reflect.Ptr {
		val = reflect.ValueOf(inspect)
	} else {
		val = reflect.Indirect(reflect.ValueOf(inspect))
	}

	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			ival := val.Field(i).Interface()
			if reflect.TypeOf(ival).Kind() == reflect.String {
				if notValid(val.Field(i).String()) {
					return true, val.Field(i).String()
				}
			}
		}
	} else if val.Kind() == reflect.String {
		if notValid(val.String()) {
			return true, val.String()
		}
	}

	return false, ""
}

// Sanitize provides sanitization of string. So if there is some unwanted
// character at string, this function sanitize the text. The values has to be
// editable. If the input is not a pointer, function a cannot sanitize content.
//
// Be aware! Function only sanitize string or struct with its top level strings.
// So if struct contains some arrays of structures, then it will be skipped.
func Sanitize(inspect interface{}) {
	if reflect.TypeOf(inspect).Kind() != reflect.Ptr {
		return
	}

	indirectValue := reflect.Indirect(reflect.ValueOf(inspect))
	if indirectValue.Kind() == reflect.Struct {
		for i := 0; i < indirectValue.NumField(); i++ {
			ival := indirectValue.Field(i).Interface()
			if reflect.TypeOf(ival).Kind() == reflect.String && indirectValue.Field(i).CanSet() {
				indirectValue.Field(i).SetString(html.EscapeString(indirectValue.Field(i).String()))
			}
		}
	} else if indirectValue.Kind() == reflect.String && indirectValue.CanSet() {
		indirectValue.SetString(html.EscapeString(indirectValue.String()))
	}
}
