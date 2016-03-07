// Code generated by go-bindata.
// sources:
// templates/dao/.pgx.tmpl.swp
// templates/dao/pgx.tmpl
// templates/dao/sqlx.tmpl
// DO NOT EDIT!

package data

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesDaoPgxTmplSwp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9a\x6d\x6c\x14\x4f\x1d\xc7\x87\xa2\xd5\x0a\x08\x41\x41\xc3\x0b\x1c\x56\xc0\xbb\xda\xdb\x6b\xab\x11\x24\x56\x43\xdb\x3b\x68\x52\xe8\xc3\x1d\xa0\xd4\x86\x6c\xef\xe6\x8e\xa3\x77\xb7\xcb\xee\x5e\x7b\xcd\xba\x1a\x13\x7d\x81\x22\x44\x34\x31\x10\xe2\x0b\x5f\x49\x88\x18\x8c\x1a\x83\x31\xd1\x04\x1f\x63\xf4\x8d\x0f\xaf\x7c\xa3\x21\xf1\x8d\x31\x11\xa3\xc1\xa8\xff\xff\x77\x76\x76\xef\xf6\x9e\xda\x6b\x29\x94\xe4\x3f\xdf\xf0\xe9\x3e\xcc\xcc\x6f\x67\x7e\xf3\x9b\xd9\x9d\xe1\x16\x06\xcf\x4f\x9c\xa1\xc7\xd4\x0f\x12\x68\x0f\x21\x95\x03\x7f\x3a\xff\x84\x3e\xee\xdd\xa6\x10\x92\x59\x30\xf5\xac\x4d\xd6\x54\xc1\x88\x0d\x7d\x78\x38\x36\xf4\xa1\xe3\xb1\xa1\xd8\xf0\x90\xca\x32\xc3\x6a\xa1\x6c\x33\xb3\xac\x15\xc3\xf9\x3e\x2d\x0c\xc6\xa7\x4d\xfd\x0a\xcb\xd8\x56\xdc\x60\xa6\xa5\x23\x53\x9c\x55\x59\xc9\x28\x6a\x66\x2c\xaf\xc7\x97\x17\x2d\x23\x13\xb7\xcc\x4c\x3c\x5f\xb0\x2f\x57\x16\xd4\x8c\x5e\x8a\x57\x16\x98\x29\x0a\x07\x59\xe3\xb6\x77\xb4\x99\x15\xcf\x6a\x7a\xdc\xc8\x57\x55\x1b\x37\xd6\xae\xad\x94\xd4\x1b\x50\x15\x3b\x17\x3b\xbe\x8b\x7c\x60\x78\x68\x90\x5f\xbe\x57\x39\x44\xdf\xb1\xf7\xdc\x56\xd7\x4a\x4a\x4a\x4a\x4a\x4a\x4a\x4a\x4a\x4a\x4a\xea\x25\xca\x36\x7a\xc8\x67\x70\xec\xf1\xaf\x2d\xff\xb8\xad\xe9\xb8\xdd\x3f\x9e\x08\xca\x35\xa5\x4b\x49\x49\x49\x49\x49\x49\x49\x49\x49\x49\x49\x49\x49\x49\xbd\xba\xd2\xb2\x84\xc4\x70\xdc\xc7\x37\x00\xf6\xd4\xd7\xff\x4f\x77\x13\xf2\x7b\xf0\x3b\xf0\x7d\xf0\x55\x70\x1d\x7c\x11\xcc\x83\x4f\x82\x39\xd0\x0f\xfe\xf8\x76\x42\x1e\x82\x6b\xe0\xb3\xa0\x08\x92\x60\x10\x1c\x01\xef\x01\x3b\xc0\xd3\x5d\x84\xfc\x0d\xfc\x02\x7c\x03\x38\xc0\x00\xf3\x60\x0e\x5c\x04\x47\xc1\x6f\x77\x12\x72\x03\x2c\x03\x0d\x1c\x03\x3b\xc0\x7f\x77\x10\xf2\x67\xf0\x03\xf0\x75\x70\x13\x7c\x1e\x94\xc0\x0c\x38\x05\x0e\x82\x7f\xbe\x8d\x90\x5f\x82\x47\xe0\x16\xb8\x0e\x6c\x60\x82\xab\x60\x2f\xf8\x59\x1f\x21\x0f\xc0\xe7\x40\x01\x5c\x04\x69\x90\x04\xe3\x60\x0c\xbc\x0f\x7c\xe7\xad\x68\x2b\x78\x37\xd8\x0f\xf6\x81\x5b\x6f\x21\xe4\x23\xe0\x18\x78\xd6\x4b\xc8\xbf\xc1\xbf\xc0\x59\xf0\xda\x9b\x09\xf9\x2b\xf8\x2e\x78\x00\xbe\x05\xee\x80\x45\xc0\x40\x16\x4c\x03\x0a\x0e\x80\x7f\xbc\x89\x90\x5f\x83\xbb\xe0\x0e\xb8\x0d\xfe\xb2\x1d\xfe\x06\x2e\x98\x00\x63\x60\x14\xc4\xf9\x26\x0c\xf8\x0f\xfa\xea\x19\x78\x02\xbe\x07\xbe\x0d\xbe\x04\xf2\x60\x1a\x50\x70\x10\xec\xef\x11\xfd\xfa\xce\x9e\xf5\x46\x85\xe3\xb0\x72\xd6\x75\x89\x4b\x28\x64\x32\xbb\x62\x96\xa9\x95\xd1\xca\x8e\x43\x0f\xab\x29\xdb\xac\x64\xec\xf4\x8a\xc1\xce\x6a\x25\x46\x5d\x77\x56\x5f\xb6\x22\x26\xfe\x44\xbd\xfc\x59\x96\x63\x26\xe5\xd7\xea\x58\x51\xb7\x58\x44\xdc\x2e\xe4\x28\x33\x4d\x7a\x68\x84\x96\x0b\x45\xea\x04\x66\x71\x31\xe0\x25\xf8\x0f\x43\x31\x71\x7d\x62\x84\xda\x55\x75\xa6\xc2\xcc\x95\x88\x75\x15\x99\x0c\x61\x47\xe4\xe3\xc2\x5d\xfa\xfe\x11\x9a\x2b\xd9\x6a\xca\x30\x0b\x65\x3b\x17\x51\xa6\x92\xc9\x54\x22\x4d\x8f\x64\x95\x01\xaa\xe7\x72\x16\xb3\x6b\x4f\x17\x97\xf4\xa3\x74\x90\x3a\xa4\x0b\x53\x93\x13\x67\x26\xb8\x25\x0a\x53\xc5\x42\xa9\x50\xb7\xe4\x5d\x75\x6f\x68\x6a\x76\x3c\x31\x4b\x47\x3f\x41\x8f\x58\xdc\x96\x6e\x66\x99\x39\xba\x52\xaf\x97\xb8\xe6\x9e\x51\x14\x58\x0c\xcc\xc0\x01\x4a\x2a\x31\x99\x18\x4b\xd3\x7e\x9a\x9c\x9d\x3a\x43\x1d\xe7\xb0\x9a\xd6\x16\x8a\x9e\xe3\x5d\x97\x5e\x38\x9d\x98\x4d\xf0\xbb\xb9\x02\x2b\x66\x91\x94\xb7\xd4\xec\x82\x7a\x5e\x2b\x56\x78\xf2\x08\xfd\x18\x55\x48\xae\x52\xce\xd0\x88\x65\xeb\x26\xa3\xfd\xdc\xc2\xf8\xc9\x29\x51\x3e\x4a\x93\x85\x72\x36\x55\x2c\x64\xd8\xe8\x4a\xcd\x8a\xd7\xa9\x9f\xa2\xbc\xd8\x98\x66\x21\x5b\xba\x1a\x31\x42\x0f\xf1\xfb\xdd\x75\x07\xd0\x3f\xb4\x9f\xff\xf2\x26\x5d\xad\xb5\x8a\x7a\xbf\xfc\xc9\x69\x19\xe6\xb8\xbe\xdb\xf8\xad\xa0\x33\xf8\x79\x94\x46\xe6\xe6\xbd\x9a\x34\x46\x12\xb7\x88\x8e\xd7\xcd\x28\x77\xeb\xc6\x62\x6f\x03\x41\xe6\x79\x46\x1d\x1f\x95\xa1\xb6\xf5\xa1\xd6\x29\xd0\xb6\x2e\xb8\x22\xc1\x1c\xc4\xcf\xfd\xd8\x10\xee\xdc\x5c\xe7\x89\x08\x18\x5a\xdb\x87\xcf\x3b\x52\xe1\xa0\xcd\x74\x4f\xe3\xe8\x69\x71\xd2\xab\xe5\x9f\xf6\xde\xe9\xc2\x23\x8e\x63\x6a\xe5\x3c\xa3\xa2\x2c\x6f\x91\x9a\xe4\x67\x16\x5e\x92\x9e\xb7\x7c\x5f\x2d\x69\x45\x2b\x34\x4a\xe1\x36\xbc\x48\x69\xcc\xad\x0f\xda\x25\xb5\xa1\x7a\x78\x8a\x9f\xd3\x7f\x00\x5a\xc0\x30\x9b\xb6\x3c\xc8\x33\xc2\xcd\xf3\x7b\x73\xf3\xa1\x81\x80\xea\x79\x3e\x61\x55\x96\x39\x69\xe6\xad\xf6\x7d\x15\x59\xa2\x6d\x1b\x19\x6d\x34\xe6\x90\x78\x3c\xa9\x9b\x25\xcd\xa6\x5a\xfb\x02\x7c\x90\xe9\x48\xb4\xf8\x58\xa6\xf6\x65\xe4\x44\x80\xd0\x05\x46\x2b\x16\xcb\x22\x15\xff\x2c\x66\xa2\x3c\x5a\x5e\x31\xb2\x9a\xcd\x30\xc5\xe2\x6f\x89\x95\x6d\xab\x29\xb6\x4c\x66\x55\x8a\x36\xe6\x62\x4c\xd0\x4d\xb3\x9b\x9f\x84\xce\xd7\x0c\x03\x5e\x8c\xd4\xf2\xe2\x24\x4a\x3b\xf9\x17\x89\xcd\x1e\x86\x89\x25\xdc\xf2\x3c\x4b\x5d\x77\x1d\xfe\x0e\x19\xe5\x49\x65\xb6\x1c\x69\xeb\xc3\x5a\xc6\x7a\x91\x50\x23\x83\x97\x4f\x2d\xad\xe9\x25\x55\xbb\x1f\xed\xd4\x28\xae\xa3\xb5\x36\xac\x2f\x66\x82\xf2\xfe\xdb\xce\xfb\x2c\x4b\xa1\xc7\x22\x1d\x23\x54\x33\x69\xe8\x59\xed\x46\xcc\x3a\x1f\x9f\xd3\xfd\xcf\xc1\xb3\xac\x6a\x47\xa2\x7e\x8b\x3b\x7c\x26\x06\xfd\x0e\x3b\x25\x6d\x91\x75\x9e\xcd\x07\x07\x50\x85\x22\x2b\xd7\x87\x62\x54\x0c\x84\x2e\x3e\x16\xc4\x64\xc8\x2f\xbb\x7a\x5f\xc4\xe3\xdc\x63\x16\xa2\x9d\x79\x15\x0e\xc6\x00\x02\x82\x76\x28\xed\x07\x7a\x3d\xce\x45\x8c\x37\xf9\xbb\x8b\x68\xed\xde\xcd\x5d\x45\x69\x38\xa8\x5b\xa3\xb3\x5d\x64\x46\x9b\x2b\xdd\x1a\x8d\xdd\xd7\xb1\x1e\x85\x7e\x10\x06\x4b\x8d\x6e\x23\xaf\xbb\x47\xad\x1d\x07\x94\xc7\x41\x2d\x0c\xba\x78\x05\x34\xc5\x40\x38\x04\x3a\x04\x00\xb1\x71\x81\x46\xd4\xdf\x50\x98\x08\x79\x26\x7c\x97\x8e\x8f\x8a\x67\x8f\xe9\xe5\xf2\xb4\xae\x17\xd1\x2f\x35\x57\x14\x4a\x86\x8e\xf9\x53\xe1\x9e\xf0\x4e\xd5\x09\xef\xc0\xb2\xc2\x8a\x12\x72\x82\xc8\xca\x5b\x2f\xf2\x78\x2f\xa4\xc0\x40\xe8\x07\xf4\x57\xb4\xcc\x62\x86\xff\x4a\x5e\xa9\xa5\xe2\xbb\x51\xf1\xd6\xff\x3f\xc1\x9a\xf6\xa7\x7d\x62\xfd\x1f\xfc\xff\xfe\x8f\xb1\xae\xbf\x07\xbe\x00\x5c\x7f\xcd\x7f\x01\x9c\x03\x69\xd0\x0b\xfe\x8e\x35\xfd\xaf\xc0\x8f\xc0\xd7\xc0\x97\xc1\x4d\x70\x03\x4c\x82\x53\x60\x3f\xcf\x83\x35\xfd\x43\x70\x1b\x7c\x05\xdc\x02\x49\x30\x0c\xf6\x81\xdd\x80\x80\xff\x63\x8d\xff\x3f\xf0\x43\xf0\x08\x5c\x03\x23\xe0\x30\xd8\x09\x7a\xf9\x7d\xac\xef\xef\x83\x7b\xe0\x9b\xe0\xe3\x20\x05\x14\xbe\x37\x80\x75\xfd\x6f\xc0\x43\x70\x1f\xdc\x05\x77\xc0\x6d\x30\x01\x4e\x83\x77\x81\x3f\xa0\xad\x3f\xef\x13\x6d\x7e\xdc\xb7\x19\xbb\x28\x52\x52\x52\x52\x52\x52\x52\x52\x2f\x59\x0d\x4b\xea\xe0\x53\xfe\x52\x78\x17\x3d\x51\x65\x19\xb1\x31\xb3\xc6\x2e\x81\x1e\x55\x55\xb5\x71\xeb\xc6\x71\xc4\x2a\x3e\x35\x33\x39\x9d\xaf\xd6\x3e\xb3\xc3\x1b\x38\x0a\x69\xde\x96\x09\xef\xca\x88\xe2\xe9\x6a\x44\xf7\x52\x5a\xbf\xb1\xc3\x9b\x53\xde\xf7\x76\xcb\x1e\x54\xd0\xa8\xda\x02\x44\xc7\x5a\x92\x65\xad\x09\xb1\xd5\x30\x42\x73\x5a\xd1\x62\xe1\x25\xcc\x48\x78\x09\xd3\xde\x17\x1b\x71\x85\xd8\xdb\x58\xd5\x15\xbc\xc0\x6a\xee\x10\x26\x36\xc5\x1d\x97\x5a\x37\xb1\x45\xeb\x6c\xbe\x27\x3c\x33\xb9\x9e\xee\x6e\x59\x18\x1a\xdc\x4a\xa2\xbe\x30\x0c\xae\x1b\x57\x87\xfc\xe6\x40\x3d\x2d\x5c\x95\x69\x93\x19\x9a\xc9\x22\x4a\xfb\x87\x5f\x12\x71\xa1\x0c\xc0\xbb\x2f\x24\xe2\x3a\x38\xf8\x85\xc7\xd8\xf3\xf5\xc5\x16\x75\x86\x88\xca\xb6\x9d\xb1\x69\x31\xdf\x6d\x87\xb4\xb4\x5e\xb4\xa2\x3e\x8f\xf0\x31\xe2\x6f\x01\xba\x94\xa1\x57\x42\x9b\x68\x0d\x45\xea\x63\xcd\x2b\x12\x78\xaf\xb1\x67\x9d\xd5\x6a\x9e\xd2\x96\x36\x38\x75\xad\xda\x0a\xf4\x76\xf7\xf5\x47\xe6\x0d\xd6\x7c\x4d\x8f\xbf\x1e\x00\x00\xff\xff\x78\xba\x00\xa1\x00\x40\x00\x00")

