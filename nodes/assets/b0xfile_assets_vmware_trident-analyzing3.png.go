// Code generaTed by fileb0x at "2024-04-08 14:11:59.291986 -0500 CDT m=+1.094718584" from config file "b0x.yml" DO NOT EDIT.
// modified(2011-12-13 11:19:06 -0600 CST)
// original path: ../../assets/vmware/trident-analyzing3.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsVmwareTridentAnalyzing3Png is "assets/vmware/trident-analyzing3.png"
var FileAssetsVmwareTridentAnalyzing3Png = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xc3\x4f\x28\x43\x01\x1c\xc0\xf1\xef\x9a\xc4\xb4\xda\x38\xcc\xbf\x83\x03\xe5\x68\x4d\x0e\x2e\x6b\xf6\xc7\x9e\xda\xcc\x38\xcc\x69\xaf\x51\x3b\xce\xd2\x0a\xa1\x3d\x93\x9a\xda\x0e\x48\xcb\x9e\x8b\xf2\x8e\x84\xab\x8b\x03\x57\x29\x27\x12\xe5\xf4\x76\xf2\x8e\xcb\xf8\xc9\xe1\x53\x4e\xc4\xa7\x9d\x8e\x7e\x07\xe0\x54\xa2\xa1\x24\xe0\x06\x5c\x1d\xed\xc0\x97\xf7\x28\x0d\x74\xae\x44\x17\x57\xa1\xab\xe7\x9f\x8d\x13\xbd\x17\x68\xcb\x06\x62\x01\xb8\xaa\x6e\x7e\xcf\x1d\x02\x43\x4b\xd1\x64\x0c\x36\x46\xa0\x58\x82\xa6\x40\xd1\x84\xc2\x18\x34\x54\x98\xac\x81\x27\x77\x90\xbe\x8f\x00\x05\x25\x14\x58\x58\x7b\xc9\xb4\x44\xfc\xf6\xf3\x62\x59\x73\xa8\xcf\xae\xc1\xbd\xf1\x37\xdd\x7a\x6f\xa6\x1a\x62\xd4\x7d\x89\xb3\xfd\x54\xfe\xd8\xa8\xfb\x7e\x73\x5b\xeb\x35\x4d\xd3\x34\xdb\x8c\x5c\xce\x3f\xaa\xc3\x61\xdd\xfc\x78\xbd\x19\x8d\x97\x2c\xc5\x9d\xbd\xf6\x4c\x44\xdc\xbb\x96\xd2\x2d\x22\xfe\xe0\x67\xc5\x94\xd3\xaa\xd1\xca\xdd\x7a\x7f\x44\x44\xfa\xd4\x5a\x2a\xff\xf0\xb4\x3c\x3b\x10\xc4\xce\x5d\x62\xdb\x96\xf4\x56\x32\x00\x4a\x38\x1e\xba\x98\x52\x77\xfe\x02\x00\x00\xff\xff\xcc\x31\x05\x24\xff\x00\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsVmwareTridentAnalyzing3Png)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/vmware/trident-analyzing3.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
