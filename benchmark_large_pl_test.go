package main

import (
	"encoding/json"
	"testing"

	"github.com/buger/jsonparser"
	gojson "github.com/goccy/go-json"
	"github.com/json-iterator/go"
	"github.com/valyala/fastjson"
)

// jsonParser
func BenchmarkJsonParserLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		count := 0
		jsonparser.ArrayEach(largeFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			count++
		}, "topics", "topics")
	}
}

// go-json
func BenchmarkGoJsonLarge(b *testing.B)  {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data LargePayload
		gojson.Unmarshal(largeFixture, &data)
	}
}

// fastjson
func BenchmarkFastJsonLarge(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	count := 0
	for i := 0; i < b.N; i++ {
		var p fastjson.Parser
		v, _ := p.ParseBytes(largeFixture)
		vals := v.GetArray("topics")
		for i:= 0; i < len(vals); i++ {
			vv := vals[i].GetObject("topics")
			if vv != nil {
				_ = vv.String()
				count++
			}
		}
	}
}

// jsoniter
func BenchmarkJsoniterLarge(b *testing.B) {
	iter := jsoniter.ParseBytes(jsoniter.ConfigDefault, largeFixture)
	b.ReportAllocs()
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
func BenchmarkEncodingJsonLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		payload := &LargePayload{}
		json.Unmarshal(largeFixture, payload)
	}
}
