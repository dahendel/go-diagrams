// Code generaTed by fileb0x at "2024-04-08 14:11:59.473707 -0500 CDT m=+1.276438626" from config file "b0x.yml" DO NOT EDIT.
// modified(2011-12-13 11:16:30 -0600 CST)
// original path: ../../assets/vmware/maintainenceMode.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsVmwareMaintainenceModePng is "assets/vmware/maintainenceMode.png"
var FileAssetsVmwareMaintainenceModePng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\xad\x01\x52\xfe\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x10\x00\x00\x00\x10\x08\x06\x00\x00\x00\x1f\xf3\xff\x61\x00\x00\x00\x04\x67\x41\x4d\x41\x00\x00\xb1\x8e\x7c\xfb\x51\x93\x00\x00\x00\x20\x63\x48\x52\x4d\x00\x00\x7a\x25\x00\x00\x80\x83\x00\x00\xf9\xff\x00\x00\x80\xe8\x00\x00\x75\x30\x00\x00\xea\x60\x00\x00\x3a\x97\x00\x00\x17\x6f\x97\xa9\x99\xd4\x00\x00\x01\x38\x49\x44\x41\x54\x78\x9c\x62\xfc\xff\xff\x3f\x03\x32\x60\x04\x62\x03\x97\x48\xb0\xa0\x4d\x74\x05\xe3\x91\xa5\x1d\x60\xf6\xb1\x5d\xcb\x19\x39\x99\x18\x30\x00\x40\x00\x61\x08\x81\x34\xef\x59\x3b\x07\xcc\xfe\xf7\xef\x1f\x98\x06\xf1\x55\xd4\x0d\xfe\xa3\xab\x05\x01\x80\x00\xc2\x62\x26\x02\x1c\x5b\xde\x05\xd6\xf4\xf6\xe3\x37\x9c\x6a\x00\x02\x88\x05\x97\xc4\xca\x79\x93\xf0\x99\x0d\x07\x00\x01\x84\xd5\x05\x2e\xc1\x29\x0c\xc2\xfc\x5c\x70\xbe\xa3\xb3\x2b\x4e\x03\x00\x02\x08\xc3\x80\xf3\x7b\x96\x83\xc2\x91\xe1\xef\x3f\x54\x2f\x3f\xbd\x7d\x81\x11\x9b\x01\x00\x01\x84\x61\x00\x4c\xd5\xef\x3f\x7f\xc1\x86\xc0\x0c\x7a\xfd\xed\x2f\x56\x17\x00\x04\x10\xde\x40\x44\x06\xac\x4c\x58\x1d\xc0\x00\x10\x40\x38\x03\xf1\xcf\xdf\x7f\x0c\x3f\x7e\xfe\x46\x18\xc0\x8c\x5d\x1d\x40\x00\xe1\x34\x00\xe4\x85\x4f\xdf\x7e\x22\x14\xe2\x70\x01\x40\x00\x11\xed\x05\x5c\x06\x00\x04\x10\x4e\x17\xfc\xfa\xfd\x97\xe1\xc3\xe7\xef\x08\xfe\xaf\xdf\x0c\x9c\xec\xac\x18\xea\x00\x02\x88\x11\x3d\x2f\x18\x02\x93\xf2\xc2\x19\x7d\x0c\x67\x6f\x3c\x41\x11\xaf\x29\x4c\xc5\x1a\x95\x00\x01\x84\xd5\x05\xa0\x00\x04\x69\x80\x81\x92\xe6\xc9\xb8\x1c\xca\x00\x10\x40\x58\xc3\xe0\xef\x5f\x48\x26\x82\xd9\x08\xe3\x63\x03\x00\x01\x84\xe1\x05\x10\x90\x56\x85\xe4\x3c\x90\x01\xc8\x6c\x6c\x06\x00\x04\x18\x00\x94\x75\x6d\x7b\x1a\xe5\x52\xde\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\x32\x83\x1b\x19\xad\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsVmwareMaintainenceModePng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/vmware/maintainenceMode.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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