// Code generaTed by fileb0x at "2024-04-08 14:11:59.490286 -0500 CDT m=+1.293018251" from config file "b0x.yml" DO NOT EDIT.
// modified(2011-12-13 11:15:56 -0600 CST)
// original path: ../../assets/vmware/healthRed.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsVmwareHealthRedPng is "assets/vmware/healthRed.png"
var FileAssetsVmwareHealthRedPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x8e\x01\x71\xfe\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x2c\x00\x00\x00\x13\x08\x02\x00\x00\x00\x64\xca\x78\x2e\x00\x00\x00\x01\x73\x52\x47\x42\x00\xae\xce\x1c\xe9\x00\x00\x00\x04\x67\x41\x4d\x41\x00\x00\xb1\x8f\x0b\xfc\x61\x05\x00\x00\x00\x20\x63\x48\x52\x4d\x00\x00\x7a\x26\x00\x00\x80\x84\x00\x00\xfa\x00\x00\x00\x80\xe8\x00\x00\x75\x30\x00\x00\xea\x60\x00\x00\x3a\x98\x00\x00\x17\x70\x9c\xba\x51\x3c\x00\x00\x01\x0c\x49\x44\x41\x54\x48\x4b\x63\xfc\xff\xff\x3f\xc3\x80\x03\xa0\x23\x06\x1c\x30\x0c\xb8\x0b\x40\x51\x81\xe9\x88\x5d\x3c\x3c\xc8\x08\x97\x2b\x03\x50\x01\x25\xca\x50\x1c\x31\x23\x37\x17\x68\xfd\xff\x35\x8e\xc8\x08\x28\x02\x14\x47\xb6\xa3\x77\xce\x06\xa0\x03\xd0\x6c\x05\x8a\x00\xc5\xc9\x50\x86\x1e\x12\x50\x17\xb4\xaa\xfd\xf7\x63\x80\x22\x20\x7b\x8d\x23\x48\x1c\x09\x60\xba\x00\x22\x89\x26\x4e\xa4\x32\x14\x47\x40\x5d\xd0\xa4\xfe\xdf\x8d\xf1\xbf\x1b\x13\x98\x84\x21\x77\xc6\xc9\x0c\xd0\x30\xc3\x65\x34\x9a\x3b\x88\x54\x06\xd1\x85\x88\x0e\x90\x23\x16\x5a\xff\xb7\x62\x86\x22\x6b\xe6\xff\x10\x64\xc5\x02\x24\xe9\xe8\x88\x5e\xa3\xff\xba\xac\xff\x75\xd9\xc0\x08\xc8\x40\xb0\xe9\xe8\x88\x22\xcd\xff\x0a\x1c\x50\x24\xcf\xf1\x1f\x8e\x14\x38\xe8\xe8\x88\x60\xc5\xff\x22\x5c\x18\x88\x13\x28\x42\x27\x47\x00\x13\x48\x1d\x30\xf5\x19\x4b\xfc\xe7\xe1\x46\x43\x29\xb0\x54\x89\x35\x17\xc0\xf3\x0d\x15\x72\x07\xd0\x2c\x60\x79\x00\x72\x07\xaa\x23\x80\x2e\xa0\x6b\x39\x01\xf1\x13\xd0\x56\x64\x44\x49\x51\x08\x29\x3c\x90\x01\x56\xd3\x06\x6b\xdd\x41\xff\x2a\x6d\x34\x24\x60\x61\x3e\x28\x42\x02\x00\x29\x71\x0a\xde\xd3\x46\x66\xca\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\xfd\x1e\xc6\x71\x8e\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsVmwareHealthRedPng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/vmware/healthRed.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
