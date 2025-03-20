package rt

import (
	"unsafe"
)

type GoMapIterator struct {
	K     unsafe.Pointer
	V     unsafe.Pointer
	T     *GoMapType
	It    unsafe.Pointer
}
