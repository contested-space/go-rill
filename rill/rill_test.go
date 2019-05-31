package rill

import (
	"testing"
	"os"
	"path/filepath"
)

func TestAccOpenClose(t *testing.T) {
	os.RemoveAll("AccTest")
	os.Mkdir("." + string(filepath.Separator) + "test",0777)
	acc := AccOpen("test", 1000000)
	acc_file := filepath.Join("./test", "acc")
	if _, err := os.Stat(acc_file) ; os.IsNotExist(err) {
		t.Errorf("File %v not created by rill_acc_open", acc_file)
	}
	AccClose(acc)
}
