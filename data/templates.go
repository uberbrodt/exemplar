// Code generated by go-bindata.
// sources:
// templates/dao/.pgx.tmpl.swp
// templates/dao/.sqlx.tmpl.swp
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

var _templatesDaoPgxTmplSwp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9a\x5b\x6c\x1c\x57\x19\xc7\x4f\xeb\x24\x36\xa1\x69\x9a\x58\x42\xaa\xc4\xe5\x74\xea\xa4\xb3\x89\x3d\x9b\x34\x40\xdb\x08\xb7\xf5\x65\x4d\xad\xda\xb1\x63\xaf\x5b\x48\x62\xa2\xf1\xee\xd9\xcd\x26\xb3\x33\xdb\x99\x59\x7b\xdd\xed\x90\x87\x12\x5e\xe0\x01\xa1\x04\x55\xaa\x04\x48\x3c\x44\xa2\x55\xfa\x44\x11\x0f\xd0\x82\xb8\x14\x1e\x00\xa9\x84\x88\x97\x22\x10\x0f\xdc\xca\xa5\x50\xa8\xb8\xa8\xfc\xcf\x9c\x99\xd9\xd9\xcd\xae\x77\x9c\xba\x71\x50\xcf\xa7\xfc\x32\x7b\xe6\x5c\xe6\x7c\xdf\xf9\xce\xcc\xf9\x8e\xcf\xd2\x81\x47\x26\xa7\xe9\x3d\xda\x21\x02\xb9\x8d\x90\xde\xfe\x9e\x47\x5f\xfe\xe6\xfb\x7b\xcb\xfd\x84\xe4\x96\x6c\x2b\xef\x92\xae\x22\xca\x0d\xb9\xb6\x5e\x72\x4b\x85\x55\xcd\xb0\x72\xba\xd1\xa6\xdc\x27\x45\xc1\xf4\xac\x6d\x9d\x66\x39\xd7\x49\x57\x98\xed\x58\xa6\x6e\xa4\x59\x8d\x95\x2b\x86\x6e\x0f\x15\xad\xf4\xca\x19\xa7\x92\x4b\x3b\x76\x2e\x5d\x2c\xb9\xa7\xaa\x4b\x5a\xce\x2a\xa7\xab\x4b\xcc\x16\x95\xc3\xa2\x69\xd7\xbf\xba\xcc\x49\xe7\x75\x2b\x5d\x29\xd6\x34\x17\x37\xba\xf7\x56\x8a\x94\x77\xa0\x54\xdd\xc2\xd0\xbd\x3b\xc8\xa1\xbb\x0f\x1e\xe0\xc9\x3b\x95\x3b\x68\xff\xee\x85\xcd\xee\x95\x14\x29\x52\xa4\x48\x91\x22\x45\x8a\x14\x29\x52\xae\xa3\xb8\x95\x1e\x72\x16\xd7\x9b\x83\x74\x26\xb8\xde\xd4\x72\xdd\x12\x5c\x47\x82\xeb\x58\x4b\x7e\x4f\x70\x7d\x20\xb8\x3a\x2d\xf9\x52\xa4\x48\x91\x22\x45\x8a\x14\x29\x52\xa4\x48\x91\x22\x45\x8a\x94\xcd\x13\x3d\x4f\xc8\xfe\x5e\x42\x8e\xf7\x11\xff\xef\xff\x61\xfc\xff\xb7\x9d\x84\x5c\x06\x3f\x03\xcf\x83\xf3\xe0\xb3\xe0\x33\xe0\x18\x48\x01\x15\xdc\x05\xfe\x71\x2b\x21\x5f\x01\x1e\x58\x00\xd3\xe0\x5e\xb0\x1b\xfc\x67\x07\x21\x7f\x05\x7f\x00\x57\xc0\x0b\xe0\x1b\xe0\x3c\x30\xc1\x47\xc1\x47\x40\x0a\xdc\x05\xf6\x82\xd7\x6e\x21\xe4\x69\x70\x0c\x8c\x81\x21\xd0\x03\xae\xbc\x9b\x90\x97\xc0\x25\xf0\x29\x70\x1a\x1c\x07\x47\xc1\x7d\xe0\x7d\xa0\x1f\xfc\x7e\x3b\x21\x2f\x82\x0b\xe0\xd3\xe0\x13\xe0\x63\xe0\x41\x70\x3f\x18\x06\xaf\xbc\x8b\x90\xcf\x83\xc7\xc1\x2c\xf8\x10\xd8\x0b\x28\xd8\x0d\x6e\x03\x3b\xc1\x45\xd8\x65\xb1\x4f\xd8\x67\x01\x3c\x08\x7e\x0d\x7b\x9d\x00\xdb\xc1\x56\xf0\x83\x6d\x84\x7c\x0f\x7c\x17\xdc\x0e\x7e\xbc\x95\x90\xaf\x83\x27\x81\x07\x9e\x00\x25\x70\x18\x7c\x10\x1c\x02\x1f\x00\x7f\xde\x42\xc8\xef\xc0\xb7\xc1\xd3\xe0\x0c\x28\x81\x53\xe0\xb9\x1e\x42\xf2\x60\x1a\xdc\x0e\xfa\xc1\x6e\xf0\xe6\xcd\x84\xfc\x14\xfc\x08\xfc\x10\x7c\x0d\x9c\x03\x67\xc1\x22\xb8\x07\x50\xf0\x97\x9b\x08\xf9\x13\x78\xb5\x69\x03\xc6\x23\x14\x62\x33\xb7\x6a\x9b\xd4\xc9\xe9\x66\xbd\x4e\x07\xb4\x79\xd7\xae\xe6\xdc\xec\x6a\x85\x1d\xd1\xcb\x8c\x7a\xde\x9c\xb5\xe2\xa8\x36\xfe\x4b\xf9\xe5\xf3\xac\xc0\x6c\xca\xd3\xda\x98\x61\x39\x4c\x15\xb7\x4b\x05\xca\x6c\x9b\xde\x31\x4c\xcd\x92\x41\xeb\x61\xb3\x48\x0c\xfa\x19\xc1\xc3\x50\x4d\xa4\x0f\x0f\x53\xb7\xa6\x1d\xad\x32\x7b\x55\x75\x1e\x43\xa1\x8a\x68\x47\x94\xe3\x82\xbb\x74\xff\x30\x2d\x94\x5d\x6d\xbe\x62\x97\x4c\xb7\xa0\x2a\x33\x13\x13\xf3\x99\x2c\xdd\x93\x57\x06\xa9\x55\x28\x38\xcc\x8d\x9e\x2e\x92\xf4\x7e\x7a\x80\xd6\x49\x82\xa6\xa6\x26\xa7\x27\x79\x4b\x14\x4d\x19\xa5\x72\xa9\xd1\x92\x9f\x4a\xde\xd0\xcc\xdc\x78\x66\x8e\x8e\x7e\x9c\xee\x71\x78\x5b\x96\x9d\x67\xf6\xe8\x6a\xa3\x5f\x22\xcd\x2d\xa3\x28\x68\x31\x6c\x06\x06\x50\xe6\x33\x53\x99\xb1\x2c\xdd\x47\x27\xe6\x66\xa6\x69\xbd\x3e\xa0\x65\xf5\x25\xc3\x37\xbc\xe7\xd1\x47\x1f\xca\xcc\x65\x70\xd7\x44\x72\xc6\xce\xea\x45\x3a\x50\x28\x31\x23\x8f\xac\x61\xfa\x00\x55\x48\xa1\x6a\xe6\xa8\xea\xb8\x96\xcd\xe8\x3e\x5e\x7b\xa2\x64\xe2\x61\xa2\x7a\x8a\xf2\xd4\xbc\x51\xca\xb1\xd1\x55\x64\xfa\x55\x35\x7f\x4c\x9f\xa0\xbc\xe6\x98\xee\xa0\x58\xb6\xa6\x56\x68\x94\x1d\x0e\xbb\xe7\x0d\x62\x78\xe8\x3e\x7e\x90\x26\x5b\x8b\x94\xa2\xd0\x98\xd9\x05\x3d\xc7\xea\x5e\x60\x35\x7e\x2b\x1c\x0b\xfe\x3b\x45\xd5\xe3\x8b\x7e\x67\x9a\x1d\x89\xb7\x88\x71\xb7\xec\x14\xb7\xea\xb5\xb9\xde\x35\xf8\x98\x6f\x1c\x6d\x7c\x54\x7a\xda\xa6\x7b\x5a\x27\x3f\xdb\x3c\xdf\x52\xc3\x37\x10\xff\x1d\xb8\x86\xb0\xe6\xc6\xd9\x4e\x0c\xfe\xc1\x44\x26\x7c\xab\xf3\x14\xf6\x79\x4b\xd6\xf1\xd6\x92\xa0\xde\x9a\xd6\x6d\x9b\xe9\xc5\x3c\x3b\xd1\xa8\x34\xcf\xd9\x68\x6c\x74\xbb\xe8\x68\x9a\x86\x11\x4a\x60\xca\x19\x93\xf1\x5a\x98\xff\x98\x3b\x45\x51\x99\x1e\x5f\x8c\xf9\x58\x52\x63\x05\x5d\x5e\xd6\x0d\x27\xa6\x08\x7a\xcf\xcc\x3c\x1d\xf2\x1a\x33\x76\x59\x6b\x1a\x3d\xb4\x14\x94\xb4\x75\xb3\xc8\xe8\x00\xef\x23\x5e\xa5\xa2\x08\x77\x2c\x74\x1b\xbf\x1c\xbf\x11\xde\x3c\xbf\xd7\xd4\x45\x74\xc1\x57\x95\xd5\x58\x6e\x04\x0a\xb4\x37\x99\xba\x4c\xdb\x2a\x92\x6a\x6e\xac\x4e\xd2\xe9\x09\xcb\x2e\xeb\x2e\xd5\xdb\x57\xe0\x53\xcc\x42\xa6\xc3\x67\x32\x75\x4f\xa1\x24\xc6\x89\x2e\x31\x5a\x75\x58\x1e\xb9\xf8\xe7\x30\x1b\xf5\xa1\x79\xb5\x92\xd7\x5d\x06\xfb\xe2\xff\x32\x33\x5d\xa7\xc5\xb5\x6c\xe6\x54\x0d\x17\x2f\x62\xbc\x9d\x5b\x5e\x6d\x41\x16\xa6\x87\x5e\xa9\xc0\x8a\x6a\x54\x16\x3f\x52\xb4\x93\x7d\x91\xd9\x6a\x61\x34\xb1\x8c\x5b\xbe\x65\xa9\xe7\xad\xc3\xde\xb1\x46\x79\x96\xc9\x56\xd4\xb6\x36\x8c\x0a\x36\xaa\xc4\x94\x0c\xbf\x3c\x51\x5e\xcb\x17\x2a\xba\x9f\xea\xa4\x14\x97\xbd\x91\x0e\xeb\xf3\x99\xb0\x7e\xf0\xa9\xf3\x97\x64\xf3\x18\x31\xb5\xa3\x87\xea\x36\x8d\x3d\xab\xdd\xeb\x64\x9d\x8f\x2f\x58\xc1\x52\xf0\x08\xab\xb9\x6a\x2a\xd0\xb8\xc3\x12\x31\x1c\x77\xb4\x53\xd6\xcf\xb0\xce\xef\xf2\x03\x83\xe8\x82\xc1\xcc\xf0\x71\x7c\x1c\xfc\x89\x90\x60\xa5\x20\xde\x85\x3c\x99\xe8\x6b\x91\x4e\x73\x8b\x39\xf0\x76\xe6\x77\x38\x9c\x03\x70\x08\xda\xa1\x76\xe0\xe8\x0d\x3f\x17\x3e\xde\x62\xef\x04\xde\x9a\xdc\xcc\x89\xbc\x34\xee\xd4\x57\x7b\x67\x3b\xcf\x4c\xb5\x76\xfa\x6a\x6f\x4c\xde\xc7\x86\x17\x06\x4e\x58\xaf\xa3\xe9\x20\x33\x89\xe7\x25\x7b\x54\x77\x3f\xa0\xdc\x0f\x22\x37\x48\xf0\x9a\x6f\xf1\x81\xb8\x0b\x74\x70\x00\xe2\x22\x01\x25\xb4\x69\xfe\x06\x0b\x3e\x3c\xfc\x5b\x83\x72\x58\x97\x8e\x8f\x8a\xc7\x8f\x59\xa6\x39\x6b\x59\x06\x86\x26\xac\x10\xff\x52\xad\x59\x21\x32\x5f\xa9\x5c\xb1\xf0\xce\x55\xb8\xf5\xfc\x9f\xda\xa4\x7f\x61\x79\xd1\x8a\x12\x33\x9c\x28\xca\x2d\x26\xca\x38\xbc\xb3\x61\x03\xb1\x03\xf4\xa7\xf5\xdc\x99\x1c\x3f\x25\xaf\x44\xb9\x58\x68\x2a\x7e\xfc\x7f\x19\x31\xec\xab\x41\xfc\x1f\xfe\xfd\xfe\x0d\xc4\xf5\xff\xe4\xb1\x3d\x78\x0e\x7c\x15\xac\x82\x41\xd0\x0f\xfe\x8d\xd8\xfe\x75\xf0\x77\x70\x11\x3c\x05\xbe\x00\x2e\x80\x23\xc1\x1e\xc0\x7b\xc1\x4f\x10\xcf\x3f\x0f\xbe\x08\x9e\x02\x0f\x07\xf1\xfe\x61\x70\x1f\xf8\x0d\xe2\xfa\x5f\x82\x4b\xc0\x04\x27\xc1\x22\x38\x01\x7a\xc1\x36\x70\x19\x31\xfd\x79\x70\x0e\x3c\x09\x18\x18\x05\xfb\xc1\x3e\xf0\x2b\xc4\xf3\x17\xc1\x31\x30\x02\x52\xe0\x4e\xa0\x80\xdf\x22\x86\xff\x05\xb8\x02\x7e\x0e\xce\x82\x0c\x18\x07\xef\x01\xbd\x60\x2b\xf8\x2f\xf4\x7f\xbd\x4f\xd8\xe1\x8f\x7d\xd7\x71\x03\x46\x8a\x14\x29\x52\xa4\x48\x91\xf2\xf6\x4a\xb8\xc0\x6a\xbd\x36\x85\xb1\xe1\xf2\xf9\x64\x7c\xd7\x3a\x83\x78\x5c\xec\x49\x58\xad\x6b\xfc\x94\xd8\xc7\x0b\xb7\x8e\xc6\x33\x53\x99\x6c\x26\xdc\x38\xa2\xb1\x9d\x23\xbe\x98\x0d\xf6\x8e\x4a\x79\x44\x06\x03\xd1\x16\x51\xd9\x5f\x53\x8a\xa5\x67\x7c\x7d\x99\xa2\xe3\xcc\x60\x2e\xcb\xd6\x54\xab\xfd\xc2\xb4\x79\x2f\xc8\x5f\xdf\x62\x71\x49\xba\x69\x24\x1e\xc8\x37\x5c\x84\x62\x2e\xdf\x73\x3c\x3a\xd5\x59\xbb\xab\x22\x8d\x0a\xaf\x93\x69\x44\x1a\x61\xba\x39\xdc\xe0\x37\x07\x1b\x79\x4d\x4f\x9e\xb5\x59\x45\xb7\x99\xaa\xb4\x5f\xd6\x9f\xcc\xfb\xaa\x2b\x83\x30\x6c\xea\x6d\xb7\x70\x27\xfb\xc6\x6d\x8a\x67\xf1\x90\xea\x31\x3a\xc0\x37\xfe\xa8\x52\xca\x2b\x22\xa6\x5b\xc1\x32\x3b\xb8\xc9\x43\xb6\xab\x76\x05\x1b\xcb\xf5\xb6\xf1\xcd\x7a\xbc\xaf\xcb\xbe\x90\x95\xe2\xdb\x65\x4d\xd6\xaa\xd7\xc5\xbe\x0d\x86\x77\xb6\x58\x8b\x1e\x1c\xdf\xd0\x54\x48\x57\x2b\x2d\xf8\x6d\x84\x7e\x98\xcc\x0d\xbb\x7a\x61\x5c\xb3\x6b\x51\x4c\xec\x4d\xad\xa9\x18\xaf\xd0\x55\xb9\x49\xbf\x9d\x0d\x56\xae\xe3\x14\x4b\xae\xe9\x26\x4c\x3a\xe1\x2b\x6d\x26\xdd\x46\xba\x51\x07\x3b\x5f\x37\xdb\x6e\x96\x71\x85\xbf\xb6\x35\xee\xc6\xba\x72\x77\x03\x47\x9f\x3e\x1e\xff\xef\xda\x4a\xc8\xc1\x6d\xc4\x8f\xff\xc3\xf3\xfd\x6f\x22\xc6\xbf\x00\x8a\x60\x04\xec\x02\x3b\xc1\xad\x3b\x45\xac\xcf\x63\xfc\x29\xf0\x30\x78\x19\x71\xfc\x33\xe0\x1c\x58\x00\xf3\x60\x0e\x7c\x18\x7c\x0b\xf1\xbb\x01\x54\xb0\x17\xec\x01\x5f\x42\xcc\x3e\x0d\x1e\x02\xdb\x41\x1f\xe8\x05\x27\x10\xaf\xef\x02\xff\x42\x3c\xfe\x7d\xf0\x22\x78\x01\x3c\x0b\x3c\xb0\x02\x96\xc1\x12\xb8\x1b\x0c\x82\x2d\xe0\x15\xc4\xeb\x97\xc0\xb3\xe0\x19\xf0\x5a\x2f\x21\x17\xc1\xe7\x82\xbf\xd3\x3f\x02\x16\xc0\x28\xe8\x07\x3b\xc0\x2d\xe0\x0d\xe8\xfc\x12\xf8\x0e\xf8\x32\x78\x1c\xe4\xc0\xa1\x6d\x81\x3d\xa4\x48\x91\x22\x45\x0a\x97\xf5\xfe\x55\x57\x1e\x21\x5a\xab\x4f\xff\xd7\x07\x3b\xe4\x11\x22\xe9\x69\xf2\x08\x91\x3c\x42\x94\xfc\x28\x4f\x68\xa3\x1b\xca\x3c\xed\x8d\x93\xc0\x20\x1d\x37\x74\xa2\xb3\x02\xd7\xe3\x35\xd5\xfc\x29\x0c\x0f\x4b\x25\x31\xca\x88\x61\xc0\x35\xd6\x3e\x2d\xd5\xe2\x23\x1b\x3b\x87\x36\xe8\xc0\xd7\x7a\x95\xb8\xb1\x3f\x31\xeb\x1c\xc1\xee\xa7\xdd\x12\x8c\xd9\xff\x02\x00\x00\xff\xff\x6c\xef\xd8\xec\x00\x50\x00\x00")

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

	info := bindataFileInfo{name: "templates/dao/.pgx.tmpl.swp", size: 20480, mode: os.FileMode(420), modTime: time.Unix(1459819281, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDaoSqlxTmplSwp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\xda\x7b\x6c\x14\xc7\x1d\x07\xf0\x31\x06\xdb\x3c\x6c\x68\xa9\xd4\x96\x0a\x34\x5c\x0d\xdc\x19\x7b\xcf\x94\x52\x28\xad\x69\x31\x3e\x83\x91\x8d\x8b\xef\xca\xa3\x94\xa2\xf5\xed\xdc\xb1\xb0\x77\x7b\xde\xdd\xc3\xe7\x1e\xd7\x56\x85\xd2\x3f\x90\x50\x1f\x12\x56\xff\xa8\x5a\xf1\x90\x20\x91\x20\x46\x09\x0a\x16\x91\x42\x10\x48\x11\x44\x8a\x22\xa1\x3c\xa4\x28\x08\x85\x20\x08\x09\xff\x44\x49\x14\x43\x92\xef\x3e\xee\xe1\xf3\x99\xb3\x85\x6d\x2c\x31\x3f\xe9\xa3\xbd\xdd\x9d\x9d\x9d\x99\x9d\xd9\xbb\xbd\x9d\xce\xfa\x2d\x2d\x6d\x74\xa5\xb0\x9c\x20\xe6\x10\xf2\xe8\x7b\xa5\x5b\x83\xaf\x2d\x28\x8f\xcc\x25\x24\xd8\xa9\xa9\x92\x41\x8a\x86\x9d\xae\xce\xd0\x44\xd9\x90\x43\x3d\x82\xa2\x06\x45\xa5\x40\xba\x3f\xd9\x09\xbd\xbf\xd1\xd4\x3d\x2c\x68\xe8\xde\x18\xd3\x74\x35\x2a\x2a\x5e\x96\x60\x91\x98\x22\x6a\x75\x61\xd5\xdb\xbd\x57\x8f\x05\xbd\xba\x16\xf4\x86\x65\x63\x77\xbc\x53\x08\xaa\x11\x6f\xbc\x93\x69\xf6\xc1\xe9\xa4\x5e\xc3\x5a\x1a\x4c\xf7\x4a\xa2\xea\xd5\xbb\x94\x84\x60\x60\x4b\xf1\xe2\xf2\xe0\xf1\xec\x45\xdc\x08\xd5\xad\xaa\x24\xcb\x7f\xb2\xac\xde\x5c\xfd\xb1\x6b\x21\x9d\xfb\xdd\xdf\x3e\xed\x52\xf1\xe0\xc1\x83\x07\x0f\x1e\x3c\x78\xf0\xe0\xc1\x63\x02\xc3\x88\x95\x92\x3f\x63\x39\xc5\x59\x5f\xee\x2c\x4b\xf2\x96\x53\x9d\xe5\x1e\x67\x59\x99\xb7\xbf\xd4\x59\xce\x76\x96\xfb\xf2\xf6\xf3\xe0\xc1\x83\x07\x0f\x1e\x3c\x78\xf0\xe0\xc1\x83\x07\x0f\x1e\x3c\x9e\x5e\x88\x12\x21\x57\x2b\x08\x39\x33\x9d\x58\xef\xff\xd3\xcf\xff\x1f\xe0\x41\xfe\x1c\x9c\x84\x13\x70\x08\x24\xf8\x3d\xec\x80\x45\xf0\xa0\x8a\x90\x4f\xe1\x10\x84\x41\x80\x1f\xc1\x0f\xe1\x07\x70\xa6\x92\x90\x83\xb0\x0d\x7e\x0a\xf3\xa1\x1c\x6e\xce\x22\xa4\x1f\x0e\x83\x06\x0a\xfc\x0e\xb6\xc1\x56\xb8\x39\x93\x90\x4b\x70\x04\x92\xd0\x03\x09\xf8\x66\x06\x21\x37\xe0\x2c\x1c\x06\x03\xfe\x00\x2b\xe1\x3b\x70\x0b\x75\x78\x1d\x2e\x43\x1f\x9c\x9d\x6e\xd7\xab\x16\x96\x42\x17\xea\xb9\x1e\x1e\x96\x13\x72\x1b\xfa\xe0\x79\x78\x0e\x7a\x61\x37\x74\x82\x08\x6d\x30\x1f\xbe\x0f\x9f\x95\x11\xf2\x26\x5c\x80\xe3\xb0\x1f\xb6\xc0\x12\xa8\x82\x0a\xf8\x7c\x1a\xd2\xc1\x7d\xf8\x18\xee\xc1\x09\xf8\x1f\x74\xc1\x7a\x58\x05\x3f\x83\x15\xf0\x60\x2a\xda\x0e\x5e\x85\xa3\xf0\x77\xf8\x1b\x1c\x04\x2f\xd4\xc0\xc3\x52\x42\xde\x85\x2b\x70\x0a\x8e\xc1\xbf\xe1\x9f\xf0\x0f\xf8\x05\xac\x86\x59\x70\x77\x0a\x21\xef\xc1\x3b\xd0\x0f\xbd\xf0\x17\x58\x09\x8f\x4a\x90\x17\x0c\xe4\xff\x19\x43\x11\x7a\x97\x42\x57\x37\x50\x57\x93\xaf\xd5\x17\xf0\xd1\xe6\x8e\xf6\x36\x9a\x4c\xd2\x6a\x21\x20\x76\x2a\x6c\x93\x18\x61\xb4\x2e\x95\xa2\x5b\x37\xf8\x3a\x7c\x54\x96\x68\x03\xad\x5e\xe6\x22\xa1\x78\x34\x48\xdd\x11\x31\x16\x63\x1a\xad\x49\x26\xab\x85\x36\xeb\xb3\x99\x3e\x95\xf2\xd0\x26\xa6\x30\x83\xb9\x55\x7b\x9f\xdf\xd0\xe2\x41\x23\xd0\x13\x63\xe9\xfd\x4c\xd3\x54\x8d\x26\x29\xc1\xa9\xe4\x10\x65\x5d\xb4\x3a\x6a\x9e\xca\x25\x4b\x2e\x9a\x4a\x61\x73\xb7\x6c\xec\x76\x36\xa2\x78\xe6\xb2\x5d\x0b\x88\x61\x5a\x1d\x92\x99\x22\x59\x49\x34\x31\x1a\x66\xce\x06\x33\x91\xd0\x6c\x7e\xd2\xcd\xf2\x92\x94\x55\x39\x8d\x19\x71\x2d\x6a\x9e\xcd\x5a\xdd\x55\x6b\x7e\x34\x93\x1a\x09\xc1\x2c\x8a\xe4\x4b\xb0\xa0\x1b\x4d\x50\x4b\x55\xcf\xa0\xe6\x48\x26\xe3\x31\x49\x34\x98\x7f\x73\x6b\x26\xdf\x6c\x93\xa4\x52\x2e\x32\xa4\x0d\xf2\x9a\xc0\x3e\x3e\x90\xb0\x1b\x61\x48\x1b\xd4\xa2\x0c\xb4\xc6\x9a\xb6\x12\x48\x64\x1a\x84\x14\x2e\xb8\xbd\x11\xa1\x0a\x9b\x18\x93\xf4\x96\xa8\xce\x34\x03\xd7\x22\x24\x2a\x3a\xb3\x76\x9a\xad\x88\xba\x35\xa0\xad\x64\x05\x19\x8d\xba\xbe\xb2\x95\xe7\xf0\xf5\x35\x53\x17\xab\xb3\x9d\xc7\xd8\xd4\x39\x5b\x78\xfb\x74\x42\x53\xe3\x04\x5d\xb3\x61\x4a\x3f\x1e\x05\x1e\xbb\x46\x2f\x5e\x68\x8c\x18\x16\x95\xf2\xc7\x46\xa2\xd6\xec\x30\x83\x7b\x19\x4d\xa0\x6b\x99\xe7\x88\x4a\x6e\xec\x97\x3d\x99\x1d\xe8\x65\x7a\x50\x8c\xfa\x50\xcf\x85\x4e\x4f\xcb\x66\x94\xde\x93\xcd\x27\xbd\x05\x95\xd5\xd4\x6e\xdd\x29\x9d\x1f\x5b\xdd\xb9\x99\x5a\x23\x9c\x75\xbb\x0b\xde\x2b\xac\x64\x21\x54\xc1\xca\x61\x13\x4b\x18\x6e\x8f\xd3\xbf\x25\x16\x62\xce\xf6\x75\x8a\xaa\x33\xb7\x27\x77\x28\x0c\x2d\xa0\xb5\xd9\xa9\x3d\x0e\xca\x1d\x1d\x9b\xe3\x4c\xeb\x49\xd8\x57\x29\xe6\xc9\x6b\x0e\xf3\x7a\x2d\xc5\x60\x8b\x18\x82\x3f\xa6\xc9\x51\x23\xe4\x76\xb5\x37\x37\xfb\x7d\x01\xba\x48\x72\xe1\xb2\x86\x42\x3a\x33\x32\x27\xb7\x57\xe9\x1a\x5a\x6f\x36\x7b\xf1\xac\x5a\x5b\xda\x5a\xcc\x9c\x28\xb2\x52\xe4\x88\x9c\xcd\xc9\x5a\x1b\x79\x46\xed\x1d\x4d\xbe\x0e\xda\xb8\x9d\x2e\xd2\xcd\xbc\x54\x4d\x62\x5a\x63\x4f\xb6\x5c\xf6\xba\xd9\x30\x2e\x97\xd3\x84\xe9\xae\xe8\xc7\xed\x7f\x5d\x80\xd6\xa4\xbf\x00\xaa\x73\xfb\xa1\x73\xf7\x4f\x26\x87\xde\x86\xd1\x51\x7e\x45\xed\x7e\x9a\xb0\xbb\xfe\x5e\xe6\xde\xb1\xb3\xe0\x6d\xbf\x96\xd6\x7b\x9c\xce\xac\x1b\xaa\xc6\xec\x2f\x87\x66\x39\x2a\x65\x3b\xb3\xb9\xe6\x57\xe4\x20\x6b\xec\xc1\x4e\xeb\x24\xd6\x20\xa2\xfb\xa9\x79\xe4\x3a\x51\x47\x32\xdc\x60\x62\x34\xb3\x7b\x98\xfb\x4b\xa6\xfe\x18\x20\x06\xd3\x42\x62\x90\x25\x53\x4e\x03\x9b\x9b\xd2\x97\xcd\xfc\xec\xa1\xc3\x97\xd9\x1a\x3f\x9e\x21\xa3\x7e\x82\xc7\x8d\x7d\x72\x3e\x74\x8b\x0c\x5d\x51\x0b\xeb\x9e\xb1\xed\x8d\x6b\x15\x25\x60\xe5\x4e\x75\x03\xe3\x2c\x6c\x9f\x84\xee\xd8\x39\xa8\x5b\x0d\xfa\x62\x1b\x75\x67\x5a\x6c\x57\xd0\xda\x94\x57\xb1\x0e\xb5\x3b\xb7\x6e\xb9\xd7\x61\x31\x4e\x65\x1d\xb2\x4f\xd4\x50\xdf\x42\x67\x1c\x49\x05\xdb\xa3\x6c\xd4\x15\x9c\xd4\x63\x65\x52\x76\x54\xeb\x12\x98\x3f\x09\xc6\xbf\xbb\x16\xb9\x96\x63\xd3\x3f\x07\xd7\x67\x42\x7a\x69\xf1\x7a\x15\xad\x95\x81\xad\x34\xef\x17\x94\x99\x25\x0e\xc0\x45\x6c\x6a\xa4\x11\x55\x62\x8a\xe0\x73\x26\x99\xfb\xd1\xe1\x71\x31\xd3\x47\xe5\x96\xaa\xf8\x51\x99\xdf\x5b\x72\x24\xa6\xe2\xd7\x3a\x7e\xee\x55\xdb\x1f\x85\x16\x6b\xc1\xa4\xf4\xcf\xbc\xcc\xb3\x8c\x93\xd4\x7c\x98\xb1\xd3\xe8\x38\x3e\x93\x01\xbe\xec\x5d\x99\x95\xc7\xcf\x8e\xb7\x4a\x54\x30\xf1\x9e\x88\x2a\x6b\x6a\xd4\x9a\x31\xef\xb2\x9e\xff\xe3\x78\x56\x7f\xa3\xca\x7e\xfe\x4f\xbf\xbf\xbf\x86\x0f\x17\xa1\x1f\x2e\x80\x08\x4b\x60\x31\x7c\x89\xb4\x77\xe1\x0e\x7c\x04\xb7\xe1\x7d\xb8\xe1\xe4\x73\xbd\x6a\xa2\xfe\xbd\xe0\xc1\x83\x07\x0f\x1e\x3c\x78\xf0\xe0\xc1\x83\x07\x0f\x1e\xcf\x64\xa4\x1f\xb8\xf3\x97\x45\x5f\xd8\x18\x09\x21\xe7\x25\x8d\x30\xe8\x9f\x67\xf3\xef\x21\x2b\xf9\x78\xbe\x29\x4d\xbf\x32\x2b\xf4\xef\x45\xa1\x77\x66\x74\x14\xef\xa0\x1e\x5b\x33\xf3\xf9\x7f\x63\x09\x21\x55\xe6\x04\xfe\x39\xd9\xf9\xfd\x1f\xe2\x39\x5f\x81\x36\x18\xc0\xf3\xfc\x15\x38\x07\x2f\xc0\x59\x70\x43\x05\x9c\xaf\x24\xa4\x0f\xfe\x0a\x2a\x44\x40\x81\x4b\xb3\x08\xe9\x85\x56\x58\x01\x25\x30\x30\x93\x90\xaf\xe0\x0e\xbc\x0c\xe7\xa0\x0f\xfe\x0b\x11\x08\xc3\x1a\x98\x07\xa5\xf0\xc9\x0c\x42\xae\xc2\xff\xa1\x0b\xb6\x41\x07\xfc\x1a\xd6\x40\x03\x1c\x98\x4e\xc8\x76\xa8\x82\x2f\x2a\x08\xb9\x0c\xaf\xc0\x45\x38\x0d\x7f\x04\x03\x74\xd8\x05\x5e\xf0\x40\x19\xdc\x2a\x27\xe4\x1a\xbc\x04\x47\x9c\xf7\xff\x3f\x07\x0a\xf3\xa0\x02\xca\xe0\xeb\x32\x42\x1e\xc1\x43\x38\x0f\x7d\x70\x00\xb6\xc3\x46\xd8\x00\xeb\xa1\x04\x08\xbc\x35\x8d\x90\x53\xf0\x1f\xe8\x85\xa3\xb0\x16\x7e\x09\x55\xe6\xfc\x80\xa9\x84\xbc\x0d\xfd\xf0\x22\x9c\x84\xe3\x70\x0c\xda\xa0\x15\x16\xc0\x00\xae\xc9\x7d\xb8\x07\xd7\xe1\x34\xfc\x0b\x5a\x60\x76\xa9\x73\xcd\x78\x8c\x65\x8c\xf7\x38\xe7\x33\x22\xf8\x8c\x08\x3e\x23\x82\xcf\x88\x98\xec\x6f\xab\xf8\x8c\x08\x3e\x23\x62\x92\xcf\x88\x98\x94\xe3\xa6\xe0\x5b\x5e\x3e\x7a\x26\xf9\xe8\x19\x6e\xec\x8c\xeb\x70\x19\xd1\x9c\x8f\x58\xfe\xab\x74\x2b\xf9\xd8\xb5\xb4\xdd\x55\x96\xd9\x0d\x3e\xfc\x0b\x7a\x3a\x92\x66\x7e\xe2\xfb\xd3\xa8\xe7\x95\x8c\x62\x5a\xc2\x90\x86\x9c\xf8\x46\x7c\xd2\x36\x2c\xdc\x82\x23\x68\xb4\x61\x9f\x0b\x50\xa6\x6f\x03\x00\x00\xff\xff\xe4\xd8\xf6\x13\x00\x50\x00\x00")

func templatesDaoSqlxTmplSwpBytes() ([]byte, error) {
	return bindataRead(
		_templatesDaoSqlxTmplSwp,
		"templates/dao/.sqlx.tmpl.swp",
	)
}

func templatesDaoSqlxTmplSwp() (*asset, error) {
	bytes, err := templatesDaoSqlxTmplSwpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dao/.sqlx.tmpl.swp", size: 20480, mode: os.FileMode(420), modTime: time.Unix(1459819791, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDaoPgxTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\xdb\x6e\xe3\x36\x13\xbe\xf7\x53\xcc\x2f\x64\x7f\x48\xa9\xaa\xec\xde\x16\x70\x8b\x75\x6c\xa3\x01\x92\x75\xd6\x76\x51\x14\x45\x11\x28\x16\xe5\xa8\xab\x53\x48\x3a\x71\xa0\xea\xdd\x3b\x24\x45\x45\x72\xe4\x48\xb2\xf7\x74\xd1\x20\x90\x25\x1e\xe6\xf0\xcd\x37\x1c\x92\x41\x94\x26\x94\x83\xe1\x47\xdc\x18\xe8\x8f\x75\xc0\xef\x36\xb7\xce\x2a\x89\xce\xfe\x76\x57\x9f\x56\x67\xe9\x7a\x6b\x0c\x06\x59\x46\xdd\x78\x4d\xe0\xa4\x18\xf7\xd3\x10\x9c\x0b\xf9\xca\xf2\xbc\x9c\x9c\x65\x45\x7f\xd1\x47\xbc\x0f\x6e\x44\xf2\xdc\xc0\xf9\x24\xf6\x70\xe4\x80\x3f\xa5\x04\xb2\xcc\x99\x06\xb1\x47\xa8\xea\x06\xc6\xe9\x66\xc5\x21\x83\xf1\x08\x4e\x51\xa1\x73\x9e\xc4\xf1\x75\x92\x84\x90\x97\x13\xae\xdc\x34\xed\x34\x61\x30\x38\x3b\x5b\xac\xdc\x98\x01\xbf\x23\x40\x93\x47\x08\x62\x9e\x80\x0b\x31\x79\x84\x53\x34\xd1\x59\xc8\xd9\x4b\x14\xac\xc4\x0d\xfc\x4d\xbc\x02\x86\x73\xb2\x0c\x76\xbb\x21\xcf\xe7\x28\xc3\x14\x82\xa4\x2a\xfc\xb2\xc0\x6c\x14\x64\x03\xa1\x34\xa1\x16\x64\x15\xc0\xd0\xcf\xad\x0d\x27\x7e\x40\x42\x4f\xe2\x36\x15\x6f\x0c\x7e\x44\xc5\x80\x7f\x0f\x2e\x85\x07\x81\x9c\x18\x88\xbe\xe1\xab\x1c\xeb\x54\x0c\xd4\xf0\x89\xf1\xa8\x42\x88\x41\x7b\x1c\xe1\xa6\xd9\x43\x95\xf8\xfb\x7f\x45\x97\x8d\x73\x01\x25\x97\x03\x2c\xf9\x0c\x7c\xa9\xe4\x7f\x43\x88\x83\x10\x7d\xd1\x53\x29\xe1\x1b\x1a\x8b\x46\xe9\xa8\x6c\x57\xf3\x28\x61\x42\x1f\x22\x6c\x36\xe1\x62\xf5\xb0\x11\x45\x39\x25\x04\x45\xb8\x87\xcf\x00\x81\x44\xa3\x6a\x74\x61\x15\xce\xb3\x85\x69\x83\xfc\x05\x01\x58\x95\x01\x7f\xfe\x75\x28\x07\x98\x29\x65\x69\x12\x30\x64\xc1\x1e\x61\x15\x1e\x14\x2e\x6d\x42\x2e\x11\x8a\xdc\x4f\x64\xff\xac\xb7\x36\x46\x3f\x24\xb1\x06\x45\x20\x27\x04\x78\xc4\x27\x54\x7a\xe2\x9c\x87\x09\x23\xa6\x6a\xf6\x93\xa2\xf1\x03\xd9\x72\x53\x6b\xeb\x47\x87\x2e\xec\x53\x52\x6b\x44\xa9\x13\x91\x15\x4c\xec\xaf\xfe\x25\x23\xf7\x28\xb3\xca\xb7\x7d\xec\xdc\xc7\xd0\x67\x96\x76\x62\x6a\x7f\x1f\x3a\x30\x76\x8f\x53\x9a\x18\x43\x10\x8b\x5b\xec\x99\x45\x83\x2d\x7a\x2c\xa8\x25\x98\x26\xb9\xea\x2f\x89\x3e\x4d\x68\xe4\x72\xe4\x76\x23\xa5\x34\xf1\x59\x18\xac\x08\xa6\x03\x8e\xc4\x40\xc1\x2d\x81\x0d\x23\x1e\xf6\xe2\x3f\x23\xb8\x74\xbb\x68\xd9\x26\xf5\x5c\x4e\x70\x6d\xc5\x67\x44\x62\xce\x54\x4e\x90\x2d\x59\xbd\xa7\x6b\xd6\x9c\x17\xe6\x43\xb3\x66\x0b\x13\x0d\x95\x13\xea\xbb\x2b\x92\xe5\xd9\xe0\xc1\x0d\x25\xf2\xb5\xe6\xc3\x28\xbb\x8b\x76\x23\x69\xc4\xb3\x40\x4d\xa8\x16\x68\x49\x77\x4c\xc6\x13\x4a\x94\xd1\xd5\x22\x64\x81\xf8\x9a\xc5\xc4\x64\xf7\xa1\x28\x30\x41\xbc\xb6\xc1\x45\xc7\xeb\x26\x77\x58\xfb\x2b\x01\x7b\x75\x41\x51\xa6\x38\xe3\x91\xf3\x71\x43\xe8\x93\x6c\xb9\x0f\x95\x52\xc7\x71\x2c\xab\xa3\xd1\xef\xc3\xb0\xdd\xe8\x8e\x4b\x15\x26\xb3\xad\x33\xbb\x6e\x5e\xdd\xb6\xa6\x42\xb1\x9b\x7e\x75\xee\x76\x58\x5b\xbb\x3a\x8c\x51\x5a\x6e\x5b\x5c\xb6\x81\x6f\xd5\x5a\xbd\xdc\x7e\x9f\x31\xeb\xeb\xc2\x01\x11\xe4\xdb\xaf\x1b\xbb\x32\x91\x77\x13\x58\x17\xd8\xd7\x41\x19\x3d\xd5\x32\x1b\xfe\x01\x31\xe9\xdc\x65\x38\xc2\x4c\x9b\x6a\x53\xc7\xc8\x0a\x9c\xd1\x18\x63\x31\xb9\x9c\x9c\x2f\xe1\x14\xa6\xf3\xd9\x95\x90\xe7\x2c\xdd\xdb\x50\x2f\x96\xbf\xff\x3a\x99\x4f\xb0\x35\xc6\xcf\x19\x5d\xba\xeb\xc2\x0f\xb9\x9e\xff\x02\x97\x17\x57\x17\x4b\x78\x67\x1c\xcd\x95\xb4\x33\x4b\x5e\x03\x04\xe9\xd3\x08\xc9\x01\xc4\xff\xfc\xf0\x74\xc7\x48\x53\xf4\x00\x74\x16\xa2\xaa\xf5\xe7\x8c\x0d\x09\x45\x49\xa3\x27\xa8\xe5\x5a\x18\x44\x01\x17\x4d\xd8\xef\xfb\x8c\xc8\xf7\xce\x69\xf7\x59\x20\x34\x74\x66\x6a\x03\x31\x3b\x0d\xa3\xb2\xc7\x11\x5a\x7e\x18\x02\x9e\xd8\x9c\x45\x8a\xcb\x06\xf7\x4d\x63\x36\x1f\x4f\xe6\x30\xfa\x03\xde\x30\x30\x4a\xdf\xac\xa2\x08\x6a\x89\xca\xbb\x9f\xe1\x6d\x8b\x34\x15\xc5\x37\x9e\x10\x25\xe7\xec\x0a\x2a\xb0\x69\x97\x34\x9b\x4e\x17\x13\x21\xca\xd0\x80\x5a\xd5\x0d\x4d\x5b\xa1\x49\xbf\x6d\x89\x69\xe5\x56\x97\xf4\xfb\x8f\x69\xdf\x1b\xd3\xea\x05\xb1\x17\xc7\xf6\x1c\xc2\xfa\x51\xaf\xbc\x07\x29\x38\x18\xc9\x5b\x0d\x45\xc2\xea\x0d\x87\x05\x17\x72\x63\x6e\x26\xa2\xaf\x61\x77\x2d\x29\x51\xe0\x52\xa3\x45\x96\xa9\x2d\xfd\xe2\xe3\xe5\xf5\x7a\x5b\x6e\xa0\xab\xe4\x50\xc1\x67\x3c\x42\x02\xa6\xe2\x67\xa2\xc0\x51\xc6\x88\x3c\xbc\xa6\x24\x75\x29\x31\x8d\x66\x97\x6e\x94\x0a\xc3\x46\xa5\x25\x84\xa5\xa4\x7d\xf7\x06\x7a\x40\x35\xcc\x37\x65\x64\x9e\x95\x4f\xf0\xc0\x61\x8a\xb1\x0e\xfa\x60\xb7\x9d\x3f\x12\xab\xdc\xd2\x14\x7a\xc4\xc1\x2f\x6f\x07\xf8\x37\x79\xdc\x69\x05\x78\x07\x5c\x75\x48\xfa\xa2\xe0\x2a\x15\xc7\x81\xfb\x8d\xb1\x55\xe4\xc5\x25\xb2\x19\xdd\xfa\x16\xe5\x78\x2a\xdf\x54\xf3\x5b\xb9\x88\xe9\xfd\x65\x89\xd3\xd3\xb9\xc3\x68\xd4\xec\xd8\xa1\x61\xdb\xbb\x3b\x97\x87\xe6\x2c\x7b\x0c\xf8\x1d\x9c\x88\x8a\x21\xaf\x48\x5e\x54\x0e\x71\xf1\x26\xd6\xca\xfb\x62\x90\x11\x78\x86\xb8\xdd\x68\x45\x6c\x4c\x42\xa2\x53\xad\xe9\xaa\xa0\x00\xa9\x1e\xfd\x31\xd6\xb7\xe5\x44\x57\x37\xa8\x94\x37\x61\x6e\x51\xe0\x02\x0f\xcb\xd9\xc9\xbb\xe3\x93\xce\x93\x26\x7e\x9d\x15\x2d\xd9\xbd\xc0\x38\x84\x85\x0a\x53\xcd\xc2\xa6\x5d\x43\x03\x0d\x8f\x41\xf8\x35\x42\x76\xf1\x48\x97\xbe\xdd\xdf\x7f\x03\x00\x00\xff\xff\xff\xe2\xa9\x7f\x87\x18\x00\x00")

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

	info := bindataFileInfo{name: "templates/dao/pgx.tmpl", size: 6279, mode: os.FileMode(420), modTime: time.Unix(1459819271, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDaoSqlxTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x57\x6d\x6f\xe2\x46\x10\xfe\xce\xaf\x98\x5a\xb9\x93\x49\xa9\x73\xfd\x5a\x89\x56\x25\x18\x35\x52\x72\xe8\x00\xa9\xaa\x4e\xa7\xca\xc1\x63\xce\xad\xed\xf5\xad\x97\xc6\x91\xcb\x7f\xef\xec\x9b\x63\x83\x09\xd0\x12\x55\x27\x25\x5f\x62\x76\x5e\x76\x5e\x9e\x79\x76\x37\x4e\x73\xc6\x05\x38\xab\x58\x7c\x5e\xdf\x7b\x4b\x96\x5e\xfd\x91\xb2\x98\xb3\xec\xaa\xf8\x92\x94\x4e\xaf\x43\x61\x7d\x8f\xfc\x9e\xb3\x50\x5c\x61\x89\x69\x9e\x04\xfc\x2a\x65\x21\x26\x4f\xca\x51\x2a\x9c\x5e\xaf\xaa\x78\x90\xad\x10\x2e\xcc\xf2\x0f\x43\xf0\x6e\xd4\x67\xb1\xd9\xd4\xba\x55\x65\xe4\x46\x86\xe1\xfb\x20\xc5\xcd\xc6\x21\x7b\xcc\x42\xd2\xec\x89\xc7\x1c\xa1\xaa\xbc\x49\x9c\x85\xc8\xb5\x18\x0a\xc1\xd7\x4b\x01\x15\x8c\x47\xa0\xb6\xf7\x7c\x13\xcd\x9c\x02\x87\x4d\x6d\x75\x17\xe4\xf9\xf1\x56\xbd\x68\x9d\x2d\xc1\x2d\x04\xe3\x08\x97\x14\x5c\x6b\xd7\x3e\xc8\x5f\xd3\x0c\x5d\xaa\x8e\x74\x16\x67\xab\x01\x04\x7c\x55\xc0\xc7\x4f\x71\x26\x90\x47\xc1\x12\x2b\xd2\x73\x95\xed\x5c\x6d\xb7\xa0\x48\xb4\xfd\x00\x90\x73\xc6\xfb\x50\xf5\x80\xfe\xfe\x0a\x38\x94\xd0\xa5\xa8\xc4\xa4\x2b\x8b\xa6\x62\xf1\xc6\x23\xef\xc3\x1a\xf9\xe3\x8c\x3d\x94\x72\x77\xbd\x6d\xdf\x58\xce\x97\x41\xe6\xbe\x2d\xfb\xa0\x0c\x39\x8a\x35\xcf\xe0\x6d\xa9\xf6\xeb\x1d\x97\xd5\xcf\x49\x72\x38\xab\x8f\x9f\x8e\xc9\xab\x94\x61\xa7\xc1\x9f\xb8\x5f\xff\x5d\x5f\x47\xca\x1e\x8a\x41\x77\xa2\xcd\x2c\x95\x6e\x1c\x29\xc5\x6f\x86\x90\xc5\x09\xb5\xd0\x64\xa9\x93\x04\x5d\xb2\x10\x23\xe4\xca\xab\x77\x9d\xb0\x02\x5d\x6d\x1a\x31\xb3\xf8\x1e\x4b\xe1\xda\x30\x95\x53\xb9\x71\x86\x0f\x6e\x57\x98\xfd\x5a\xad\xa0\xfa\xfa\x3a\x4a\xe5\xa7\x51\xf5\xf8\x49\x8b\x22\xb4\x8a\xbb\x51\x5a\xc9\xa6\x56\x2f\x61\x08\x12\x9c\x59\xe8\x52\x12\xc6\xcf\xa6\xd9\x41\x5a\x26\x2f\x47\x36\x90\x60\xb9\x28\x0f\xb4\x70\x00\xa2\x84\x4b\x39\xda\xde\xa2\x3c\x2b\x4a\x45\xf9\xf2\xf8\x3c\x39\xbd\x17\x85\xab\xcd\xf8\x15\xa8\x4d\xa0\xd6\xac\x1f\xc5\x98\x84\x8a\xf4\x27\xf2\x4b\x72\xfe\x11\x5d\x1e\x3d\xd2\xba\x32\xf5\xe4\x32\xfc\x0d\xd2\xe8\x3a\x28\x48\xc3\xcd\xa1\x16\x36\x72\x3f\x0b\x8a\x25\xae\x28\x56\x67\xee\xdf\xfa\xd7\x0b\xb8\x84\xc9\x6c\x7a\xa7\xd4\x17\xc1\x7d\x62\x34\xe1\xd7\x5f\xfc\x99\x4f\xab\x19\xfd\x9c\xf2\x45\xb0\x32\x69\x92\x68\x08\x3f\xc1\xed\xcd\xdd\xcd\x02\xbe\x77\x8e\x61\xef\xfc\x8c\xa3\xf1\x5c\xd1\x68\x66\x3a\xcb\x76\x5e\x26\x80\x17\x2a\xe2\xf3\x0c\xb3\x53\xc3\x7f\x5f\xc2\x79\x12\x2f\xf1\x74\xf0\x0d\x80\x71\xf2\x34\x7a\x84\x16\x0b\x25\x71\x1a\x0b\xb9\x44\xf2\x28\x2a\x50\x7d\xbf\x00\x1f\x9d\xa5\xe0\x8e\x25\x2d\x9b\x0a\x51\x82\xe3\x34\xa8\x47\xee\xf2\xed\x10\xe8\x56\xe7\xcd\x73\xa2\x5e\x11\xb9\xce\x74\x36\xf6\x67\x30\xfa\x0d\xde\x14\xe0\xd4\x55\xb0\xcc\x60\x3d\xea\x3a\xfc\x08\xef\x0e\x78\xd3\x3d\x7f\x13\x4a\x57\xca\x66\xdb\x91\xa9\xe2\x61\x4f\xd3\xc9\x64\xee\x4b\x57\x8e\x2d\x7d\x8b\xad\x0e\x5e\x37\xf2\x57\x0a\x3f\xcb\xb8\x1c\x45\x3b\xaf\xd3\xf3\xf5\x4e\xcf\xd6\xed\xe7\x75\x6e\x9e\xde\xab\x66\x80\x52\xf5\xf0\x94\x13\xd4\x7a\x83\xf6\x09\xcd\x05\x72\xe1\x32\x25\xda\xc9\x47\xa3\xd9\xa4\xdf\xc2\x69\x55\x69\xcb\xf9\x87\x5b\x7b\xa9\x82\x26\x54\x35\x14\x7f\xaf\x5b\xa4\x03\x90\x0c\x27\xe5\x21\xbd\x74\x97\xba\x59\xac\x75\x50\xb6\x0e\xc9\xbd\x41\xaf\xf3\x30\x10\x78\x30\xe8\xad\x80\xb5\xd5\xff\x12\xb0\xae\x15\xf1\x50\x77\xc8\x5b\xf7\x9f\xb3\x15\x9d\xe6\xa2\x3b\x78\x33\x1a\x43\x03\x3a\x8b\x30\x46\x78\xc7\xb0\xb8\x51\x9b\x10\xd8\xa2\x20\x29\x70\x17\x65\xa7\x34\xe9\xd4\x9c\xff\x53\xcf\xf6\xe6\xdb\x0a\x7c\xef\xab\x00\xbe\xdb\x48\xe9\x43\x2c\x3e\xc3\x85\xa4\x58\x35\xeb\x3b\x54\x4b\x2a\xaa\x80\x5f\x8c\x92\x13\x87\x0e\xd4\x2f\x8a\xa7\x7a\x5c\x6c\x15\x64\x8c\x09\x5a\xd4\x76\x70\x87\x2d\x41\xbb\xe5\x63\x3a\x0f\x16\xbe\x3d\x0d\xa0\x71\x1c\xc8\x68\xcd\x81\x10\x87\xd4\xac\x0b\xba\xa6\xee\xc5\x70\xa3\x22\x5e\xeb\xa0\xb4\xa4\xf5\x7c\x6b\xbb\x53\xb1\xbd\xed\x3a\xd6\xba\x9a\x7b\xc6\xcc\xa8\xd3\x27\xa6\x54\x53\xe2\xf6\xff\x7f\x02\x00\x00\xff\xff\xfd\x78\x2b\x1f\x77\x14\x00\x00")

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

	info := bindataFileInfo{name: "templates/dao/sqlx.tmpl", size: 5239, mode: os.FileMode(420), modTime: time.Unix(1459819773, 0)}
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
	"templates/dao/.sqlx.tmpl.swp": templatesDaoSqlxTmplSwp,
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
			".sqlx.tmpl.swp": &bintree{templatesDaoSqlxTmplSwp, map[string]*bintree{}},
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

