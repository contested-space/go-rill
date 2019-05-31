package rill

// #cgo LDFLAGS: /home/fabien.lf/code/rill/build/librill.a
// #include </home/fabien.lf/code/rill/src/rill.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"

type Acc struct {
	acc *C.struct_rill_acc
}

type Rows struct {
	len C.size_t;
	data *C.struct_rill_row
}

type Row struct {
	a C.rill_val_t;
	b C.rill_val_t
}

func AccOpen(path string, capacity int) Acc {
	acc := C.rill_acc_open(C.CString(path), C.size_t(capacity))
	return Acc{acc}
}

func AccClose(acc Acc) {
	C.rill_acc_close((*C.struct_rill_acc)(acc.acc))
}

func AccIngest(acc Acc, val_a int, val_b int) {
	C.rill_acc_ingest(
		(*C.struct_rill_acc)(acc.acc),
		C.rill_val_t(val_a),
		C.rill_val_t(val_b))
}

func AccWrite(acc Acc, file string, ts int) {
	C.rill_acc_write(
		(*C.struct_rill_acc)(acc.acc),
		C.CString(file),
		C.rill_ts_t(ts))
}

func Rotate(dir string, now_ts int, expire_time int) {
	C.rill_rotate(C.CString(dir), C.rill_ts_t(now_ts), C.rill_ts_t(expire_time))
}
