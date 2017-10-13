package validator

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"strings"
)

// 1. int, bool 类型声明的时候必须用指针
// 2. 如果是嵌入对象，那么必须是required，不然会panic
// 4. 类似链表嵌套
// set default value in array/slice is not support now

func ValidateJson(in []byte, v interface{}) error {
	err := json.Unmarshal(in, v)
	if err != nil {
		return fmt.Errorf("Invalid request: malformed %s", err)
	}

	err = ValidateParameters(v)
	if err != nil {
		return err
	}

	return nil
}

func ValidateParameters(in interface{}) (err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//Indirect returns the value that v points to. If v is a nil pointer,
	//Indirect returns a zero Value. If v is not a pointer, Indirect returns v.
	//Using this we don't need to care about if the interface is point to v or *v -gs
	v := reflect.ValueOf(in).Elem()
	t := reflect.TypeOf(in).Elem()

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		sv := v.FieldByName(sf.Name)
		if isRequired, ok := sf.Tag.Lookup("is_required"); ok && isRequired == "true" {
			// fmt.Printf("%s\n", sf.Type.Kind())
			switch sf.Type.Kind() {
			case reflect.String:
				if sv.String() == "" {
					return fmt.Errorf("Invalid request: missing %s", sf.Name)
				}
			case reflect.Int, reflect.Int64:
				//FIXME: 0 is meaningful here
				if sv.Int() == 0 {
					return fmt.Errorf("Invalid request: missing %s", sf.Name)
				}
			case reflect.Struct:
				err = ValidateParameters(sv.Interface())
				if err != nil {
					return err
				}
			case reflect.Ptr:
				if sv.IsNil() {
					return fmt.Errorf("Invalid request: missing %s", sf.Name)
				}
				if sv.Elem().Kind() == reflect.Struct {
					// fmt.Printf("%s:%s\n", sv, reflect.ValueOf(iv).Type())
					err = ValidateParameters(sv.Interface())
					if err != nil {
						return err
					}
				}
			case reflect.Slice, reflect.Array:
				if sv.Len() == 0 {
					return fmt.Errorf("Invalid request: missing %s", sf.Name)
				}
			}
		} else {
			switch sf.Type.Kind() {
			case reflect.String:
				if defaultV, ok := sf.Tag.Lookup("default"); ok && sv.String() == "" {
					sv.SetString(defaultV)
				}
			case reflect.Int, reflect.Int64:
				if defaultV, ok := sf.Tag.Lookup("default"); ok && sv.Int() == 0 {
					tempV, _ := strconv.Atoi(defaultV)
					sv.SetInt(int64(tempV))
				}
			case reflect.Struct:
				// fmt.Printf("TEST %s\n",)
				err = ValidateParameters(sv.Interface())
				if err != nil {
					return err
				}
			case reflect.Ptr:
				// fmt.Printf("%s\n", sv)
				tx := reflect.TypeOf(sv.Interface())
				// fmt.Printf("%s\n", tx.Elem().Kind())
				switch tx.Elem().Kind() {
				case reflect.Bool:
					if defaultV, ok := sf.Tag.Lookup("default"); ok && sv.IsNil() {
						tempV, _ := strconv.ParseBool(defaultV)
						newV := reflect.ValueOf(&tempV)
						sv.Set(newV)
					}
				case reflect.Int, reflect.Int64:
					if defaultV, ok := sf.Tag.Lookup("default"); ok && sv.IsNil() {
						tempV, _ := strconv.Atoi(defaultV)
						fmt.Printf("%s\n", sv)
						newV := reflect.ValueOf(&tempV)
						sv.Set(newV)
					}
				case reflect.Struct:
					if sv.IsNil() {
						newStruct := reflect.New(tx.Elem())
						sv.Set(newStruct)
					}
					// fmt.Printf("%s:%s\n", sv, reflect.ValueOf(sv.Interface()).Type())
					err = ValidateParameters(sv.Interface())
					if err != nil {
						return err
					}
				}
			}
		}

		if maxLen, ok := sf.Tag.Lookup("max_len"); ok {
			maxLenInt, _ := strconv.Atoi(maxLen)
			switch sf.Type.Kind() {
			case reflect.String:
				s := sv.String()
				if len(s) > maxLenInt {
					return fmt.Errorf("%s exceed the max len %d", sf.Name, maxLenInt)
				}
			case reflect.Slice, reflect.Array:
				if sv.Len() > maxLenInt {
					return fmt.Errorf("%s exceed the max len %d", sf.Name, maxLenInt)
				}
			}
		}

		if max, ok := sf.Tag.Lookup("max"); ok {
			switch sf.Type.Kind() {
			case reflect.Int, reflect.Int64:
				maxInt, _ := strconv.Atoi(max)
				s := sv.Int()
				if s > int64(maxInt) {
					return fmt.Errorf("%s exceed the max %d", sf.Name, maxInt)
				}
			case reflect.Ptr:
				tx := reflect.TypeOf(sv.Interface())
				switch tx.Elem().Kind() {
				case reflect.Int, reflect.Int64:
					maxInt, _ := strconv.Atoi(max)
					s := sv.Elem().Int()
					if s > int64(maxInt) {
						return fmt.Errorf("%s exceed the max %d", sf.Name, maxInt)
					}
				}
			}
		}

		//FIXME: need to check if it's a empty value
		if min, ok := sf.Tag.Lookup("min"); ok {
			switch sf.Type.Kind() {
			case reflect.Int, reflect.Int64:
				minInt, _ := strconv.Atoi(min)
				s := sv.Int()
				if s != 0 && s < int64(minInt) {
					return fmt.Errorf("%s exceed the min %d", sf.Name, minInt)
				}
			case reflect.Ptr:
				tx := reflect.TypeOf(sv.Interface())
				switch tx.Elem().Kind() {
				case reflect.Int, reflect.Int64:
					minInt, _ := strconv.Atoi(min)
					s := sv.Elem().Int()
					if s < int64(minInt) {
						return fmt.Errorf("%s exceed the min %d", sf.Name, minInt)
					}
				}
			}
		}

		if enum, ok := sf.Tag.Lookup("enum"); ok {
			if sf.Type.Kind() == reflect.String {
				enumList := strings.Split(enum, ",")
				s := sv.String()
				isIn := false
				for _, v := range enumList {
					if s == v {
						isIn = true
					}
				}
				if !isIn {
					return fmt.Errorf("%s contain a unsupport type '%s'", sf.Name, s)
				}
			}
		}
	}
	return
}
