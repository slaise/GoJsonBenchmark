package main

import (
	"encoding/json"
	"testing"

	"github.com/buger/jsonparser"
	gojson "github.com/goccy/go-json"
	"github.com/json-iterator/go"
)

// jsonParser
func BenchmarkJsonParserLargePl(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(largeFixture)))

	for i := 0; i < b.N; i++ {
		count := 0
		jsonparser.ArrayEach(largeFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			count++
		}, "users", "username", "avatar_template")
	}
}

// go-json
func BenchmarkGoJsonLargePl(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(largeFixture)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data LargePayload
		gojson.Unmarshal(largeFixture, &data)
	}
}

// jsoniter
func BenchmarkJsoniterLargePl(b *testing.B) {
	iter := jsoniter.ParseBytes(jsoniter.ConfigDefault, largeFixture)
	b.ReportAllocs()
	b.SetBytes(int64(len(largeFixture)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iter.ResetBytes(largeFixture)
		count := 0
		for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
			if "topics" != field {
				iter.Skip()
				continue
			}
			for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
				if "topics" != field {
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
func BenchmarkEncodingJsonLargePl(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(int64(len(largeFixture)))
	for i := 0; i < b.N; i++ {
		payload := &LargePayload{}
		json.Unmarshal(largeFixture, payload)
	}
}
