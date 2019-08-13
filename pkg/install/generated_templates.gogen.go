// Code generated by vfsgen; DO NOT EDIT.

package install

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 8, 13, 11, 48, 4, 868967991, time.UTC),
		},
		"/flux-helm-operator-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-helm-operator-account.yaml.tmpl",
			modTime:          time.Date(2019, 8, 13, 7, 52, 26, 296527299, time.UTC),
			uncompressedSize: 948,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x39\x6f\xdc\x30\x10\x85\x7b\xfe\x8a\x01\x5c\x38\x09\x2c\x05\xee\x02\x75\xb6\x8b\x14\x09\x52\x28\x47\x13\xa4\x18\x92\x4f\x59\xc6\x5c\x8e\x30\x24\x37\x87\xb0\xff\x3d\x90\xb4\x06\xbc\x8e\xed\x34\xdb\x8d\xe6\xd2\x9b\xc7\xaf\x69\x1a\x73\x46\x9f\x36\xa0\x0c\xdd\x05\x07\x62\xe7\xa4\xa6\x72\x41\x2e\xd6\x5c\xa0\xa4\x12\x91\x2f\x88\x93\x3f\x4a\x91\x0d\xc9\x87\xf4\x9d\x58\x61\xce\x48\x52\xfc\x4d\x09\xf0\xf0\x34\x88\xd2\xbb\x6a\xa1\x09\x05\x99\x7e\x86\xb2\x59\x46\x1a\xcb\x19\x7e\xfe\x03\x72\x26\x27\xa9\xa8\x44\x7a\xd1\x5f\x5f\xdd\xbc\x6c\x0d\x8f\xe1\x0b\x34\x07\x49\x1d\xed\x2e\xcd\x6d\x48\xbe\xa3\x8f\xab\xaa\xab\x55\x94\xd9\xa2\xb0\xe7\xc2\x9d\x21\x8a\x6c\x11\xf3\x1c\x11\x25\xde\xa2\xa3\x21\xd6\x5f\xcd\x06\x71\xdb\xc8\x08\xe5\x22\x6a\x9e\x2e\x4d\x13\x85\x81\xda\x0f\xbc\x45\x1e\xd9\x81\xf6\xfb\x43\xf7\xf2\xd9\xd1\x34\x1d\x57\xa7\x89\x90\xfc\x7e\x6f\x66\xcf\xee\x8b\x55\xcb\xae\xe5\x5a\x36\xa2\xe1\x0f\x97\x20\xa9\xbd\x7d\x93\xdb\x20\xaf\x77\x97\x16\x85\xef\x6e\xb9\x59\xdd\xeb\x25\xe2\x94\x87\x18\xad\x11\xcb\x78\x43\x3c\x86\xb7\x2a\x75\xcc\x1d\x7d\x3d\x7f\x75\xfe\x6d\xd9\xa9\xc8\x52\xd5\xe1\x28\xb9\x83\xda\x7b\x89\x86\x92\xa4\xfe\xd0\xf8\xb9\x7f\xff\x74\xef\x09\xae\xbf\x5e\xc9\x39\xad\x09\x12\xd1\x63\x98\x17\xdc\x99\xf0\x8c\x36\x43\xf4\xef\x9b\x3c\xb3\x3d\x57\xfb\x03\xae\x1c\x5c\x7e\x14\xcd\xff\x08\x7f\x88\xd6\x43\xf6\x1e\xa3\x2d\xe6\x39\xf2\x18\xb8\xc6\xb2\xe2\x37\x53\xfa\x37\x00\x00\xff\xff\xad\xec\xff\x2b\xb4\x03\x00\x00"),
		},
		"/flux-helm-release-crd.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-helm-release-crd.yaml.tmpl",
			modTime:          time.Date(2019, 8, 13, 11, 45, 48, 539576538, time.UTC),
			uncompressedSize: 4106,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x97\xdd\x8f\xe3\x34\x10\xc0\xdf\xf3\x57\x58\x07\x52\x01\x5d\x0a\x2b\x10\x82\xbc\xa0\x13\xa7\xd3\xa1\x03\xee\xb4\xd5\xee\xcb\x6a\x91\xa6\xc9\x34\x35\x75\x6c\x33\x9e\x84\x2d\x88\xff\x1d\xd9\x69\xba\x4d\x36\x1f\xdd\x76\x17\x5e\xce\x4f\xad\xe7\xc3\x33\xbf\xf1\xc7\x24\x8e\xe3\x08\xac\xbc\x46\x72\xd2\xe8\x44\x80\x95\x78\xc7\xa8\xfd\x3f\x37\xdf\x7c\xe7\xe6\xd2\x7c\x59\x5d\x2c\x91\xe1\x22\xda\x48\x9d\x25\xe2\xc7\xd2\xb1\x29\x2e\xd1\x99\x92\x52\x7c\x8d\x2b\xa9\x25\x4b\xa3\xa3\x02\x19\x32\x60\x48\x22\x21\x34\x14\x98\x88\x35\xaa\x82\x50\x21\x38\x74\x73\xff\x67\xbe\x52\xe5\x5d\x9a\xcd\xa5\x89\x9c\xc5\xd4\x6b\xe6\x64\x4a\x5b\xab\x1e\x48\x6b\x0f\xce\x2b\x08\x51\xaf\xfb\x16\x55\x71\x59\x3b\x0b\xb3\x4a\x3a\x7e\xd7\x95\xfc\x2c\x1d\x07\xa9\x55\x25\x81\x6a\x87\x10\x04\x6e\x6d\x88\x7f\xbd\x77\x1e\x8b\x35\x45\x42\xb8\xd4\x58\x4c\x44\x10\x58\x48\x31\xf3\x73\xe5\x92\x76\x69\xee\x94\x1d\x03\x97\x2e\x11\x7f\xff\x13\x09\x51\x35\xd0\xaa\x8b\xfb\x7f\x7b\xaf\x35\x81\x20\x0a\x96\x48\x15\x66\x89\x60\x2a\xb1\x99\x62\x43\x90\xe3\x7e\xae\x02\x25\x33\xf0\x28\x6b\x1f\xc6\xa2\x7e\xf5\xe1\xa7\xeb\xaf\x17\xe9\x1a\x0b\x48\x76\x66\x96\x8c\x45\x62\xd9\xc4\x14\x5c\xed\x60\x36\x83\xf0\x8f\x52\x92\x5f\xef\x66\x96\xae\x81\x78\x76\x7b\x20\xed\xf3\x50\x5b\x05\x4e\x9e\x41\x5b\x20\x04\x6f\x3d\x1d\xc7\x24\x75\xde\x11\x59\x60\x46\xd2\x89\x78\xf1\xdb\x0d\xc4\x7f\x7d\x15\x7f\x7f\xfb\xd9\x4d\xbc\xfb\xf5\x45\x33\xf5\xf9\x0f\x9f\xbe\x68\x19\x32\x50\x8e\xbc\xe7\xfd\x1f\x2c\x28\x0b\x34\x25\xf7\x2f\x24\x35\x63\x8e\xd4\x91\xad\x0c\x15\xc0\x41\xfa\xed\x37\x1d\x54\x0e\xf9\x1a\x54\xd9\x65\xd8\x38\x5c\x1a\xa3\x10\x74\xd4\x71\x97\xe2\x95\xcd\x09\xb2\x81\x7c\xfb\xac\xc8\x28\xb5\x84\x74\xd3\x6f\x61\x96\xbf\x63\xca\x5d\x42\x03\x05\xf6\x03\x35\x2c\xd5\x83\xe5\xc7\x43\xd8\x07\xff\x78\x33\xc2\x94\x10\xf8\x04\xcb\x4c\x3a\x1f\xe9\x5b\x63\x36\x3d\x69\x4c\x59\x0f\x54\x5b\x4c\x54\x5c\x8c\x56\xdd\x8f\x3f\x41\x8e\x78\xed\x0b\xa7\xf2\xdb\xe4\x8d\x54\xb8\xf0\x2c\x78\x60\xc3\x00\x11\x6c\x3b\x12\xc9\x58\xf4\xe4\x3e\x52\xf9\xf6\xd1\xf7\x77\x50\xeb\xe4\xd7\x63\x6c\x7b\xec\xae\xee\x9e\xf9\x91\x33\x19\x32\x74\x6f\xc8\x14\xcf\x9b\xdb\x78\xe0\xa9\xd1\x2b\x99\xff\x02\xf6\x1d\x6e\x2f\x71\x35\x96\xc3\x80\x7f\x71\x1c\xbf\xe9\x50\xc4\x28\x47\x31\x7e\xbf\x35\x63\x83\xdb\xb3\xec\x8d\xf5\x6f\x09\xa8\x29\x27\x43\x47\xc8\x3f\x5a\x7e\xc3\x7e\xc4\x19\xc6\xf9\x38\x7d\x67\x45\x1a\xd4\x22\xb4\x14\x4f\xc3\xb4\x24\x75\x32\xd2\x92\x26\x93\x79\x66\x22\xa1\x37\xf1\x57\xe3\xd3\xc0\xb0\xc0\xeb\x93\x69\x78\xe3\xff\x15\x87\xd1\xf8\xbe\x87\x42\xdc\xee\xe6\xda\x97\x5c\x4f\xb6\x6d\xfd\xc3\x23\x3c\xa9\xfc\x60\x83\x4e\x5a\x1c\x16\xb0\xa3\x5c\x8d\x74\x47\x3d\xf5\x0c\x9e\xba\xda\xbd\x48\xda\x11\xe4\x92\x67\x2f\xc5\x50\xe9\xc7\xcb\x9e\xf7\x3f\xe5\x47\x54\xbb\xe9\x11\x72\xc9\xe2\x13\xa1\x0d\x8b\xcc\x7f\x08\x61\x26\x96\x5b\xf1\xfe\xd5\xa2\xc7\x68\x78\x7f\x4d\xac\x46\xe3\x67\x63\xd0\xce\x6d\xa4\x7d\x8d\xf6\xca\x66\x03\xfd\xd7\xf8\x86\x6c\x63\x26\xb4\xc6\x49\x36\xb4\xf5\xb4\xc3\x4d\xfe\x52\xcc\x76\x1f\x3d\x8f\x06\x7f\xef\xed\x4c\xfe\x25\xa9\x63\xf9\x9f\xd0\xd7\xd4\xa3\xf9\xce\x3b\x2f\x52\x87\x45\x85\x74\x6c\xb0\xe1\x34\x7c\x28\x95\xaa\x7b\xc6\xfe\xb5\x9f\xf4\xd1\xfc\x37\x00\x00\xff\xff\xd9\xc1\x96\x44\x0a\x10\x00\x00"),
		},
		"/helm-operator-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "helm-operator-deployment.yaml.tmpl",
			modTime:          time.Date(2019, 8, 13, 8, 42, 32, 244594645, time.UTC),
			uncompressedSize: 5690,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\xcd\x8e\xdb\xc8\x11\xbe\xcf\x53\x14\x66\x0e\xbe\x8c\xc8\x31\xd6\xd9\x03\x0d\x1f\x92\xdd\xec\xda\x80\xed\x0c\x32\x46\x80\x9c\x76\x4b\xcd\x92\xd8\x51\xb3\x9b\xe9\x2e\x4a\x21\x84\xcd\xb3\x07\xd5\x4d\x51\xa4\x28\x69\x66\x36\x87\x9d\x8b\x65\xb2\xfe\xeb\xab\x3f\x2e\x16\x8b\x1b\x6c\xf4\x3f\xc8\x07\xed\x6c\x01\xd8\x34\x21\xdf\xbe\xbd\xd9\x68\x5b\x16\xf0\x23\x35\xc6\x75\x35\x59\xbe\xa9\x89\xb1\x44\xc6\xe2\x06\xc0\x62\x4d\x05\xac\x4c\xfb\x9f\x45\x45\xa6\x5e\xb8\x86\x3c\xb2\xf3\xfb\x3d\xe8\x15\x64\x5f\xb1\xa6\xd0\xa0\x22\xf8\xed\xb7\x9e\x3a\xfe\xb7\x80\xfd\x7e\xfa\x76\xbf\x07\xb2\xa5\x90\x85\x86\x94\x88\xf6\xd4\x18\xad\x30\x14\xf0\xf6\x06\x20\x90\x21\xc5\xce\xcb\x1b\x80\x1a\x59\x55\x9f\x71\x49\x26\xa4\x07\x97\x2d\x11\x5e\xf6\xc8\xb4\xee\x12\x29\x77\x0d\x15\xf0\x77\x52\x9e\x90\xe9\x06\x80\xa9\x6e\x0c\x32\xf5\xa2\x47\xde\xc9\x9f\x99\x68\xb9\xaa\x47\xfe\xd0\x5a\xc7\xc8\xda\xd9\x11\x4f\xe3\x5d\x4d\x5c\x51\x1b\x32\xed\xf2\xa0\x3c\x8a\x09\xb7\xec\x5b\xba\x8d\x44\x07\x9f\xe3\x6f\xf2\x5b\xad\xe8\xcf\x4a\xb9\xd6\xf2\xd7\xeb\xea\xb6\xce\xb4\x35\x85\xa2\x8f\xf7\x5f\x2d\x2e\x0d\x7d\xd3\xc6\x90\xff\xf6\xf9\x29\x45\x3d\xfd\x2d\x7a\xcb\xa3\x14\x36\x61\xa1\x70\x78\x07\xa0\x9c\x5d\xe9\xf5\x17\x6c\x8a\xd1\xc3\xb9\xb3\x89\x6f\x91\xa8\x27\x94\x25\xad\xb0\x35\xfc\xc5\x95\x54\xc0\xc3\xf7\x0f\x0f\x57\x14\x93\xe7\x30\x62\x0e\x92\x0a\x9e\x2a\x4e\xcf\x92\xf3\xbd\x6f\x83\x57\x3f\x90\xe7\xa7\xe1\x7d\xc2\xce\xf5\xb7\x64\x82\xfc\x3a\xf1\x83\x3c\x1f\x51\x77\xd9\x97\x77\x0f\x0f\x47\x11\x3d\xdd\xdd\xe1\x5f\xf8\xa7\x6b\x61\xa7\x8d\x01\x4b\x54\x02\x57\x14\x08\x78\xe7\x0e\x89\x11\xcb\x3b\x21\x41\xcb\xc0\x0e\x28\x30\x2e\x8d\x0e\x15\x6c\xd1\xe8\x12\x99\x4a\xf8\xf6\xf9\x69\x10\xa7\x9c\xb5\xa4\x22\x7c\x00\xd7\xa8\x6d\x60\x48\xae\xcd\x34\x5f\x4e\xe8\xdd\xb9\x84\xde\xbd\x38\xa1\x77\x57\x13\x7a\x07\x29\xba\xb1\x8e\x60\xd3\x2e\xc9\x5b\x62\x8a\xc8\x66\x13\x66\xe6\xcd\x83\x3e\x52\x33\x4d\xfd\xdd\x1f\x99\xfa\x73\x5e\xbf\x3b\x7a\xbd\xdf\x93\x2d\x47\xc4\xdf\x2a\x82\x95\x33\xc6\xed\xb4\x5d\xf7\xd9\x06\x1d\x60\xe5\x3c\xb4\x41\x9e\x21\xa8\x36\xb0\xab\x75\xa0\x12\x36\xd6\xed\xec\x2f\x95\x0b\x1c\x60\xa5\x0d\xdd\x0f\x82\x76\x95\x56\x55\xc2\xc8\x11\x46\x0e\x4a\x77\x80\x8e\x30\xc9\x0f\x0f\x6e\x67\x61\xad\x59\x3a\xa3\x03\x8f\x5c\x1d\x51\x01\x5c\xa1\xed\x15\xaf\x35\x57\xed\x12\x9c\x17\x38\x82\xd1\x1b\xca\x04\xa6\x6f\x8c\x01\x34\xc1\x0d\x2a\x6a\xe9\x2f\xa0\x8f\xf9\xd0\x96\x5d\xe4\x51\xce\x32\x6a\x4b\xfe\x1e\x96\x64\xdc\x2e\x3b\x0b\xfb\x1a\xbb\x24\x70\x27\x78\x66\x27\x6d\x6e\xab\x4b\x02\xb4\x10\x42\xf5\x4b\x02\xd5\x89\xbb\x32\x41\xb4\xb3\x62\x67\xed\x3c\x25\xbb\x9d\x25\xf8\xf5\x53\x29\xaf\xb8\xfb\x49\x1b\xfa\xf5\x7d\x0c\xa4\xc0\x1f\xad\xa2\xfb\x3e\x16\x6f\x3c\x0d\x82\x92\xaf\x53\x19\x3f\x6b\xfe\xd8\x2e\x63\x7c\x32\xf8\xfa\x97\xe8\x0b\x59\xf6\x1d\x6c\xa8\x83\x50\xb9\xd6\x94\xb0\x3c\xca\xb8\x4d\x26\xde\xf6\xc1\x4c\x82\x6e\x8f\xb6\xdf\x8a\xde\x18\x26\x2a\x41\x5b\xf8\x6f\x9e\x85\x50\xe5\xf3\x70\x1c\xc0\x1e\x42\x55\x6a\xff\xaa\x32\x0c\xa1\x7a\xbe\xfc\x52\x0f\x92\x52\x78\x7a\xfa\x38\x81\xf8\xcd\xb1\x2c\x9f\x3e\x46\x37\xd9\x01\x2a\x45\x21\x44\xf7\x7f\xee\xf1\x12\x34\x3b\xdf\xcd\x9a\xf2\x5a\xf3\x62\x43\xdd\xeb\xba\xf1\xdc\x88\x31\xf1\xcc\xf2\x08\x72\xb2\x43\x20\x3d\x61\xb9\x70\xd6\x74\xf7\xb0\x23\xd8\x39\xfb\x86\x61\x49\x20\x93\x4b\x8c\x57\x55\xed\xca\x9b\x57\xb4\x5c\x1d\x86\xfa\x3b\xa0\x64\x28\xc1\xa1\x5c\xb8\xc2\x23\xd0\x85\x31\x08\x4c\x0f\x31\x13\xb0\xa5\xa0\xbd\x07\xca\xd6\xd9\x3d\xe0\x01\x4c\x65\x5c\x7c\x84\x2a\x83\x4f\xab\x41\xc4\x44\xcf\xbf\xda\xc0\x11\x80\xa1\x55\x55\xd4\x77\x1f\x83\xdf\x87\x62\x54\x0d\x03\x3f\x1a\x09\x43\x07\x8d\xd3\x96\x03\x20\x43\x4e\xac\x72\x81\x44\x99\x0b\xc8\x74\x5f\x0e\x80\x01\xf0\xa0\x5e\xd4\x1e\x3b\x47\x3f\x53\xda\x40\x27\x75\xb0\xa1\xee\x3e\x5a\x38\x6a\x28\x87\xe2\x3c\x74\x92\x41\xcc\xa8\x54\x71\xe9\xb6\x74\x0f\x3b\xcd\x95\x44\x67\x5a\x92\x7d\x25\xc5\xd5\x4b\x9c\x26\x54\xd5\x20\x44\x82\xa8\x6d\x74\x3a\x81\xe5\x50\xe8\x54\x42\x45\x9e\x2e\x97\xcc\x14\x81\x2f\x19\x0a\xb1\x6c\x84\x2d\xa5\xe6\x7a\xd9\xfc\x2e\xf0\x5d\xee\xf9\x27\xd3\x1d\x3d\x45\xec\x44\xe9\x09\x71\x43\xb1\x69\x0a\x59\x87\xb5\x39\x69\x80\x68\xcb\x3e\x17\xfd\x90\x40\x25\x48\xd1\x3e\xae\xb7\x5d\x16\x27\x8b\x41\x66\xf2\x32\x4f\x24\x7d\x24\x4d\x4b\x61\x1b\x8e\x9d\x6b\x50\xc8\x71\x0e\xf9\x9a\x7c\xaa\x89\x1a\x37\x94\x7a\xb8\xc8\xcd\x8f\x82\x8f\x9e\x5f\xce\xc5\xd8\xf6\x85\xd8\xfe\xda\xac\xc4\x09\x3b\x96\x72\x5d\x45\xb4\x71\x24\x91\xea\x86\xbb\x1f\xb5\x2f\x60\x3f\x34\xb6\x61\x16\x0d\xfb\xf4\x7c\xb3\x38\x59\x8a\xfb\x5c\x79\x8a\xf9\xb1\x0e\x6e\x0b\xd9\xef\x03\xdf\x82\xae\x71\x4d\x69\x4a\x4f\x38\x33\xf8\x49\xdb\xb8\xbf\x41\x2d\xf3\xd6\x93\x92\x53\xe7\x28\xcf\x93\x21\x0c\x24\x53\x35\xca\x80\x6d\xba\x93\xa4\x72\x2b\xe6\x26\x14\x79\x5e\xb5\xcb\xac\x74\x6a\x43\x3e\x53\xae\xce\x7d\xbe\x23\xdc\xd2\xce\xf9\x4d\xc8\x27\xda\x72\xc6\x75\x18\x09\x17\x4c\xc8\xb9\x23\xa7\x90\x98\xc0\xb8\x9e\x54\x0d\x24\x9d\x05\xf4\xd2\xb5\x8b\x8d\x42\x95\x53\xb1\xc5\x43\xf6\xf6\x21\x7b\x3b\x65\x7a\x6c\x8d\x79\x74\x46\xab\xae\x80\x4f\xab\xaf\x8e\x1f\x3d\x85\xb1\x6f\x8d\xf3\x3c\xba\x55\x86\xb5\x92\xb9\x99\x1e\x08\x29\x0d\x8f\xce\x73\x01\xdf\x3d\x7c\x77\xdc\xf0\x3d\x05\xd7\x7a\x45\x61\x3c\x35\x3c\xfd\xbb\xa5\xc0\x61\x3a\x49\x54\xd3\x16\xf0\xa7\x87\x7a\xf2\xb0\xa6\xda\xf9\xae\x80\xef\xdf\x7d\xd1\xc3\x8b\x54\x62\x5f\x04\xe7\x23\x19\x77\xf0\xc9\x2a\xd3\x96\x94\xfa\x7e\xbf\x22\x4d\x37\x9a\x8b\x8b\x97\xf3\xf3\x4e\x2c\x22\xa5\x40\xdf\x1f\xfa\xe5\x68\x45\xaa\xe8\x30\x58\x4a\x52\x06\x3d\x95\xa9\x43\x66\x23\xde\xb3\x93\x3f\xa1\x39\x5a\xf3\x88\x5c\x15\x90\x7b\xe7\x38\x2e\x0f\x13\x0a\x29\xc9\xbf\x59\xd3\x15\x20\xa7\xe0\x33\x13\x1e\xae\x8e\xed\xa9\xba\xc9\x28\x99\xcf\xd2\xcb\xfd\x77\x6e\xf9\x44\xd4\xfc\x5e\x7a\xbe\x79\xcc\x45\x6e\xd1\xf7\x22\x05\xbe\xf9\x99\xf5\xe4\x05\xfd\xe2\x15\x62\xf3\xc8\x77\xf9\x44\x3e\x13\xe0\x8b\xc7\xea\xa5\xe0\x08\xc3\x84\x6c\x9a\xdb\xab\x2a\xf0\x85\xf2\x4f\x29\x4f\x54\x5c\x4d\xf3\x05\x87\x2e\x67\x7b\xe2\xd0\x1c\xac\x57\x55\xe0\x0b\xe5\x9f\x52\xce\xea\xe1\x04\x6c\xe8\xd7\x93\x5e\xf0\xd1\xed\xa4\xec\x57\xd2\xba\x27\x57\xb2\x44\x79\xb1\xe0\xf8\x68\x31\x7c\x75\xfa\x30\x39\x27\x4f\xbe\x3e\x9d\x7f\xda\x07\x54\xce\xdc\x45\xe8\x02\x53\x3d\x58\xf5\xfc\x27\x97\x3b\x39\xed\xfb\x13\xa0\xf5\xf1\xa3\xd0\x39\xfb\x0e\x37\x38\x79\x5e\x34\xc8\xd5\x87\x79\x94\x72\x85\x99\xf2\x7c\x89\x9b\xa2\x0d\x1f\x26\x79\x39\x21\xd9\x50\x77\x56\xb8\xdc\xec\xd9\xb8\xfe\x4f\x0d\xbb\x64\x55\x64\xbc\x62\xd3\x96\xbc\x5e\x75\x57\x6d\x7a\x99\xd3\x67\x61\xfd\x83\xab\x65\xc1\x04\xd7\xc6\x15\xf8\xf2\x97\x95\xe7\xbf\xa8\x24\x08\xff\x3f\xd9\x98\xf1\x9f\xcb\xc7\x8c\xe8\xc5\x19\x99\x9b\xf7\xe2\x9c\xcc\x58\xcf\x65\xe5\x77\xba\x7f\x2c\xcf\xff\x05\x00\x00\xff\xff\x5b\x01\xe0\x43\x3a\x16\x00\x00"),
		},
		"/tiller-ca-cert-configmap.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "tiller-ca-cert-configmap.yaml.tmpl",
			modTime:          time.Date(2019, 8, 13, 11, 48, 4, 869077929, time.UTC),
			uncompressedSize: 226,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xc1\x0a\xc2\x30\x0c\x40\xef\xfd\x8a\xfc\x40\x07\x82\xa7\xde\xa4\x78\x53\x2f\x0e\xef\xb1\xcd\xb4\xd8\x66\xa3\x8b\x22\xd4\xfe\xbb\x6c\x4c\xd1\x63\x78\x2f\x2f\x29\x05\x42\x07\xcd\x96\xf1\x1c\xa9\x0d\x31\x52\x6e\x77\x47\xa8\x55\x6b\xad\x70\x08\x27\xca\x63\xe8\xd9\xc0\x63\xa5\x6e\x81\xbd\x01\xdb\x73\x17\x2e\x7b\x1c\x54\x22\x41\x8f\x82\x46\x01\x30\x26\x32\xd0\xc5\xfb\x53\x5f\x29\x26\x2d\x71\xd4\x0e\xb5\x9b\xe5\xe5\xc8\x01\x13\x8d\x03\x3a\x82\x5a\x97\x95\x79\x34\x50\xca\x3f\x2d\x05\x88\xfd\xa4\x7d\xfa\x0e\x1b\x97\xc5\xc0\x4b\x4d\x31\xf6\xc4\x02\x6b\x68\xbe\x1f\xdb\x8d\xa5\x2c\xb6\x67\x99\xc8\x6f\xe1\x1d\x00\x00\xff\xff\x60\xd7\x1c\xac\xe2\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-helm-operator-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-helm-release-crd.yaml.tmpl"].(os.FileInfo),
		fs["/helm-operator-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/tiller-ca-cert-configmap.yaml.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}