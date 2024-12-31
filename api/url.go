package api

import (
	"fmt"
	"net/url"
	"path"
	"reflect"
	"slices"
	"strings"
)

type URL url.URL

func (u URL) WithExtraPath(p string) URL {
	u.Path = path.Join(u.Path, p)
	return u
}

func (u URL) String() string {
	return (*url.URL)(&u).String()
}

func compareStructField(a, b reflect.StructField) int {
	return strings.Compare(a.Name, b.Name)
}

func (u URL) WithGetParams(p any) URL {
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		v = v.Elem()
	}
	// fmt.Println("TYPE", t.Name())

	q := (*url.URL)(&u).Query()
	appenders := make([]string, 0)
	for _, field := range slices.SortedStableFunc(slices.Values(reflect.VisibleFields(t)), compareStructField) {
		// fmt.Println("FIELD", field.Name)
		jsonTag, ok := field.Tag.Lookup("json")
		if ok {
			tag, tagOpt := parseTag(jsonTag)
			fieldValue := v.FieldByIndex(field.Index)
			// handle pointers
			if fieldValue.Kind() == reflect.Pointer {
				fieldValue = fieldValue.Elem()
			}
			if fieldValue.Kind() == reflect.Slice {
				acc := make([]string, 0, fieldValue.Len())
				for i := range fieldValue.Len() {
					acc = append(acc, url.PathEscape(fmt.Sprint(fieldValue.Index(i).Interface())))
				}
				fmt.Println(">>", acc)
				fieldValue = reflect.New(reflect.TypeOf("")).Elem()
				fieldValue.SetString(strings.Join(acc, ","))
				if len(fieldValue.String()) > 0 {
					// u.RawQuery = u.RawQuery + fieldValue.String()
					appenders = append(appenders, fmt.Sprintf("%s=%s", tag, fieldValue.String()))
				}
				continue
			}
			// handle zero
			if !(fieldValue.IsZero() && tagOpt.Contains("omitempty")) {
				q.Set(tag, fmt.Sprint(fieldValue.Interface()))
			}
		}
	}
	u.RawQuery = strings.Join([]string{q.Encode(), strings.Join(appenders, "&")}, "&")
	return u
}
