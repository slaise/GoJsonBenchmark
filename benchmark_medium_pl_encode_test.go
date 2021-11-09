package main

import (
	"bytes"
	"encoding/json"
	"testing"

	gojson "github.com/goccy/go-json"
	"github.com/json-iterator/go"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

/*
   encoding/json
*/
func BenchmarkEncodeStdStructMedium(b *testing.B) {
	var data MediumPayload
	json.Unmarshal(mediumFixture, &data)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		json.Marshal(data)
	}
}

// jsoniter
func BenchmarkEncodeJsoniterStructMedium(b *testing.B) {
	var data MediumPayload
	jsoniter.Unmarshal(mediumFixture, &data)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsoniter.Marshal(data)
	}
}

// go-json
func BenchmarkEncodeGoJsonMedium(b *testing.B) {
	var data MediumPayload
	gojson.Unmarshal(mediumFixture, &data)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		gojson.Marshal(data)
	}
}

// easyjson
func BenchmarkEncodeEasyJsonMedium(b *testing.B) {
	var data MediumPayload
	lexer := &jlexer.Lexer{Data: mediumFixture}
	data.UnmarshalEasyJSON(lexer)
	b.ReportAllocs()
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		writer := &jwriter.Writer{}
		data.MarshalEasyJSON(writer)
		buf.Reset()
		writer.DumpTo(buf)
	}
}