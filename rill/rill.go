package rill

// #cgo LDFLAGS: /home/fabien.lf/code/rill/build/librill.a
// #include </home/fabien.lf/code/rill/src/rill.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"

type Acc struct {
	acc *C.struct_rill_acc
}

func AccOpen(path string, capacity int) Acc {
	acc := C.rill_acc_open(C.CString(path), C.size_t(capacity))
	return Acc{acc}
}

func AccClose(acc Acc) {
	C.rill_acc_close((*C.struct_rill_acc)(acc.acc))
}
