package api

import (
	"fmt"
	"net/url"
	"path"
	"reflect"
)

type URL url.URL

func (u URL) WithExtraPath(p string) URL {
	u.Path = path.Join(u.Path, p)
	return u
}

func (u URL) String() string {
	return (*url.URL)(&u).String()
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
	for _, field := range reflect.VisibleFields(t) {
		// fmt.Println("FIELD", field.Name)
		jsonTag, ok := field.Tag.Lookup("json")
		if ok {
			tag, tagOpt := parseTag(jsonTag)
			fieldValue := v.FieldByIndex(field.Index)
			if !(fieldValue.IsZero() && tagOpt.Contains("omitempty")) {
				q.Set(tag, fmt.Sprint(fieldValue.Interface()))
			}
		}
	}
	u.RawQuery = q.Encode()
	return u
}
