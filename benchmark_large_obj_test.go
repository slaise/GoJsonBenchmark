package main

import (
	"encoding/json"
	"testing"

	"github.com/buger/jsonparser"
	gojson "github.com/goccy/go-json"
	"github.com/json-iterator/go"
)

// jsonParser
func BenchmarkJsonParserLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(largeStructText)))

	paths := [][]string{
		{"statuses", "user", "entities", "default_profile"},
		{"search_metadata", "completed_in"},
	}

	for i := 0; i < b.N; i++ {
		jsonparser.EachKey(largeStructText, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
			switch idx {
			case 0:
				jsonparser.ParseString(value)
			case 1:
				jsonparser.ParseFloat(value)
			}
		}, paths...)
	}
}

func BenchmarkEasyJsonLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(largeStructText)))
	for i := 0; i < b.N; i++ {
		var s LargeStruct
		err := s.UnmarshalJSON(largeStructText)
		if err != nil {
			b.Error(err)
		}
	}
}

// go-json
func BenchmarkGoJsonLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(largeStructText)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data LargeStruct
		gojson.Unmarshal(largeStructText, &data)
	}
}

// jsoniter
func BenchmarkJsoniterLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(largeStructText)))
	iter := jsoniter.ParseBytes(jsoniter.ConfigDefault, largeStructText)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iter.ResetBytes(largeStructText)
		count := 0
		for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
			if "in_reply_to_screen_name" != field {
				iter.Skip()
				continue
			}
			for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
				if "in_reply_to_screen_name" != field {
					iter.Skip()
					continue
				}
				for iter.ReadArray() {
					iter.Skip()
					count++
				}
				break
			}
			break
		}
	}
}

// encode/json
func BenchmarkEncodingJsonLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(largeStructText)))

	for i := 0; i < b.N; i++ {
		payload := &LargeStruct{}
		json.Unmarshal(largeStructText, payload)
	}
}
