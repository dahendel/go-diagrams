// Code generaTed by fileb0x at "2024-04-08 14:11:59.244698 -0500 CDT m=+1.047430084" from config file "b0x.yml" DO NOT EDIT.
// modified(2011-12-13 11:20:40 -0600 CST)
// original path: ../../assets/vmware/vm-suspend.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsVmwareVMSuspendPng is "assets/vmware/vm-suspend.png"
var FileAssetsVmwareVMSuspendPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xea\x0c\xf0\x73\xe7\xe5\x92\xe2\x62\x60\x60\xe0\xf5\xf4\x70\x09\x62\x60\x60\x10\x00\x61\x0e\x36\x06\x06\x06\xf9\xcf\xff\x13\x19\x18\x18\x58\xd2\x1d\x7d\x1d\x19\x18\x36\xf6\xd5\xfc\x0e\x9c\xcc\xc0\xc0\xa0\x90\xec\x11\xe4\xcb\xc0\x50\xa5\xca\xc0\xd0\xd0\xcc\xc0\xf0\xf3\x3f\x03\x43\xc3\x0b\x06\x86\x52\x03\x06\x86\x57\x09\x0c\x0c\x56\xd3\x19\x18\xc4\xf3\xa7\xaf\x9c\x79\x85\x81\x81\x61\xbb\xa7\x8b\x63\x48\xc5\x9c\xa4\x3f\xff\xff\xdb\x33\xab\x32\x28\x34\xcc\xd0\xb8\xc4\xcb\xc2\xc0\x78\x84\x29\x79\x8e\x6f\xfa\x91\x82\xae\x8f\x67\xfe\xf7\xb4\x3c\x61\x60\xea\x60\x0d\x35\xe0\x3d\xe7\x67\x3e\xf3\xa8\xc1\x64\x0e\xcb\x06\x06\x27\xc3\x86\x87\x97\x8d\xd9\xda\x56\x47\x6f\xd8\x2d\xda\x2e\xe3\xc0\xb0\x50\xea\xc3\x5f\x6e\xe6\xc3\xf1\x6f\x7e\x15\x9e\x64\x64\x11\x58\xc5\x78\xdf\xce\xd8\xe0\xcf\xb1\x44\x06\x3e\xcb\x06\x86\x43\x26\x0f\xf8\xe5\x98\x37\x9c\x00\x4b\xcc\xff\xb8\x83\x01\xa4\x23\x30\x87\x01\xac\xed\x23\x37\xdf\xdd\x1d\x4f\x19\x98\x38\xae\x32\xc4\xfc\x3b\xae\xcc\x2e\xe0\xb0\x82\xed\x41\x91\x9d\x1c\xd8\x16\x85\x5a\xa5\xbc\xbf\x3b\x40\x0e\x98\x78\xe9\x0e\x87\x80\x83\xc8\x73\xb7\x06\x06\x66\x86\x15\xd5\x8e\x8d\x73\xe7\x6c\x09\x61\x60\x60\x60\xf0\x74\xf5\x73\x59\xe7\x94\xd0\x04\x08\x00\x00\xff\xff\xf9\xd1\x83\xe0\x2c\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsVmwareVMSuspendPng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/vmware/vm-suspend.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
