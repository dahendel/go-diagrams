// Code generaTed by fileb0x at "2024-04-08 14:11:58.69877 -0500 CDT m=+0.501503584" from config file "b0x.yml" DO NOT EDIT.
// modified(2011-12-13 11:16:18 -0600 CST)
// original path: ../../assets/vmware/host-yellow.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsVmwareHostYellowPng is "assets/vmware/host-yellow.png"
var FileAssetsVmwareHostYellowPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x9d\x01\x62\xfe\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x10\x00\x00\x00\x10\x08\x06\x00\x00\x00\x1f\xf3\xff\x61\x00\x00\x00\x04\x67\x41\x4d\x41\x00\x00\xb1\x8e\x7c\xfb\x51\x93\x00\x00\x00\x20\x63\x48\x52\x4d\x00\x00\x7a\x25\x00\x00\x80\x83\x00\x00\xf9\xff\x00\x00\x80\xe8\x00\x00\x75\x30\x00\x00\xea\x60\x00\x00\x3a\x97\x00\x00\x17\x6f\x97\xa9\x99\xd4\x00\x00\x01\x28\x49\x44\x41\x54\x78\x9c\x62\xfc\xff\xff\x3f\x03\x25\x00\x20\x80\x58\x90\x39\x4e\x4e\x4e\x04\x4d\xdb\xb7\x6f\x1f\x23\x32\x1f\x20\x80\x18\x40\x2e\x80\x61\x47\x47\xc7\xff\xb8\xc0\xb6\x6d\xdb\x40\x86\xff\x47\x56\x0f\xc2\x00\x01\xc4\x84\xcd\x16\x46\x46\x46\x0c\x8c\x0b\x00\x04\x10\x0b\x36\x41\xa0\x6d\x84\x7c\x02\x07\x00\x01\x84\xd5\x05\xa4\x00\x80\x00\xc2\xea\x02\x74\xb0\xae\xc9\x0b\x4c\xa7\x58\x60\xca\x01\x04\x10\xd1\x2e\x98\xdd\x84\x5d\x1c\x20\x80\x08\x1a\x00\xb2\x1d\xa4\xf9\xf2\x53\x88\x21\xa9\x96\x8c\x28\x51\x0d\x10\x40\x44\xb9\x00\xa4\x59\x2f\x11\x42\xa3\x03\x80\x00\xc2\x08\x83\xed\xdb\xb7\x33\x1c\x3a\x74\x08\xcc\x7e\x73\xa0\x03\x6e\x3b\x0c\xcc\x9e\x02\x71\xc5\xec\xe3\xff\xc1\x71\x0b\x10\x40\x18\x06\x78\x7a\x7a\x22\x9c\x0f\x34\x00\x03\x48\xa3\x72\x01\x02\x08\x23\x25\x82\x52\x5c\x45\x45\xc5\x7f\x60\x88\xff\xff\xbf\x0b\x8a\xcf\x30\xfc\xbf\x34\x1f\x48\x3f\x07\xe2\xd7\x10\x0c\x96\x07\xea\x01\x08\x20\x8c\x30\x00\xb9\xc0\xce\xce\x0e\x21\x20\x04\xf4\xc2\x65\x68\x18\xec\xc0\x74\x10\x40\x00\x61\xb8\x00\x28\x06\x49\xf3\x50\x9b\xc1\xf8\x39\xd4\x05\x20\xdb\x5f\x42\x5d\x72\x1f\xe2\x0a\x80\x00\x62\xc4\x96\x9d\xd3\xbd\x19\xff\xff\x7b\x87\xc5\x36\x34\xc0\x04\x74\x1d\x40\x80\x01\x00\x19\x9a\xb9\x41\x6c\xaa\x16\xba\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\x51\x09\x23\x96\x9d\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsVmwareHostYellowPng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/vmware/host-yellow.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
