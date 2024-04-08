// Code generaTed by fileb0x at "2024-04-08 14:11:59.379985 -0500 CDT m=+1.182717668" from config file "b0x.yml" DO NOT EDIT.
// modified(2011-12-13 11:18:42 -0600 CST)
// original path: ../../assets/vmware/task.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsVmwareTaskPng is "assets/vmware/task.png"
var FileAssetsVmwareTaskPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\xd6\x01\x29\xfe\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x10\x00\x00\x00\x10\x08\x06\x00\x00\x00\x1f\xf3\xff\x61\x00\x00\x00\x04\x67\x41\x4d\x41\x00\x00\xb1\x8e\x7c\xfb\x51\x93\x00\x00\x00\x20\x63\x48\x52\x4d\x00\x00\x7a\x25\x00\x00\x80\x83\x00\x00\xf9\xff\x00\x00\x80\xe8\x00\x00\x75\x30\x00\x00\xea\x60\x00\x00\x3a\x97\x00\x00\x17\x6f\x97\xa9\x99\xd4\x00\x00\x01\x61\x49\x44\x41\x54\x78\x9c\x62\xfc\xff\xff\x3f\x03\x25\x00\x20\x80\x58\x60\x8c\x59\xb3\x66\xfd\x07\x62\xa2\x35\xa6\xa5\xa5\x81\x30\x23\x40\x00\x31\xc2\x5c\x60\x62\x62\xf2\x3f\xed\xcc\x19\xb8\x02\x4d\x20\xb6\x05\xe2\xc3\x40\x7c\x1d\x8b\x01\xb3\x4c\x4c\x18\xce\x9c\x39\xc3\x08\x10\x40\x2c\xc8\x82\xe9\x67\xa1\xa6\x8b\x00\xb1\x3c\x03\x03\xe3\x59\x46\x86\xff\xc6\xff\x19\xec\xce\x62\x1a\x60\x0c\xa5\x01\x02\x88\x09\x26\x10\x14\x14\xc4\xc0\x70\xf6\x2c\x58\xf3\x4c\xa8\x66\x06\x5c\x3e\x3a\x8b\x30\x11\x20\x80\xc0\x2e\x00\x2a\xfe\xff\xbf\xea\x3f\xc3\xc3\x25\x0f\x51\x34\xff\x9f\xf9\x9f\x21\x1d\x28\xc6\x70\xf8\x0d\xce\xb0\x00\x08\x20\x06\x70\x18\xa4\x31\xfc\x67\x38\x03\x62\xfe\x07\xd3\x20\x3e\x08\xfc\xfb\xf7\xff\xff\xfd\xfb\xf7\xc1\x34\x08\x9f\x3c\x79\x12\xce\x36\x36\x36\x06\x2b\x07\x08\x20\xb0\x17\x80\x36\x31\x82\x6c\x44\xb6\x19\x64\x2e\xd3\x39\x46\x06\x05\x05\x05\xdc\xb6\x03\x01\x40\x00\xc1\x03\x11\xa4\x89\x31\x9d\x11\x4c\x83\x03\xe7\x1c\x34\x0c\x66\x32\x30\x9c\x3e\x7d\x0a\xae\x01\x99\x0d\x02\x00\x01\x04\xf1\x02\xd4\xf9\x30\x67\xc3\xbc\x41\x8c\x17\x00\x02\x08\x12\x0b\x30\xe7\x23\xd9\xfc\x6f\x06\xc4\x25\xaf\x5e\xbd\xc2\xeb\x05\x80\x00\x62\x81\x85\x01\xd0\xf9\xff\x19\xd3\x10\x61\xf0\xe0\xc1\x03\xb8\x66\x7c\x5e\x00\x08\x20\xb0\x17\xa0\x98\xa0\xb3\xb1\x79\x01\x20\x80\xe0\x09\xa9\xad\xad\x8d\x68\x67\x23\x03\x80\x00\x82\xc7\x82\x88\x88\x08\x83\xa9\xa9\x09\xd1\x1a\x61\x00\x20\x80\x18\x29\xcd\xce\x00\x01\x06\x00\xb0\xe4\x07\xe7\xed\x7b\x43\x9e\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\x3d\xa6\xcf\x1c\xd6\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsVmwareTaskPng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/vmware/task.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
