// Code generaTed by fileb0x at "2024-04-08 14:11:59.006174 -0500 CDT m=+0.808907168" from config file "b0x.yml" DO NOT EDIT.
// modified(2011-12-13 11:16:34 -0600 CST)
// original path: ../../assets/vmware/migration-recommendation.png

package assets

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileAssetsVmwareMigrationRecommendationPng is "assets/vmware/migration-recommendation.png"
var FileAssetsVmwareMigrationRecommendationPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x48\x01\xb7\xfe\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x14\x00\x00\x00\x0d\x08\x02\x00\x00\x00\x26\x32\xd9\x09\x00\x00\x01\x0f\x49\x44\x41\x54\x78\xda\x63\xfc\xff\xff\x3f\x03\xb9\x80\x05\xce\x52\xaa\xbb\x8f\x29\x7d\xaf\x49\x91\x28\xcd\x40\x70\xac\x0e\x45\xa9\x55\xdd\x79\x62\x6d\x66\xf8\xf1\xe1\xfd\x1f\x06\xef\x36\x24\xfb\x59\x04\xc0\xce\xf9\x71\xaf\x49\x13\x22\xc0\x38\x9b\xf1\x7f\x2a\xc2\x9b\x4c\x08\xa5\x5f\xde\x7f\xf8\x01\xb5\x1f\x82\x9c\x74\x39\x41\x6e\xf9\xf3\x03\xd9\x36\xa0\x7e\x14\x9b\x95\xea\xae\x33\x30\x70\x30\x88\x28\x46\xf7\x81\xac\xb5\x6a\x02\xd9\xb6\xb7\x0e\x64\xdb\x77\x20\xfe\x83\xae\x07\x6e\x3f\xd8\xe6\x3f\x3f\x80\x36\x64\x87\x49\x82\xc8\x28\xc5\x63\x4d\x20\xdb\x7e\x80\xf5\x80\xc8\x1f\x3f\x30\x7d\xcb\x99\xc5\x09\xd7\x0c\xb6\x81\x81\x61\xe9\x8d\x1f\x60\x12\x24\xd2\xbb\xf9\x05\x54\xf3\x1f\x90\x24\xc7\x79\x0e\x20\x82\xe8\x04\x32\xbe\x4f\xfb\x0e\x0b\x30\xb0\x6f\x9d\x14\x40\x72\xfb\x1e\xfc\x00\x32\xa6\x7e\x79\x1f\x64\x6b\x08\xe4\x3e\xff\x02\x94\x05\x99\x08\x51\x0d\x74\x30\x5c\x27\x4c\xf3\x1f\x06\x88\x23\x21\x00\xa8\x1f\x28\x92\x53\xb4\x0d\xd3\xb5\xc8\x3a\x61\x9a\x7f\xfc\xf8\x0e\xf3\x97\xa1\x08\x07\x44\xe4\xde\x32\x2f\x4c\xcd\xc8\x3a\x11\x36\xa7\x60\xb3\x87\x20\x60\xa4\x24\x6d\x03\x00\x64\x56\x81\x78\xae\x9e\x5b\x34\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\x94\x93\x89\x6f\x48\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileAssetsVmwareMigrationRecommendationPng)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "assets/vmware/migration-recommendation.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
