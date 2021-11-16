package main

import (
	"encoding/json"
	"sync"
	"testing"
)

var buf, _ = json.Marshal(User{Name: "Test", ID: 1})

var userPool = sync.Pool{
	New: func() interface{} {
		return &User{Name: "Test", ID: 1}
	},
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		user := &User{}
		json.Unmarshal(buf, user)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		user := userPool.Get().(*User)
		json.Unmarshal(buf, user)
		userPool.Put(user)
	}
}
