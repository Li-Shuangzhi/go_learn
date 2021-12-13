package main

type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)

//(iota在每个const开头被重设为0)
const (
	c0 = iota
	c1 = 9
	c2
	c3 = iota + 1
)

//func main() {
//	var b int
//	b = int(Bool)
//	fmt.Println(b)
//	fmt.Println(c0, c1, c2, c3)
//	fmt.Println()
//}
