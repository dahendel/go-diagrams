// Code generaTed by fileb0x at "2024-04-08 14:11:58.965038 -0500 CDT m=+0.767770876" from config file "b0x.yml" DO NOT EDIT.
// modified(2011-12-13 11:18:08 -0600 CST)
// original path: ../../assets/vmware/refresh.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsVmwareRefreshPng is "assets/vmware/refresh.png"
var FileAssetsVmwareRefreshPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x0b\x02\xf4\xfd\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x10\x00\x00\x00\x10\x08\x06\x00\x00\x00\x1f\xf3\xff\x61\x00\x00\x00\x04\x67\x41\x4d\x41\x00\x00\xb1\x8e\x7c\xfb\x51\x93\x00\x00\x00\x20\x63\x48\x52\x4d\x00\x00\x7a\x25\x00\x00\x80\x83\x00\x00\xf9\xff\x00\x00\x80\xe8\x00\x00\x75\x30\x00\x00\xea\x60\x00\x00\x3a\x97\x00\x00\x17\x6f\x97\xa9\x99\xd4\x00\x00\x01\x96\x49\x44\x41\x54\x78\x9c\x62\xfc\xff\xff\x3f\x03\x08\xc8\x4c\x96\x82\x30\xd0\xc0\x8b\x13\x2f\x18\x24\x2c\x24\xc0\xf4\x9f\xa5\xff\x18\xd1\xe5\x01\x02\x88\x05\x99\xb3\x2a\xb7\x1d\xc3\x00\xbb\x13\x89\x0c\x4f\x72\x9f\x31\xca\x30\x48\xfd\x67\x89\x66\x02\x5b\x82\x6c\x10\x40\x00\xc1\x0d\x00\x2b\x02\xba\xa2\x38\xce\x12\xcc\x7f\xf3\xfe\x1b\xc3\x6b\x20\x46\x91\x67\x80\xc8\x83\x0c\x82\x19\x02\x10\x40\x4c\x30\x05\x20\xcd\x30\x17\x80\x34\x7f\xf9\xf6\x0b\x9b\x8f\x18\x8e\x5d\x78\xcc\xd0\x3d\x2d\x90\x01\xe6\x1a\x80\x00\x62\x42\xd6\x7c\xf2\xe3\x26\xbc\x9a\x41\xae\x38\x7e\xe9\x09\x83\x12\xaf\x2e\x5c\x0c\x20\x80\x98\x90\x15\xc0\x34\x7f\xff\xf9\x87\x61\x6a\xeb\x7e\xac\x86\x80\x80\x28\x93\x22\x43\xcb\x04\x3f\xb0\x2b\x00\x02\x88\x05\xd9\x76\x98\xe6\x79\x3d\x87\xc1\x01\x05\x73\x26\x36\xf0\xf3\xd7\x1f\x30\x0d\x10\x40\x28\xb1\x80\xac\x19\xc4\xc7\x16\x6d\x62\x8a\xa2\x0c\xaf\xff\xdd\x07\xab\x05\x01\x80\x00\x42\xf1\x02\xb2\x66\x6c\x00\xe4\xda\x29\x3e\x45\x0c\xf7\x3e\x5f\x66\xf8\xfe\xe3\x37\x58\x0c\x20\x80\x98\x70\x29\x86\x69\x80\x25\x30\x64\xaf\x3e\x7e\xf1\x09\xee\x02\x80\x00\xc2\x6b\x00\x08\x80\x34\x21\x6b\x7e\xf6\xea\x33\xd8\x76\x98\x6b\x01\x02\x88\x09\x14\x35\x61\x93\x2b\x51\xa2\x06\x1d\x14\x44\x5b\x80\x35\xbf\x7c\xfb\x15\x6c\x33\xcc\x76\x10\x00\x08\x20\xb8\x0b\x40\xfe\x0a\x4c\x33\x61\xc0\x16\xf2\xbf\x7e\xff\x01\x47\x31\xc8\x66\x10\x5e\x32\xf1\x18\x3c\xac\x00\x02\x08\x6c\x00\xc8\x15\xbd\x8b\x8e\x33\x58\x19\xc8\x62\x75\xc1\x8f\x5f\x7f\xe1\x36\x2f\xec\x3f\x8a\x12\xd0\x00\x01\x04\x77\x01\xcc\x10\x50\xce\x43\x77\x05\xcc\xe6\x1f\x40\x03\xd0\x5d\x09\x10\x40\x8c\xb0\xec\x8c\x0c\x40\x81\x86\x9c\x8d\xb1\x01\x98\x2b\x00\x02\x0c\x00\xf7\x45\xfe\xe3\x6d\xd9\x5c\x59\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\x5a\x42\x7f\x9f\x0b\x02\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsVmwareRefreshPng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/vmware/refresh.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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