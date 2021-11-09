package main

import (
	"encoding/json"
	"testing"

	"github.com/buger/jsonparser"
	gojson "github.com/goccy/go-json"
	"github.com/json-iterator/go"
	"github.com/mailru/easyjson/jlexer"
	"github.com/valyala/fastjson"
)

/*
   encoding/json
*/
func BenchmarkDecodeStdStructMedium(b *testing.B) {
	b.ReportAllocs()
	var data MediumPayload
	for i := 0; i < b.N; i++ {
		json.Unmarshal(mediumFixture, &data)
	}
}

// jsoniter
func BenchmarkDecodeJsoniterStructMedium(b *testing.B) {
	b.ReportAllocs()
	var data MediumPayload
	for i := 0; i < b.N; i++ {
		jsoniter.Unmarshal(mediumFixture, &data)
	}
}

// easyjson
func BenchmarkDecodeEasyJsonMedium(b *testing.B) {
	b.ReportAllocs()
	var data MediumPayload
	for i := 0; i < b.N; i++ {
		lexer := &jlexer.Lexer{Data: mediumFixture}
		data.UnmarshalEasyJSON(lexer)
	}
}

// jsonParser
func BenchmarkDecodeJsonParserMedium(b *testing.B) {
	b.ReportAllocs()
	paths := [][]string{
		{"person", "name", "fullName"},
		{"person", "github", "followers"},
	}
	for i := 0; i < b.N; i++ {
		var data = MediumPayload{
			Person: &CBPerson{
				Name: &CBName{},
				Github: &CBGithub{},
			},
		}
		jsonparser.EachKey(mediumFixture, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
			switch idx {
			case 0:
				data.Person.Name.FullName, _ = jsonparser.ParseString(value)
			case 1:
				v, _ := jsonparser.ParseInt(value)
				data.Person.Github.Followers = int(v)
			}
		}, paths...)
	}
}

// go-json
func BenchmarkDecodeGoJsonMedium(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data MediumPayload
		gojson.Unmarshal(mediumFixture, &data)
		_ = data.Person.Name.FullName
		_ = data.Person.Github.Followers
	}
}

// fastjson
func BenchmarkDecodeFastJsonMedium(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var p fastjson.Parser
		var data = MediumPayload{
			Person: &CBPerson{
				Name: &CBName{},
				Github: &CBGithub{},
			},
		}
		v, _ := p.ParseBytes(mediumFixture)
		data.Person.Name.FullName = string(v.GetStringBytes("person", "name", "fullName"))
		data.Person.Github.Followers = v.GetInt("person", "github", "followers")
	}
}


