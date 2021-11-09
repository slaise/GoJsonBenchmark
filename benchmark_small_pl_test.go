package main

import (
	"encoding/json"
	"testing"

	"github.com/buger/jsonparser"
	gojson "github.com/goccy/go-json"
	"github.com/json-iterator/go"
	"github.com/mailru/easyjson/jlexer"
	"github.com/minio/simdjson-go"
	"github.com/valyala/fastjson"
)

// jsonparser
func BenchmarkJsonParserSmall(b *testing.B) {
	b.ReportAllocs()
	paths := [][]string{
		{"uuid"},
		{"tz"},
		{"ua"},
		{"st"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data SmallPayload

		jsonparser.EachKey(smallFixture, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
			switch idx {
			case 0:
				data.Uuid, _ = jsonparser.ParseString(value)
			case 1:
				v, _ := jsonparser.ParseInt(value)
				data.Tz = int(v)
			case 2:
				data.Ua, _ = jsonparser.ParseString(value)
			case 3:
				v, _ := jsonparser.ParseInt(value)
				data.St = int(v)
			}
		}, paths...)

	}
}

// go-json
func BenchmarkGoJsonSmall(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data SmallPayload
		gojson.Unmarshal(smallFixture, &data)
	}
}

// fastjson
func BenchmarkFastJsonSmall(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var p fastjson.Parser
		var data SmallPayload
		v, _ := p.ParseBytes(smallFixture)
		data.Uuid = string(v.GetStringBytes("uuid"))
		data.Tz = v.GetInt("tz")
		data.Ua = string(v.GetStringBytes("ua"))
		data.St = v.GetInt("st")
	}
}

// simdjson-go
func BenchmarkSimdJsonSmall(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d, _ := simdjson.Parse(smallFixture, nil)
		iter := d.Iter()
		var data SmallPayload
		obj, tmp, elem := &simdjson.Object{}, &simdjson.Iter{}, simdjson.Element{}
		typ := iter.Advance()
		switch typ {
		case simdjson.TypeRoot:
			typ, tmp, _ = iter.Root(&iter)

			if typ == simdjson.TypeObject {
				obj, _ = tmp.Object(obj)

				e := obj.FindKey("uuid", &elem)
				if e != nil && elem.Type == simdjson.TypeString {
					v, _ := elem.Iter.StringBytes()
					data.Uuid = string(v)
				}

				e = obj.FindKey("ua", &elem)
				if e != nil && elem.Type == simdjson.TypeString {
					v, _ := elem.Iter.StringBytes()
					data.Ua = string(v)
				}

				e = obj.FindKey("tz", &elem)
				if e != nil && elem.Type == simdjson.TypeInt {
					v, _ := elem.Iter.Int()
					data.Tz = int(v)
				}

				e = obj.FindKey("st", &elem)
				if e != nil && elem.Type == simdjson.TypeInt {
					v, _ := elem.Iter.Int()
					data.St = int(v)
				}
			}
		default:
			break
		}
	}
}

func BenchmarkJsnoiterPullSmall(b *testing.B) {
	b.ReportAllocs()
	iter := jsoniter.ParseBytes(jsoniter.ConfigDefault, smallFixture)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data SmallPayload
		iter.ResetBytes(smallFixture)
		for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
			switch field {
			case "uuid":
				data.Uuid = iter.ReadString()
			case "tz":
				data.Tz = iter.ReadInt()
			case "ua":
				data.Ua = iter.ReadString()
			case "st":
				data.St = iter.ReadInt()
			default:
				iter.Skip()
			}
		}
	}
}

func BenchmarkJsnoiterReflectSmall(b *testing.B) {
	iter := jsoniter.ParseBytes(jsoniter.ConfigDefault, smallFixture)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data SmallPayload
		iter.ResetBytes(smallFixture)
		jsoniter.Unmarshal(smallFixture, &data)
	}
}

/*
   encoding/json
*/
func BenchmarkEncodingJsonStructSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var data SmallPayload
		json.Unmarshal(smallFixture, &data)
	}
}

func BenchmarkEasyJsonSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		lexer := &jlexer.Lexer{Data: smallFixture}
		data := new(SmallPayload)
		data.UnmarshalEasyJSON(lexer)
	}
}