func templatesDaoPgxTmplSwpBytes() ([]byte, error) {
	return bindataRead(
		_templatesDaoPgxTmplSwp,
		"templates/dao/.pgx.tmpl.swp",
	)
}

func templatesDaoPgxTmplSwp() (*asset, error) {
	bytes, err := templatesDaoPgxTmplSwpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dao/.pgx.tmpl.swp", size: 16384, mode: os.FileMode(420), modTime: time.Unix(1457330820, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDaoPgxTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x57\x5f\x6f\xdb\x36\x10\x7f\xf7\xa7\xb8\x09\xed\x20\x65\x1a\xd3\xbd\x0e\xf0\x86\x26\xb1\xb1\x00\x69\x93\xda\xc6\x86\x61\x18\x02\xc6\x3a\xb9\x5e\x65\x49\x25\x69\xd7\x81\xe6\xef\xbe\x23\x29\x2a\x92\x2d\xd7\xb2\x9b\x61\x7d\x98\x11\x38\x32\xc9\xfb\xf7\xbb\xdf\x1d\x4f\xf3\x45\x9e\x09\x05\x5e\xbc\x50\x5e\xcf\xfd\x98\xcd\xd5\xfb\xe5\x03\x9b\x66\x8b\xf3\xbf\xf8\xf4\xc3\xf4\x3c\x9f\xad\xbd\x5e\xaf\x28\x04\x4f\x67\x08\x2f\xca\x73\x3f\xf6\x81\x5d\x9b\x47\xb9\xd9\x54\xc2\x45\x51\xee\x97\x7b\x18\xbd\xe5\x0b\xdc\x6c\x3c\x92\xc7\x34\xa2\x93\x3d\xf5\x98\x23\x14\x05\xbb\x7a\x7d\x6b\xf7\x40\x2a\xb1\x9c\x2a\x28\xe0\xea\x02\xce\xc8\x1a\xbb\xcc\xd2\xf4\x2e\xcb\x12\xa0\xe3\xbd\xf3\xf3\xf1\x94\xa7\x12\xd4\x7b\x04\x91\x7d\x82\x79\xaa\x32\xe0\x90\xe2\x27\x38\x23\x73\x6c\x6c\xa4\x27\xa4\xd5\xaa\xeb\xc5\xcb\x74\x0a\x92\x64\x8a\x02\xb6\xb7\x61\xb3\x19\x91\x0e\x5f\x2b\x32\xa6\xe8\x57\x00\x7e\xab\xa2\x10\x50\x88\x4c\x04\x50\xd4\x82\x4f\x23\x5c\x87\xf0\x22\x9e\x63\x12\x19\x0c\x86\xfa\x49\xc2\xf7\x64\x18\xe8\xb3\xe2\x02\x56\x1a\x05\x7d\x90\x62\xa3\x47\x73\x96\xd5\x1c\x74\x50\xe8\xf3\x64\x42\xab\x21\x7f\x98\x0e\xd3\x3f\xc2\x94\xfe\x7c\x5b\xb3\x15\x92\x2c\x90\xe6\xea\x40\x60\xbe\xe7\xb1\x31\xf2\x4d\x1f\xd2\x79\x42\xb1\x38\x51\x81\x6a\x29\x52\xbd\x68\x02\x35\xeb\x56\x4e\xa0\xd4\xf6\x08\x61\xbf\x0d\x97\xe0\x08\x1f\x49\x15\xab\x20\x28\xd3\xdd\x7f\x02\x08\x0c\x1a\x75\xa7\x4b\xaf\x48\x2e\xd4\xae\xf5\x36\x3b\x04\x90\x75\x06\xfc\xf1\xe7\xa9\x1c\x90\xbe\xd1\xe5\x48\x20\x89\x05\x7b\x94\xd5\x78\x50\x86\xb4\x4c\x94\x41\x68\xc1\x3f\xe0\x7e\xa9\x57\x21\x65\x3f\xc1\xd4\x81\xa2\x91\xd3\x0a\x22\x8c\x51\x98\x48\xd8\x65\x92\x49\xf4\xed\x72\x9c\x95\x8b\x6f\x71\xad\x7c\x67\xed\x38\x3a\x74\x61\x9f\xd5\xda\x20\x4a\x93\x88\xb2\x64\xe2\xf1\xe6\x77\x19\xb9\xc7\x58\x50\x3d\xed\x63\xe7\x3e\x86\x3e\xb1\xb4\x13\x53\x8f\x8f\xa1\x03\x63\xf7\x04\xe5\x88\xd1\x07\x9e\xe7\xb4\xe9\x97\x0b\xa1\xde\x09\xa0\x51\x60\x8e\xe4\x76\xbf\x22\xfa\x30\x13\x0b\xae\x88\xdb\xad\x94\x72\xc4\x97\xc9\x7c\x8a\x54\x0e\x74\x92\x12\x05\x0f\x08\x4b\x89\x11\xed\xd2\x9f\x44\x6a\xc3\x9c\x3c\x5b\xe6\x11\x57\x48\xbd\x95\xbe\x17\x98\x2a\x69\x6b\x02\xd7\x38\x7d\x2d\x66\xb2\xbd\x2e\xfc\x55\xbb\xe5\x80\x0a\x8d\x8c\xa3\x88\xf9\x14\x8b\x4d\xd1\x5b\xf1\xc4\x20\xdf\x58\x3e\x8d\xb2\xdb\x68\xb7\x92\x46\x7f\x97\xa8\x69\xd3\x1a\xad\xca\xcc\xb6\x7a\x57\xfe\xbe\x54\x99\x40\x1b\x50\x75\xd9\x04\x30\x24\xb7\x2e\x1e\x1b\x46\xe1\x6f\xd0\x12\x97\x5c\xd2\x09\x3f\x6f\x2b\x9b\x0e\x97\x84\xf6\x53\x7e\x4c\xb4\x27\xde\x78\x70\x33\xb8\x9c\xc0\x19\x0c\x47\xb7\x6f\xb4\x3e\x36\xe1\x0f\x89\xcb\xe3\x6f\xbf\x0c\x46\x83\x9a\x15\x3e\x93\x2c\x7a\x60\xbf\xf2\x64\x69\xe9\xf6\x33\xdc\x5c\xbf\xb9\x9e\xc0\x0f\x5e\x9d\x2f\x9f\xed\x67\x36\x5a\x76\x75\xc1\xde\x2d\x51\x3c\x9a\x95\x8f\x54\x38\x79\x10\x68\xb4\xbe\x04\x91\xc9\xba\x1d\x93\x10\xd4\xda\x36\xd0\xc9\xfa\x3f\xc5\xa7\x3b\x48\x6a\x7d\x1a\x3c\x63\x5d\x72\xc7\xb3\x26\x84\x4c\x44\x28\x2e\x1e\xa1\x56\x26\x21\x24\xf3\xc5\x5c\xe9\x25\xda\x8f\x63\x89\xe6\xb9\xf3\x05\xf4\x6c\x18\x7a\x6e\x3e\x70\x4e\x52\x17\xf6\xbc\x5a\x13\xd6\x96\xbe\xeb\x03\x8d\x87\x6c\x9c\x0b\xf2\x31\xf6\xbd\xdb\xd1\xd5\x60\x04\x17\xbf\xc3\x4b\x09\x5e\x15\x5f\x50\x56\xa9\xd3\x68\x23\xfc\x09\x5e\x1d\xd0\x66\xd3\xf8\x32\xd2\xaa\x8c\xcc\xb6\xa2\x12\x9f\xc3\x9a\x6e\x87\xc3\xf1\x40\xab\xf2\x1c\xa8\x41\xbd\xe3\xd2\xb5\x16\xba\x3b\xae\x59\x29\x8e\x07\x6d\xc3\xd2\xf6\x15\xd4\xec\xdf\x1d\xe6\x8b\xe7\x21\x57\x97\x02\xfc\x9f\x6a\x5f\x23\xd5\x5c\xbf\x39\x81\x64\x7b\xc6\xc4\xe3\xb8\x57\xbd\x75\x6d\x71\xb0\x4e\xc1\x31\x5f\xa1\x9f\x99\xd5\xdd\x7b\xdf\xd0\xa1\xc4\x43\x63\x44\xd3\x29\x46\xf2\xda\xce\x19\x3b\x2f\x13\xb6\xb2\xec\x14\xe2\x67\x25\x2a\x80\x89\xc4\x7d\x67\xed\x98\x42\x67\xdd\x7c\x74\xd8\x57\xaa\x86\x76\x6f\x9b\xf7\xd1\xe9\x9e\x6b\x03\x5a\x57\x77\xff\x2b\x89\x6e\x51\x38\x80\x3a\x60\xde\x28\xc3\xa2\xb0\x92\xe3\x77\x37\x77\xb3\x75\x35\x4e\xd5\x8b\xd1\x16\x9a\x54\x0b\x2a\xf8\x5c\xff\x1b\x6c\xf5\xbc\x3b\x81\x39\x17\xe8\x7b\xed\xec\xb9\xb7\x16\xbc\x90\x6c\x56\x6c\xad\x14\xed\x7b\x89\x74\x07\xea\x15\x75\xbf\xdb\x6f\x07\x34\x7c\xfa\xfa\x28\xa3\x08\xc2\x43\xb3\x68\x16\x30\xc6\x1a\x15\xd3\xdf\x32\xdf\xcc\x29\x55\x29\x8d\x86\xb8\x3b\x67\xeb\x77\x87\xcf\x27\xc4\xb1\xf0\x40\x42\xb6\x92\x61\xa5\xfe\xcd\x64\x58\x0b\x5f\x96\x8c\xe7\xce\x45\x67\x50\x9f\x8a\xe9\xc8\x6a\x3d\x8d\xf3\xf7\xf5\xa6\x6b\xa3\xa3\x9e\xfb\xb5\x71\xec\x48\x38\x4e\x63\x5c\x3b\x14\xa7\xe6\xf8\x9f\x00\x00\x00\xff\xff\x58\xc8\x61\x25\x26\x14\x00\x00")

func templatesDaoPgxTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDaoPgxTmpl,
		"templates/dao/pgx.tmpl",
	)
}

func templatesDaoPgxTmpl() (*asset, error) {
	bytes, err := templatesDaoPgxTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dao/pgx.tmpl", size: 5158, mode: os.FileMode(420), modTime: time.Unix(1457330805, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDaoSqlxTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x56\xdf\x6f\xe2\x46\x10\x7e\xe7\xaf\x98\x5a\x4d\x64\x52\x64\xd2\xd7\x4a\xb4\x2a\xc1\xa8\x48\x49\x50\x30\x6a\x55\x9d\x4e\x27\x83\xc7\x39\xdf\xd9\x5e\xdf\x7a\x9d\x18\xf9\xf8\xdf\x6f\x76\xbd\x36\x06\x1c\x41\x42\x4e\xa7\x93\xc2\x93\xb5\xf3\xeb\xfb\xbe\x9d\x19\x36\x88\x12\xc6\x05\x18\xf7\x81\xf8\x98\x2d\xac\x25\x8b\xfa\x9f\x22\x16\x70\x16\xf7\xd3\x2f\x61\x6e\x74\x5a\x1c\xb2\x05\xf2\x05\x67\x9e\xe8\x63\x8e\x51\x12\xba\xbc\x1f\x31\x0f\xc3\x8d\xb3\x1f\x09\xa3\xd3\x29\x0a\xee\xc6\xf7\x08\xbf\xea\xe3\x3f\x06\x60\x4d\xd4\x67\xba\x5e\xd7\xbe\x45\xa1\xed\xda\x86\xde\xad\x1b\xe1\x7a\x6d\x50\x3c\xc6\x1e\x79\x76\xc4\x2a\x41\x28\x0a\x6b\xf4\xf7\xb4\xb4\x41\x2a\x78\xb6\x14\x50\xc0\x68\x08\xaa\xb6\x65\x6b\x28\x0e\xa1\x86\x75\xa3\xb8\x1f\x60\xe8\xa9\xda\x63\xf9\x25\x4b\xfb\x59\xbc\x04\x33\x15\x8c\x23\x5c\x50\xfd\x4d\xe2\x2e\x8c\x83\xd8\x1b\xae\xe8\x50\xc5\x59\xf2\x18\xbe\x82\x8c\xb8\x72\x53\xf2\x30\x13\xa8\x8d\x73\xc2\x55\xc5\x99\x2a\x91\xa3\x70\x6d\xce\x7b\x80\x9c\x33\xde\x85\xa2\x03\xf4\x7b\x70\x39\xe4\xd0\xe6\xa8\xcc\xa4\xb8\x04\x6a\x38\xf6\xb5\x7d\x35\x87\x0b\x18\xcf\xa6\x37\xca\x7d\xee\x2e\x42\xed\x09\xff\xfd\x63\xcf\xec\x06\x08\xf7\x3e\xb5\xbc\x85\xf5\xaf\x1b\x66\xd2\x3c\x80\xbf\xe0\x7a\x72\x33\x99\xc3\xef\x86\xca\x4a\x08\x64\x56\x45\xd7\x1a\x0d\xad\xbb\x0c\xf9\x6a\xc6\x1e\x73\x93\xea\xf5\x20\xe9\x6a\x30\xce\xd2\x8d\xcd\xf3\xbc\x0b\x2a\x8a\xa3\xc8\x78\x0c\xe7\xb9\xa2\xd0\x21\x45\x4f\x51\x6d\x9e\xb7\xeb\xd6\x03\x91\xc3\x85\xec\x34\x6b\x9e\x9f\x2c\x22\x7c\x47\x15\x9b\x52\x8a\xfc\xa0\x88\x2f\xd4\xd0\x09\x83\x25\x3e\xbf\xfd\x7a\xc0\xb8\x87\x7c\xb8\x82\x20\x16\xc8\x7d\x77\x89\x05\x1d\x86\x41\x14\x08\x79\x44\x76\xdf\x4f\x51\x7d\x93\xca\xef\xde\x1f\xa3\x73\x2e\xa9\x46\xee\x67\x7c\xda\xff\xb2\xfb\xba\x8a\x97\x52\x07\x7e\x4d\xe7\x17\xca\x6a\x68\x3c\x55\xa5\xdf\x06\x40\xfb\xc5\x72\x12\x4e\x6c\x7c\xd3\x98\xce\x46\xf6\x0c\x86\xff\xc3\x59\x0a\x46\xad\x44\x89\x8c\x24\xd7\x19\x4b\x2d\xfe\x84\xcb\x03\xd9\xca\x4b\x3f\xf3\x64\x2a\x15\xb3\x9b\x48\x2b\x79\x38\xd3\x74\x3c\x76\x6c\x99\xca\xa8\xe4\xaf\x52\xa9\xde\x60\x8f\x69\xaf\x7d\x36\xeb\x9e\xaa\x4a\x4a\x2f\x52\x22\x0e\x42\xda\x78\xba\xa9\xca\x9e\xd2\xc9\x3c\xf4\x91\xab\x94\xd6\x55\xc8\x52\x34\xcb\x50\x9f\xe9\xc3\x5b\xcc\x85\xd9\x6d\xe0\x0d\x64\xd5\x18\x1f\xcd\xb6\x7b\xed\x6e\x68\x51\x37\xdb\x25\x44\x95\xa7\xd1\xe3\xc1\xc6\x8b\x10\x56\x8e\xfb\x28\x2b\xcb\xba\x76\xcf\xe9\xaa\xdd\x24\xa1\xdd\x6e\x12\x89\x60\x4b\x94\x9a\x1b\x65\x79\x95\x79\x39\x6a\xf1\xbc\x8d\xcf\xcf\x3d\x3e\xd5\x3e\x7e\x1b\x9c\x6a\x70\xea\xa7\xd3\xce\x00\x35\xe7\xc7\x71\x1f\xd0\x64\xea\x74\x8f\x48\xd9\xca\x9a\xb7\xbc\x35\x52\x02\xbd\x74\x12\xa7\x48\x6f\xb6\x8d\x1c\xba\x6e\xb9\xbf\x02\x65\x35\x99\x06\x06\x18\xa6\xf8\x94\x6f\x96\x78\xae\xa0\xf2\xfa\xb5\xb1\x6e\x19\xf6\x5d\xac\x34\xca\xed\x68\x77\x1e\x11\x2f\x87\x2e\x2b\xc8\x64\xc7\x13\xa8\x23\x8e\xa3\x51\x29\x74\x84\xe8\x5b\x9b\xa1\x28\xca\x48\xe7\xee\xba\x7a\xc6\x42\x73\x39\x94\x83\xff\x61\xff\xff\x44\x9a\x3d\x7a\x1a\x2f\xcb\xc9\x60\x5b\x93\x31\xd0\x3d\x57\x11\xdc\x56\x8a\xc6\xd1\x25\xfe\xfb\x4d\xd6\xfe\x94\x69\xd2\xac\x2e\xf7\x00\xcd\x1d\x8a\x65\xd4\x2b\x51\x3c\x1a\xeb\xe6\xe6\x9f\xdb\x5b\x27\xdd\x10\xad\xac\x1f\x79\x37\xcf\xe5\x7b\xd2\x55\x3d\xc9\x75\x0b\xf4\xb7\x00\x00\x00\xff\xff\xee\xcd\x35\x13\x8a\x0e\x00\x00")

func templatesDaoSqlxTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDaoSqlxTmpl,
		"templates/dao/sqlx.tmpl",
	)
}

func templatesDaoSqlxTmpl() (*asset, error) {
	bytes, err := templatesDaoSqlxTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dao/sqlx.tmpl", size: 3722, mode: os.FileMode(420), modTime: time.Unix(1457324095, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/dao/.pgx.tmpl.swp": templatesDaoPgxTmplSwp,
	"templates/dao/pgx.tmpl": templatesDaoPgxTmpl,
	"templates/dao/sqlx.tmpl": templatesDaoSqlxTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"dao": &bintree{nil, map[string]*bintree{
			".pgx.tmpl.swp": &bintree{templatesDaoPgxTmplSwp, map[string]*bintree{}},
			"pgx.tmpl": &bintree{templatesDaoPgxTmpl, map[string]*bintree{}},
			"sqlx.tmpl": &bintree{templatesDaoSqlxTmpl, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
