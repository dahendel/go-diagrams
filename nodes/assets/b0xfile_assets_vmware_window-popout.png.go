// Code generaTed by fileb0x at "2024-04-08 14:11:59.21975 -0500 CDT m=+1.022482876" from config file "b0x.yml" DO NOT EDIT.
// modified(2011-12-13 11:20:50 -0600 CST)
// original path: ../../assets/vmware/window-popout.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsVmwareWindowPopoutPng is "assets/vmware/window-popout.png"
var FileAssetsVmwareWindowPopoutPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x9e\x01\x61\xfe\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x10\x00\x00\x00\x10\x08\x06\x00\x00\x00\x1f\xf3\xff\x61\x00\x00\x00\x04\x67\x41\x4d\x41\x00\x00\xb1\x8e\x7c\xfb\x51\x93\x00\x00\x00\x20\x63\x48\x52\x4d\x00\x00\x7a\x25\x00\x00\x80\x83\x00\x00\xf9\xff\x00\x00\x80\xe8\x00\x00\x75\x30\x00\x00\xea\x60\x00\x00\x3a\x97\x00\x00\x17\x6f\x97\xa9\x99\xd4\x00\x00\x01\x29\x49\x44\x41\x54\x78\x9c\x62\xfc\xff\xff\x3f\x03\x25\x00\x20\x80\x98\x28\xd2\x0d\x04\x00\x01\xc4\x42\x48\xc1\xc6\x8d\x1b\xff\x07\x3c\x08\xc0\x2a\xb7\x41\x61\x03\x03\x40\x00\x31\xce\x9c\x39\xf3\xff\xac\x59\xb3\x30\x24\x8d\x8d\x8d\x19\x80\x72\x8c\x8c\x13\x19\xff\x1f\xcd\xd9\x8c\xd5\x00\xeb\x29\xbe\x0c\x00\x01\xc4\xb2\x79\xf3\x66\x86\x33\x67\xce\xa0\x48\xdc\xbb\x77\x8f\x21\x3f\x3f\x1f\xcc\xfe\x9f\xff\x1f\x6c\x48\x55\x42\x16\x56\x43\x00\x02\x88\x45\x5c\x5c\x1c\xa7\xf3\x7d\x7d\x7d\xc1\x21\x7c\x77\xe2\x5d\x06\xe5\x05\xca\x0c\xd8\x0c\x01\x08\x20\x16\x6c\xb1\xa0\xa4\xa4\xc4\x00\x72\x19\x0c\x00\x5d\x00\xd7\x7c\xee\xca\x75\x06\x23\x1d\x4d\xb8\x1c\x40\x00\x31\xfd\xfb\xf7\x0f\xa7\x0b\xb0\x69\xde\x71\x66\x3f\x98\x86\x01\x80\x00\x62\x21\x64\x00\x30\x0c\xc0\x86\x78\x98\x38\x82\x35\xc3\xf8\x30\x00\x10\x40\x0c\x71\x71\x71\xff\x89\x01\x0c\x13\x18\xe0\xec\xbb\x77\xef\x82\xf9\x1b\x36\x6c\xf8\x0f\x10\x40\x4c\xff\xff\x13\x97\x96\x40\x36\x23\x83\x99\x9c\x33\x19\xfc\xfd\xfd\x19\x01\x02\x88\xe9\xef\x5f\x82\x69\x09\x2b\x80\xa5\x1d\x80\x00\xa2\x38\x29\x03\x04\x10\xd3\x9b\x37\x2f\x28\x32\x00\x20\x80\x58\x82\x83\x7d\x19\x4c\x4c\x4c\x48\xd6\x68\x60\x60\x00\xa6\x01\x02\x88\x91\xd2\xec\x0c\x10\x60\x00\x26\x1e\x95\x4f\x07\x8e\x6d\x69\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\x7b\x35\x7b\x00\x9e\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsVmwareWindowPopoutPng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/vmware/window-popout.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
