package vectorcopy

import (
	"github.com/golang/protobuf/proto"
	"testing"
)

var rvm VectorMap
var rva VectorArray
var rvl VectorList

func BenchmarkMakeMap(b *testing.B) {
	v := NewRandomVector()
	var vm VectorMap
	for n := 0; n < b.N; n++ {
		vm = MakeMap(v)
	}
	rvm = vm
}

func BenchmarkMakeList(b *testing.B) {
	v := NewRandomVector()
	var vl VectorList
	for n := 0; n < b.N; n++ {
		vl = MakeList(v)
	}
	rvl = vl
}

func BenchmarkMakeArray(b *testing.B) {
	v := NewRandomVector()
	var va VectorArray
	for n := 0; n < b.N; n++ {
		va = MakeArray(v)
	}
	rva = va
}

var gb []byte

func BenchmarkMarshalMap(b *testing.B) {
	v := NewRandomVector()
	vm := MakeMap(v)
	var lb []byte
	for n := 0; n < b.N; n++ {
		lb, _ = proto.Marshal(&vm)
	}
	gb = lb
}

func BenchmarkMarshalArray(b *testing.B) {
	v := NewRandomVector()
	vm := MakeArray(v)
	var lb []byte
	for n := 0; n < b.N; n++ {
		lb, _ = proto.Marshal(&vm)
	}
	gb = lb
}

func BenchmarkMarshalList(b *testing.B) {
	v := NewRandomVector()
	vm := MakeList(v)
	var lb []byte
	for n := 0; n < b.N; n++ {
		lb, _ = proto.Marshal(&vm)
	}
	gb = lb
}