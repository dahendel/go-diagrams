// Code generaTed by fileb0x at "2024-04-08 14:11:58.930861 -0500 CDT m=+0.733593751" from config file "b0x.yml" DO NOT EDIT.
// modified(2011-12-13 11:20:26 -0600 CST)
// original path: ../../assets/vmware/vm-poweron.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsVmwareVMPoweronPng is "assets/vmware/vm-poweron.png"
var FileAssetsVmwareVMPoweronPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x87\x01\x78\xfe\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x10\x00\x00\x00\x10\x08\x06\x00\x00\x00\x1f\xf3\xff\x61\x00\x00\x00\x04\x67\x41\x4d\x41\x00\x00\xb1\x8e\x7c\xfb\x51\x93\x00\x00\x00\x20\x63\x48\x52\x4d\x00\x00\x7a\x25\x00\x00\x80\x83\x00\x00\xf9\xff\x00\x00\x80\xe8\x00\x00\x75\x30\x00\x00\xea\x60\x00\x00\x3a\x97\x00\x00\x17\x6f\x97\xa9\x99\xd4\x00\x00\x01\x12\x49\x44\x41\x54\x78\x9c\x62\xfc\xff\xff\x3f\x03\x25\x00\x20\x80\x58\xd0\x05\x64\x26\x4b\xc1\x4d\x7c\x92\xfb\x8c\x91\x90\x01\x00\x01\x84\x61\x00\x08\xac\xca\x6d\x87\x1b\x46\xc8\x10\x80\x00\x62\xc2\x25\x71\xf4\xdd\x06\x86\xae\xf8\x64\x14\x17\x61\x03\x00\x01\x84\xd5\x80\x9f\xff\x7e\x30\xfc\xfe\xf3\x97\xe1\xc1\x8f\x8b\x04\x0d\x01\x08\x20\xec\x06\xfc\xfd\x0e\x34\xe0\x1f\x18\xdf\xfe\x72\x1e\xaf\x21\x00\x01\x84\xd5\x80\xef\x7f\xbe\x31\xfc\xfa\xfd\x17\x8e\xaf\xbe\x3f\x83\xd3\x10\x80\x00\xc2\x6e\xc0\xef\xef\x28\x06\x80\xf0\xf9\x57\x27\xb1\x1a\x02\x10\x40\x58\x0d\xf8\xf2\xf3\x2b\xc3\xaf\x5f\x7f\x31\xf0\xc9\x27\xc7\x18\x5e\x9c\x78\x81\xa2\x16\x20\x80\xb0\x46\x23\xd8\x00\xa0\xad\xe8\x60\x76\xd7\x21\x86\x3f\x4b\xff\xa1\x44\x2b\x40\x00\x61\x35\xe0\xf3\xf7\x2f\x18\x06\x2c\xec\x3f\x8a\xa1\x19\x04\x00\x02\x08\xab\x01\x9f\x7e\x7c\x01\xc7\x00\x0c\x2c\x9b\x7c\x1c\xab\x66\x10\x00\x08\x20\xac\x06\x7c\xfc\xfe\x19\x9c\x0e\x40\x60\xf5\xf4\x53\x38\x35\x83\x00\x40\x00\xe1\x30\xe0\x0b\x98\x5e\x3f\xeb\x0c\x5e\xcd\x20\x00\x10\x40\x58\x0d\x00\x69\x04\x01\x42\x9a\x41\x00\x20\x80\x18\x29\xcd\xce\x00\x01\x06\x00\xf8\xa7\xa1\xa7\x0d\x8c\x79\xcd\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\x71\x80\x2e\x80\x87\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsVmwareVMPoweronPng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/vmware/vm-poweron.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
