// Code generaTed by fileb0x at "2024-04-08 14:11:59.094685 -0500 CDT m=+0.897418084" from config file "b0x.yml" DO NOT EDIT.
// modified(2011-12-13 11:20:46 -0600 CST)
// original path: ../../assets/vmware/warning.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsVmwareWarningPng is "assets/vmware/warning.png"
var FileAssetsVmwareWarningPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x40\x01\xbf\xfe\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x10\x00\x00\x00\x10\x08\x06\x00\x00\x00\x1f\xf3\xff\x61\x00\x00\x00\x04\x67\x41\x4d\x41\x00\x00\xb1\x8e\x7c\xfb\x51\x93\x00\x00\x00\x20\x63\x48\x52\x4d\x00\x00\x7a\x25\x00\x00\x80\x83\x00\x00\xf9\xff\x00\x00\x80\xe8\x00\x00\x75\x30\x00\x00\xea\x60\x00\x00\x3a\x97\x00\x00\x17\x6f\x97\xa9\x99\xd4\x00\x00\x00\xcb\x49\x44\x41\x54\x78\x9c\x62\xfc\xff\xff\x3f\x03\x25\x00\x20\x80\x98\x28\xd2\x0d\x04\x00\x01\x44\xb1\x01\x00\x01\xc4\x82\x4f\x32\xd5\x92\x11\xee\xbf\xd9\xc7\xff\x33\x62\x53\x03\x10\x40\x04\x5d\x30\xbb\x09\xbf\x3c\x40\x00\x31\x80\x02\x11\x1b\x4e\xb1\x00\x52\xbb\x18\xfe\x5f\x9a\x0f\xa1\xc1\x7c\x2c\xea\x00\x02\x08\xaf\x0b\x2e\x3f\x65\x60\xd0\x4b\x84\xd0\xb8\x00\x40\x00\x11\xb4\x1d\xa8\x06\xe2\x8a\x33\xd8\x5d\x01\x10\x40\xc4\xc7\x82\x34\x76\x61\x80\x00\xc2\x69\x3b\x18\x9f\x81\xda\xfe\x1c\x88\x5f\x43\x30\xba\x2b\x00\x02\x08\xb7\x0b\x84\x80\x7e\xbf\x0c\x0d\x83\x1d\xb8\x1d\x06\x10\x40\xd8\x6d\x3f\x03\xc5\xcf\xa1\x2e\x00\xd9\xfe\x12\xea\x92\xfb\xa8\xae\x00\x08\x20\x94\x84\xc4\x04\xb4\x35\xb5\x0e\x8b\x2d\x33\x51\xb9\x20\x75\x30\x00\x10\x40\x8c\x94\x66\x26\x80\x00\xa2\x38\x2f\x00\x04\x10\xc5\x06\x00\x04\x18\x00\x8c\xb0\xc1\x0f\xa0\x3c\x55\xb9\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\x77\xd9\x16\x01\x40\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsVmwareWarningPng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/vmware/warning.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
