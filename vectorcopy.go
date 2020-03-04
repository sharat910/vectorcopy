package vectorcopy

import "math/rand"

type Vector struct {
	Arr [48]uint32
}

func NewRandomVector() Vector {
	var v Vector
	for i:=0;i<10; i++{
		r := rand.Intn(48)
		v.Arr[r]++
	}
	return v
}

func MakeMap(v Vector) VectorMap {
	var vm VectorMap
	vm.Vector = make(map[uint32]uint32, 10)
	for i,count := range v.Arr {
		if count != 0 {
			vm.Vector[uint32(i)] = count
		}
	}
	return vm
}

func MakeArray (v Vector) VectorArray {
	var va VectorArray
	va.Vector = v.Arr[:]
	return va
}

func MakeList(v Vector) VectorList {
	var vl VectorList
	for i,count := range v.Arr {
		if count != 0 {
			vl.List = append(vl.List, &Elem{Index:uint32(i),Value:count})
		}
	}
	return vl
}
