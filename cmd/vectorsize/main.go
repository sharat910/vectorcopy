package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sharat910/vectorcopy"
)

func main() {
	rv := vectorcopy.NewRandomVector()
	va := vectorcopy.MakeArray(rv)
	vl := vectorcopy.MakeList(rv)
	vm := vectorcopy.MakeMap(rv)
	ba,_ := proto.Marshal(&va)
	bl,_ := proto.Marshal(&vl)
	bm,_ := proto.Marshal(&vm)
	fmt.Println(len(ba), len(bl), len(bm))
}
