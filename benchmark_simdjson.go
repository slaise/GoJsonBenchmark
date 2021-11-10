package main

import (
	"strconv"

	"github.com/minio/simdjson-go"
)

func findKey(iter simdjson.Iter, key string) (str string, err error) {

	obj, tmp, elem := &simdjson.Object{}, &simdjson.Iter{}, simdjson.Element{}

	for {
		typ := iter.Advance()

		switch typ {
		case simdjson.TypeRoot:
			if typ, tmp, err = iter.Root(tmp); err != nil {
				return
			}

			if typ == simdjson.TypeObject {
				if obj, err = tmp.Object(obj); err != nil {
					return
				}

				e := obj.FindKey(key, &elem)
				if e != nil && elem.Type == simdjson.TypeString {
					v, _ := elem.Iter.StringBytes()
					return string(v), nil
				}
				if e != nil && elem.Type == simdjson.TypeInt {
					v, _ := elem.Iter.Int()
					return strconv.FormatInt(v, 10), nil
				}
			}

		default:
			return
		}
	}
}
