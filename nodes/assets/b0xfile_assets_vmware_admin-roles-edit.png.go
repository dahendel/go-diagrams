// Code generaTed by fileb0x at "2024-04-08 14:11:59.465197 -0500 CDT m=+1.267929043" from config file "b0x.yml" DO NOT EDIT.
// modified(2011-12-13 11:14:34 -0600 CST)
// original path: ../../assets/vmware/admin-roles-edit.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsVmwareAdminRolesEditPng is "assets/vmware/admin-roles-edit.png"
var FileAssetsVmwareAdminRolesEditPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x01\x02\xfe\xfd\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x14\x00\x00\x00\x14\x08\x06\x00\x00\x00\x8d\x89\x1d\x0d\x00\x00\x00\x04\x67\x41\x4d\x41\x00\x00\xb1\x8e\x7c\xfb\x51\x93\x00\x00\x00\x20\x63\x48\x52\x4d\x00\x00\x7a\x25\x00\x00\x80\x83\x00\x00\xf9\xff\x00\x00\x80\xe8\x00\x00\x75\x30\x00\x00\xea\x60\x00\x00\x3a\x97\x00\x00\x17\x6f\x97\xa9\x99\xd4\x00\x00\x01\x8c\x49\x44\x41\x54\x78\x9c\x62\xfc\xff\xff\x3f\x03\x35\x01\x40\x00\x31\x51\xd5\x34\x20\x00\x08\x20\xe2\x0c\x34\x31\xf9\x0f\xc6\x40\x70\xf6\xec\xd9\xff\xb3\x66\xcd\xc2\xe9\x2d\x80\x00\x62\x00\x79\x19\x1f\x4e\x4b\x4b\xfb\xff\xff\xcc\x99\xff\xff\x67\xce\xfc\xff\xdf\xd8\xf8\x3f\x0c\x80\xc5\xb1\xa8\x07\x08\x20\xbc\x2e\x4c\x4f\x4f\x07\x69\x64\x48\x9f\x35\x0b\xe4\x34\xb8\x38\x23\x23\x23\x4e\x3d\x00\x01\x84\xd3\xc0\xb3\xe9\x8c\xff\x67\xa6\x9d\x65\x00\x7a\x8f\xc1\xd8\xd8\x98\xc1\x04\x68\xe0\xd9\x99\x33\xe1\xbe\xc2\x05\x00\x02\x88\x11\x9b\x24\xc8\x30\xe3\x34\x63\x20\x83\x01\xec\x32\x93\xb3\xc6\x0c\x33\x81\x86\x01\xc3\x0f\x6c\x38\x08\x80\x2c\x82\x01\xa0\x1c\xdc\xc9\x00\x01\x84\x11\x06\x67\xd2\x80\xd4\x19\x60\x58\xcd\x04\x62\x20\x7b\xa6\x31\x03\x30\x08\xcf\x80\x31\x28\xdc\x40\x0e\x04\x61\x10\x1b\x59\x1c\xa6\x1f\x20\x80\x88\x36\x0c\x17\x40\x37\x10\x20\x80\xe0\x61\x38\x6b\x16\xd0\x9b\xa0\x30\x82\x7a\x73\x16\x90\x4e\x3b\xf3\x9f\xf1\x2c\x52\x64\x40\x83\x03\x8c\x71\x01\x80\x00\xc2\x8c\x14\x60\xac\xc2\x0c\x03\x1b\x80\x64\x20\xc8\x20\x63\xa0\xb3\x41\x18\x97\xa1\x00\x01\xc4\x02\x63\xa4\xa7\x83\x49\x88\x99\x50\xc3\xd0\x01\xc8\xa2\x99\x67\x91\xd8\x58\xd4\x00\x04\x10\xdc\x40\xa0\xff\x71\xfb\x03\x19\x9c\xc5\x2f\x0d\x10\x40\x2c\xf8\xa5\x51\xc1\xcc\x33\xff\x19\xd2\x4d\x18\xe1\x6c\x6c\x00\x20\x80\x88\x2e\x1c\xd2\x81\x2e\x63\x04\x86\xc8\x2c\xe3\xff\x60\x0c\x62\xa7\x63\x71\x2d\x40\x00\x11\x6d\xe0\x59\x2c\x9a\xb1\x89\x01\x04\x10\xd1\x5e\x9e\x69\x0c\xd4\x6d\x4c\x58\x1d\x40\x00\x61\xcd\x7a\xe8\x00\x54\x48\xe0\xb5\x0c\x29\xeb\x01\x04\x10\x51\x06\x92\x02\x00\x02\x88\xea\x25\x36\x40\x80\x01\x00\xb2\xa4\x22\xb4\xf4\x83\xa0\xf2\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\xa4\x53\xba\xa9\x01\x02\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsVmwareAdminRolesEditPng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/vmware/admin-roles-edit.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
