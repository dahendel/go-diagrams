// Code generaTed by fileb0x at "2024-04-08 14:11:59.516532 -0500 CDT m=+1.319264209" from config file "b0x.yml" DO NOT EDIT.
// modified(2011-12-13 11:19:52 -0600 CST)
// original path: ../../assets/vmware/virtualMachine-stuck.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsVmwareVirtualMachineStuckPng is "assets/vmware/virtualMachine-stuck.png"
var FileAssetsVmwareVirtualMachineStuckPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\xf9\x01\x06\xfe\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x10\x00\x00\x00\x10\x08\x06\x00\x00\x00\x1f\xf3\xff\x61\x00\x00\x00\x04\x67\x41\x4d\x41\x00\x00\xb1\x8e\x7c\xfb\x51\x93\x00\x00\x00\x20\x63\x48\x52\x4d\x00\x00\x7a\x25\x00\x00\x80\x83\x00\x00\xf9\xff\x00\x00\x80\xe8\x00\x00\x75\x30\x00\x00\xea\x60\x00\x00\x3a\x97\x00\x00\x17\x6f\x97\xa9\x99\xd4\x00\x00\x01\x84\x49\x44\x41\x54\x78\x9c\x62\xfc\x0f\x04\x0c\x14\x00\x80\x00\x62\x41\xe6\x28\xd5\xdf\xc1\xab\xf8\x5e\xa3\x0a\x86\x18\x40\x00\x31\xfc\x47\x02\x8a\x95\xe7\xff\x3f\xf8\xf5\x1f\x2b\x06\xc9\x2d\x5c\xb0\xe0\x3f\x3a\x00\x08\x20\x14\x17\x30\xfc\xf9\xcd\xf0\xeb\x1f\x0e\xeb\x81\x72\xd8\x00\x40\x00\x31\x22\x87\x81\x52\xfe\x71\x86\x4d\xed\x96\x60\xb6\x5f\x27\x1e\xef\xfc\xff\xc1\x70\xaf\x49\x07\xcc\x04\x08\x20\x54\x2f\x64\xed\xff\x7f\xe4\xed\x7f\x30\x56\xac\xbb\x8d\xe2\x85\xf4\xd5\x4f\xb1\x7a\x07\x20\x80\xc0\x2e\x50\xaa\xbb\xc2\xc0\xc0\xc8\x81\xd5\xb2\x90\xeb\xda\x58\xc5\x75\xbc\x67\x31\xc4\xc5\xc7\x33\x00\x04\x10\x03\x72\xe0\xf5\x5e\xfa\x0a\xa6\x1f\xfd\x87\xe0\xd2\x50\x36\x20\xf9\x00\x8c\x7f\x7e\x9e\x8e\x82\x21\x72\xff\xff\x03\x04\x10\x0b\x7a\xe0\xad\xbd\xf1\x8d\xe1\x3f\x13\x33\xc3\xcb\x46\x3e\x86\xae\x55\xb7\xe0\x36\xd6\x24\xe5\x83\xe9\x96\x79\x13\xc1\x74\xd7\xa4\x26\x86\xb2\x30\x76\x06\x80\x00\x82\x18\xf0\xf3\x0f\xc3\xaf\x3f\x0c\x0c\xee\x4a\x5c\x70\x0d\x0b\xd1\x9c\x0c\x32\xec\xd7\x97\xed\x60\x36\xdb\x97\x8f\x70\x71\x80\x00\x82\xba\xe0\x27\xc3\xf7\x3f\xb8\x03\x1d\x04\xca\xc2\xd4\x50\x5c\x00\x03\x00\x01\xc4\x04\x26\x7f\xfd\x66\xf8\xf9\x97\x01\x05\xa3\x83\x30\x67\x51\x0c\xdb\x41\x00\x20\x80\xa0\x06\xfc\x61\xf8\xfe\x8b\x01\x05\xbb\x4d\xfa\x09\xb6\x15\xe4\x6c\x98\xd3\x51\x5c\x94\x57\x07\xf4\xd6\x4f\x06\x80\x00\x82\x44\x63\xf4\x66\x9c\x4e\x0f\xf9\x1d\x02\xa6\x41\x81\x06\xd3\x08\xe6\x03\x35\x83\x00\x40\x00\xc1\x53\xe2\xa2\x85\xe8\xc1\x86\x0a\xae\x6c\x4d\x83\x84\x3c\xd4\x66\x18\x00\x08\x20\x94\xa4\x4c\x08\x80\xa2\x0d\x06\x60\x86\x00\x04\x18\x00\xd1\x95\x2e\xf2\x33\x52\xe8\x2c\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\x66\xcf\x1b\x15\xf9\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsVmwareVirtualMachineStuckPng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/vmware/virtualMachine-stuck.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
