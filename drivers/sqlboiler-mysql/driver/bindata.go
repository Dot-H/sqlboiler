// Code generated by go-bindata.
// sources:
// override/templates/17_upsert.tpl
// override/templates_test/singleton/mysql_main_test.tpl
// override/templates_test/singleton/mysql_suites_test.tpl
// override/templates_test/singleton/mysql_upsert.tpl
// override/templates_test/upsert.tpl
// DO NOT EDIT!

package driver

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

var _templates17_upsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x5f\x6f\xdb\xc8\x11\x7f\x16\x3f\xc5\x9c\x70\xbd\x23\x0b\x85\x49\x81\xa2\x0f\x2e\xf4\x70\xfe\x93\x9c\x1b\x3b\x27\x59\x49\x03\x34\x08\x82\x35\x39\x94\x17\x5e\xed\xf2\x96\x4b\x27\x2a\x8f\xdf\xbd\x98\x59\x52\x22\x25\x59\x52\x92\xba\xe8\x93\xcd\x9d\xdd\x99\xdf\xfe\xe6\xef\xaa\xaa\x9e\xc1\x8f\x4e\xdc\x2a\x7c\x23\x16\x38\x93\x7a\x5e\x2a\x61\xe1\x64\x0c\xf1\x5b\x5a\x8d\x69\x19\xfe\x80\xa2\x95\xfc\x01\x4e\x3a\x85\x67\xa2\x40\x78\x56\xd7\x01\x2b\x78\x10\xf6\xe8\xe3\x89\x58\xa0\xea\x1f\x2f\x92\x3b\x5c\x08\x3e\xb0\x7d\x34\x9e\xad\xa5\x7c\x40\x66\x10\xff\x92\xa6\xaf\x94\xb9\x15\x8a\x95\x3c\x7f\x0e\xef\xf2\x02\xad\x7b\x05\xc2\x39\x5c\xe4\xae\x00\xa1\x41\x6a\x5a\x1b\x81\xd0\x29\xa4\x06\x79\xad\xcc\x53\xe1\x10\x8c\x05\x39\xd7\xc6\x22\x18\x0d\x89\xd1\x99\x92\x89\x8b\x83\xac\xd4\x09\x84\x06\xfe\x5c\x55\xdb\xa4\xd4\x75\xd4\x9a\x09\x19\x85\x36\x0e\xe2\x37\xe6\xcc\x68\x87\x5f\x5c\x5d\x27\xee\x0b\xe9\xa2\x8f\xb8\x59\x1c\x41\x55\xa1\x4e\x09\x65\x63\xfa\xcc\xa8\x72\xa1\x8b\x51\x83\xae\xf9\x84\x5b\x23\x55\xdc\x7c\x44\x80\xd6\x1a\x0b\x55\x30\xb0\xe8\x4a\xab\xc1\xc4\xde\xb0\xb7\xdb\xb5\xc9\xe7\x5e\xa1\x3b\x3f\x0d\xa3\xaa\x42\x55\x20\xe3\x18\x41\x2b\x68\x76\x36\x72\x9d\xd6\xf5\x68\x2f\x92\x28\xa8\x83\x60\x05\x3a\xf0\x7c\x13\x83\x1d\xce\xe9\xdf\x89\xd0\x32\xd9\x60\x7f\xf2\x7d\xf4\x03\xeb\x2c\x68\x8d\x09\x38\xde\x1f\x93\xa7\x76\x48\x15\x0c\x64\x46\xa8\x28\x3e\xff\x97\xde\xf8\x3b\x1b\xfd\x61\x0c\x5a\x2a\x42\x31\xc8\x89\xa3\x90\xf5\xbd\xb7\x22\xbf\xb0\x36\x44\x6b\xa3\x28\x18\xd4\xbb\x3c\xf7\x88\xab\x76\x79\x0a\x4a\xca\x52\xfa\xc6\x2f\x98\x94\xce\xd8\xaf\x49\x9d\x8e\xea\xfc\x1b\xdd\x38\xd9\x26\x94\x90\x78\xf2\x2e\x1a\x4c\x1d\x5a\xb7\x7d\xbb\xde\xde\x2c\x75\x4e\x1d\x26\xfb\x78\x9f\xef\x08\xb4\x6e\x60\x11\x8c\xa7\xf3\xeb\x8a\xe9\xa7\xf0\xe1\x0c\xb1\x47\x03\xa4\x26\x29\x17\xa8\x9d\x70\xd2\x68\xc8\x8c\x85\x3b\xf3\x19\x9c\x81\xdc\x9a\x1c\xad\x5a\x42\x59\x60\xff\xae\x6c\xb1\x77\xdd\xa3\x43\xe0\xff\x2b\x02\x56\x65\x58\x66\x60\x60\xbc\x76\x56\x53\x96\x59\x5e\xc4\x6f\xf0\x73\x38\xac\xaa\x78\x72\x3f\xa7\x5b\xd5\xf5\x09\x68\x03\x55\xd5\xe9\x63\x75\x4d\x7c\x3d\xc8\x14\x53\xe6\xb0\xe4\xdb\x0e\xd9\xbb\xc1\x80\x9a\x20\x79\x52\x91\x6f\x86\x4e\x2e\xb0\x70\x62\x91\x7f\xf2\xbb\x3e\xdd\xa1\xca\xd1\x0e\x21\x86\xda\xef\x5e\x47\xe0\xaf\xc6\xdc\x17\x1c\x16\xbd\x58\x4d\xcd\x29\x66\xc6\xa2\x27\x95\x37\x1d\x1d\xb8\xdb\xa1\xb9\xbe\x2d\xc1\x65\xb4\xcc\x65\x10\x0c\xf4\xbf\xcf\x31\x13\xa5\x72\x05\x19\xfe\xbd\x44\x2b\xb1\x88\xdf\x18\xfd\x2f\xb4\xa6\x11\xcd\x90\xdc\xba\x39\x23\xd4\x75\x43\xf3\x7b\xe9\xee\x9a\x9d\x23\x30\x51\x10\x0c\x9e\x3f\x87\xd3\x52\xaa\x14\x12\x91\xdc\x21\xdc\xe3\x12\xa4\x7e\xa6\xa4\x46\x28\xe7\x4a\xaa\x25\x3c\x83\xc5\xb2\xf8\x5d\xc1\x43\x01\x39\xfd\xcd\xad\xb9\x55\xb8\x28\x82\xc1\x6d\x99\x11\x92\xc2\xd9\x85\xd0\x73\x85\x54\x74\x4f\xcb\x2c\x43\x1b\x46\x2c\x8d\xdf\x5b\xe9\x70\xe6\xac\xd4\xf3\xb0\x70\x36\x31\xfa\x21\xbe\x74\x46\x84\xbd\xc0\x88\x5f\x4b\x9d\x52\xf6\x91\xb7\x3e\x8d\x20\x21\xad\x56\xe8\xf9\x46\xac\x53\xb0\x14\xcc\xd2\xa6\xee\x84\x9d\xbb\x5e\x3e\x5d\x3a\x0c\x7f\x8e\x7f\x3e\x04\xa3\x9f\x38\x8f\xc3\xe8\xef\xfb\x16\x18\xdb\x3a\x3b\xee\xdc\xa3\x8b\x1c\x72\x32\x06\x92\x36\x82\x28\x18\xac\x19\x9f\x94\x2d\xe3\xb7\x65\x16\x71\xc0\x6e\x3b\xdf\x47\xe6\x19\x39\xf8\xba\x74\xf1\xcd\x95\x49\xee\x49\x0d\xbb\x7c\xe4\x3d\x9f\x92\x95\x03\x87\x3f\xdc\xe3\xf2\xe3\x71\x26\xde\x69\xe5\x8d\x04\x83\x07\x61\x39\xc8\x39\x81\x03\xce\x9d\x1f\x1a\x93\x74\xef\x76\x72\xb1\xe8\x08\x42\x9f\xe9\xcb\xce\x17\x85\x76\x30\x18\xec\x34\xdf\x56\x98\x03\xf2\x6e\xf8\x1f\xb1\xd5\x94\xae\xbb\x7b\xed\x30\xfa\x8c\x82\xc1\xa0\xa9\xf0\x27\xe3\x8d\x38\x7d\xd7\xf9\xfa\x1e\xd8\x13\x2b\x17\xc2\x2e\x5f\xe3\xb2\xb3\x93\x38\x25\x12\x15\xea\x26\x8d\x22\x2a\x97\x2f\x98\xcd\xc3\xd5\xb2\xd4\x3c\xfa\x3b\xd3\xd4\xc5\xcd\xda\x39\x82\xc4\x94\x2a\xe5\xea\x75\xcb\x95\xa1\xb9\x65\xc2\x10\x40\xc9\x82\x6b\x29\x17\x53\x1f\x42\x31\xd5\xa2\x25\x8c\xfd\x7e\x1f\x0a\x53\x5a\xba\x5e\xce\xa6\x57\x61\x2a\x85\xc2\xc4\x8d\x60\xb8\x61\x6a\xd8\x36\x88\xb6\x33\x44\x2b\x8d\x16\xbd\x06\x18\x43\xb6\x70\xf1\x2c\xb7\x52\xbb\x8c\x89\x1c\xce\x2e\xae\x2e\xce\xde\xc2\x9f\x0a\x78\x79\xf3\xdb\x35\xe1\xbf\x9a\xd6\xf5\x86\xee\xaa\x8a\x6f\xa6\x75\x0d\xef\x7f\xbd\xb8\xb9\x80\xaa\xfa\x7c\x87\x16\xcf\x94\xa0\x06\x1a\x5f\x4d\x21\xbe\x99\xc2\x8b\xf6\xf5\x33\x79\x8d\xcb\xb6\x13\x11\x2c\x32\x53\x70\xb6\x15\xf1\x3f\x8c\xd4\xe1\x3a\xe1\x2e\x53\xd4\x6e\x5a\x1a\x87\x33\x25\x13\x6c\x2f\x17\x5f\x4d\x47\xd0\xfe\x7f\x33\xe5\x70\x8e\x46\x30\x1c\x0d\xa3\x95\xcf\xfc\xcd\x1e\x84\x2a\xf1\x5a\xe4\xb9\xd4\xf3\x11\x27\xc6\xba\x96\x9f\x4a\x9d\x36\xa2\x9d\x55\xfc\xed\x32\xc7\xd1\xce\x2c\x5d\x29\x5c\x13\xd9\xf4\xa8\x4e\x6f\xe9\x35\x17\x2a\x2e\x6d\x18\x11\x58\xda\xd8\xc4\xd0\xca\x05\x4f\x07\x93\x2c\x92\xa9\x1d\x20\xfb\x28\x19\x66\xed\x5b\x37\x53\xc7\x05\x14\x33\x26\xfa\x52\xa7\xd2\x62\xe2\xc2\x76\xe1\x9f\xb4\xe3\xb7\x2c\x34\x54\xc6\x1f\x84\xea\x75\x4a\x16\x16\x2f\xad\x59\xb4\xe0\x59\x61\x53\x00\x7b\x8e\x89\x7c\xd9\xf2\x48\x0a\xf8\xf0\x51\x6a\x87\x36\x13\x09\x56\xbe\xfb\x13\x6b\x9b\x34\x75\x28\x6c\x0f\xae\x8d\x4f\x9c\x7d\xdc\x74\x47\x87\xbf\xa9\xcc\xfc\x74\x74\x8e\xb7\xe5\xfc\xda\xa4\xc8\x5a\x29\x13\x5e\x72\x26\x28\x1d\xae\xe5\xdc\x35\x6c\xab\x8b\x73\x31\x3a\xbc\x9b\xd8\x59\x8d\x44\x3f\x26\x42\x5f\x89\xc2\xf9\x7a\x7b\x79\xde\xf9\x65\xe0\x6c\x43\xc2\xf3\x0f\xcf\x37\x5b\x87\x58\x34\xd8\x98\x2a\xfd\xaa\xc5\x82\x67\x8e\x66\x6a\xa2\xd9\x87\x67\xcc\xb0\x03\xda\x63\x8a\xe3\x38\x62\x2d\x34\x78\xee\x3f\xdc\x58\x08\x79\xb0\xda\xa3\xa8\x19\xe6\x7b\x3a\x77\xc3\xfc\xd4\x86\xfa\xd7\x01\xdc\x3e\xf6\xf5\xd0\xda\x39\x6f\x47\x4a\xf4\x6b\x3a\xbd\x58\xe8\xb9\xe2\xeb\xe9\xbe\xca\x4e\x33\xc7\x66\xc9\x5d\xb9\xfc\x51\x07\x52\xe0\x2b\x5a\x3d\x07\xa9\xdd\xdf\xfe\xda\x03\x47\x42\x49\x25\x50\x66\x12\x2d\x0f\x42\xfd\xe4\xd8\x93\x1d\xab\x26\x35\x37\xce\x00\xcf\x0a\xcd\x50\x7d\x10\x93\xc7\xd3\xb2\xec\xe3\x21\xee\x6c\x4b\x69\x9c\x79\x94\xb8\x0b\x6b\x67\x4b\x9d\xbc\x14\x52\xad\x03\xde\x28\xfe\xcd\x8b\x47\x8e\x14\xbf\xec\x6a\x05\xf0\x62\xed\x1c\x3a\xd0\xc9\x0a\xfe\x79\x81\xfb\xe1\x4a\x53\x6f\xeb\x5b\xe9\x94\x9f\xa9\x56\xf2\xce\x8f\x79\xb4\xd3\xc4\x1e\x85\xdf\x59\xd7\xc0\xe3\x57\x62\x54\x4c\xe5\xb3\xae\x43\x7f\x67\x7f\xaf\xc6\x1f\x5c\x5f\x7e\xfa\xe9\x71\x7e\xff\x42\xd2\x4d\xc9\x87\x17\x1f\x49\xb6\xa7\x1e\x7f\x18\xf6\xb1\x0c\x3f\x3e\xee\xa7\xde\x93\x64\x23\x14\xc6\xfd\x60\xa8\x38\xc6\xfd\xb4\xbb\x8b\x5e\x9f\x88\x44\x44\xdc\x67\x67\xd4\x4f\x8e\xef\xac\x88\xed\x2c\x71\x44\x51\xec\xdf\xc7\xa7\xe9\x2a\x34\x37\x8b\x45\x27\xe5\x59\xff\x8d\xf9\x1c\xf6\x2d\xee\x52\x18\xcf\x12\xc1\x1d\x97\x1a\x84\xb7\xd0\x2d\x25\x3b\x94\xee\xa8\x25\x5f\x6f\xa0\xe5\xf2\xbf\x50\x5d\x72\x93\x97\xfc\x7a\x4e\xfd\x38\x0c\xdc\xcb\x8a\x3d\xe5\xa6\x1b\x43\x27\x5b\x2f\x80\x23\xde\x13\xed\x8b\xe5\xd0\x5e\x7e\xa1\xc0\xd8\xb3\x74\x9c\xea\xd5\x4b\x65\xb0\xe7\xbd\xbf\xfa\x69\x38\x35\xbf\x64\x0e\xed\x37\xbd\xf5\x9b\xd4\xe9\xb4\x35\x56\xaa\xa9\x2a\xad\x93\xaa\x0e\xfe\x13\x00\x00\xff\xff\x31\x0f\x8d\x91\x27\x18\x00\x00")

func templates17_upsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertTpl,
		"templates/17_upsert.tpl",
	)
}

func templates17_upsertTpl() (*asset, error) {
	bytes, err := templates17_upsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.tpl", size: 6183, mode: os.FileMode(420), modTime: time.Unix(1528396353, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonMysql_main_testTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x5b\x6f\xdb\x3a\x12\x7e\x96\x7e\xc5\xac\x81\x9c\x95\x72\x14\xa6\x40\x81\x7d\x48\x61\x04\x8d\xe3\x14\x41\x9b\x4b\xed\xec\x16\x8b\xa6\xd8\x32\xd2\x38\x21\x2a\x91\x0a\x49\xc5\xf1\x06\xf9\xef\x8b\x21\x75\x73\x62\x19\x2d\xd0\x7d\x3b\x4f\xad\x38\x1f\xe7\xf2\x0d\xe7\xe2\x3c\x70\x0d\xfa\xf6\xf1\x6c\x35\xff\xfc\xe9\x07\xae\x60\x0c\x1a\x6f\xf1\xb1\x64\x67\x95\xb1\x13\x55\x94\x22\xc7\xe8\x7b\x74\x58\xc4\x51\x94\x5c\xcb\xf8\xf0\xda\xfc\x39\xb9\x38\x9f\x5f\xcd\xde\x9f\x9e\x5f\xb1\xdd\xc3\x93\x8b\xd9\xf4\xf4\xc3\x39\x7c\x9c\xfe\x9b\xed\x1e\x5e\xcb\xf8\xcf\xef\x71\x18\xda\x55\x89\x50\xac\xcc\x7d\x7e\x85\xc6\xa2\x06\x63\x75\x95\x5a\x78\x0a\x83\xec\x66\xa2\xa4\x84\x5d\x73\x9f\xb3\xe3\xa3\x90\x0e\xce\x79\x81\x81\xb1\x5a\xc8\xdb\x30\xb8\x53\xc6\xb6\x1f\x95\x41\xdd\x7e\x94\xdc\x98\xf6\xc3\x98\xbc\x50\x59\x77\xad\x54\xda\x06\x42\xda\x30\x0c\x54\x69\x85\x92\x27\x22\x47\xa8\xa5\x61\x60\xd1\xd8\xe3\x23\x32\xd4\x9c\x3d\x87\xe1\xa2\x92\x29\x08\x29\x6c\x14\x7b\xcf\xce\xb8\x90\x30\x86\x3f\x7a\x9e\x3f\x3d\xb7\xc8\xa8\x80\xdd\x9e\x24\x06\x83\xb6\x2a\xa3\x18\x50\x6b\xa5\x49\x03\xb1\x89\x5a\xfb\x83\x30\x0c\x1e\x44\x89\x9a\xcd\xd1\x1e\xe3\x82\x57\xb9\x8d\x46\xee\x3e\xab\x9d\x1f\x25\x30\xb2\xba\xc2\x51\x3c\x0c\xa5\xb8\x46\x09\xbc\x7d\xfb\xe6\x1f\x71\x18\x06\x05\xf3\x74\xc1\x18\xfc\x8d\x0f\x68\xe7\x2e\xa0\xe6\x42\x76\x23\x79\xe1\x54\x16\x8c\xb8\x1c\x46\x92\xd4\xe3\x88\xe6\x61\x1c\x49\x3d\x8e\x32\x30\x8c\x23\x69\x8d\x53\x7a\xcd\xee\xa9\x5c\x8f\xc7\x81\x6a\x12\x86\xf5\x35\x2c\x51\xdc\xc4\xea\x18\x1e\x78\xce\xd9\x11\xde\x0a\xf9\x2f\x9e\x8b\x8c\x53\x9e\xa3\x98\xd5\x1f\x18\x85\x41\xe0\x20\x5e\xcf\xb9\xb2\xd3\xa2\xb4\xab\xc8\x07\x98\xc0\x5a\x3c\xc9\x20\x98\x78\x69\xc1\x9e\xa4\x16\x7c\xae\x6c\xe4\xfe\x33\xbd\xaf\x78\x6e\x22\x1f\x6b\x02\x6f\xda\x0b\x3e\xc0\x2d\xea\x7d\x02\x5b\x7c\x93\xaf\xe1\x0b\x35\x0f\xed\x8d\x96\x97\x24\x0c\x62\x36\xb9\xc3\xf4\x47\x44\x1c\x89\x85\x7b\x7c\x7f\x1b\x83\x14\x39\x3d\xc7\x40\xa3\xad\xb4\xa4\xd3\x30\x78\x0e\xc3\x60\x7f\x1f\x26\x1a\xb9\x45\xe0\xa0\xb9\xcc\x54\x21\xfe\x8b\x19\x64\x37\x40\x3e\x30\xca\x4a\xaf\x50\xc6\x1d\x86\xcd\x2d\xbf\xc9\xd1\x0b\xda\x18\x7a\x46\xc7\x50\xb0\x82\xff\xc0\x8b\xb6\xf6\xa2\xf8\xdd\xb0\x3b\x4a\x1b\xf6\x45\xf3\x32\x42\x4d\x79\x49\x55\x95\x67\xf2\xef\x16\x48\x05\xf8\xfa\x85\x85\xc8\xdd\x33\x7e\x5e\xb7\x92\x69\x55\x5e\x39\x27\xb7\x5a\xa0\x7b\xfd\x6b\xa9\x8b\xfb\x27\x2f\x86\x41\x56\x15\xe5\xa4\xc8\xe0\x60\x0c\xf8\x88\x29\x9b\xa8\xa2\xe0\x32\xab\x9f\x26\x49\x47\x09\x39\xe3\x8b\xd5\xf8\x80\x13\x18\xed\xed\x49\xb5\x97\x71\xcb\xbd\xb8\xa6\x29\xf0\xd6\x87\x15\x0e\x29\x23\x4d\x37\xdc\xa0\x93\x77\xa9\x21\xe2\x75\x02\x4b\xd2\x26\x14\xbb\x14\x25\x46\x71\xeb\x34\x9b\xdb\x4c\x55\x54\x80\xcb\x9e\x65\x3a\x75\xdd\x4d\xe2\xf2\xe4\x23\xae\x8e\xd1\x58\xad\x56\xa8\xa3\x5e\xff\x4f\x40\xaf\x25\xb5\xd3\xc8\xb5\xfd\x95\x8c\x2e\xb8\xc8\x31\x03\xab\xc0\xd0\x55\x68\x69\x83\xd4\x07\xee\x33\xdb\x59\xea\xfb\xf9\x3b\x6c\xad\xdb\xd9\x10\xd2\x17\x2e\x36\x59\x59\x14\x96\x5d\x6a\x21\x6d\x2e\x49\x7d\xfc\x73\x86\x97\x5c\x58\x58\x28\x3d\x14\x67\x18\x2c\xd9\x24\x57\x06\xa3\x18\xf6\xf7\xe1\xfd\x82\x66\x62\xf3\xc6\x84\x81\x4c\x49\x4c\x20\x25\x04\xd8\x3b\x84\xa5\x16\x16\x01\x65\x06\x6a\xe1\x0e\x4a\x51\x62\xb8\x91\xae\xff\x63\x1c\x2f\x63\xa8\x15\x48\x91\x6f\x99\x89\x26\x3f\x53\x19\x46\xae\xb5\xfb\x51\x1b\xd7\xff\x92\x5f\x66\x29\x6c\x7a\x07\x4e\xfa\x14\x06\x29\x37\x58\xcf\xc0\x83\xce\xc3\xd1\x6c\xfa\xf9\x9f\xa7\xb3\xe9\xf1\xa8\x41\x2c\x78\x6e\xd6\x21\xc7\xa7\xf3\xf7\x47\x9f\x1c\xa4\xae\x9b\xbe\xf4\x72\x36\x3d\x99\xce\xbc\x86\x2d\x03\x7c\xbd\xe2\x7a\x6e\xd6\x7a\x88\xc4\x79\x49\x2c\x2e\x22\xaa\xc6\x1a\xbe\x47\xad\x69\xbc\x63\x5c\x55\x76\xdb\x46\x3c\x6c\xe8\x65\x6b\xec\x56\x06\x5b\x94\x89\xcb\x9d\x2b\xe5\xca\x8a\x9c\x5d\x61\x51\x3a\xd8\x88\x16\x04\xaf\xbf\x69\x86\xdb\x7a\xfc\x60\x56\xfd\x63\xd9\xd8\x57\xcd\xd5\xe4\x92\x4c\x3b\x82\xc3\xe0\x3f\x49\xfd\xbe\x94\xa1\x3a\xb4\xf5\x34\xf4\x86\x95\x61\xa7\x86\xe6\xd2\xa3\x30\xd6\x3d\x2a\xe7\x80\xd7\x31\x06\xca\x62\x18\x3c\x03\xe6\x06\xe1\x17\xfc\x74\xcd\x1f\xa4\xb2\x54\xbe\x16\x8a\x76\x2f\x21\x07\x29\x03\x27\x65\xfd\x8e\x1d\x57\xa3\xaf\x69\x2e\x50\xda\x6f\x04\xe9\xc4\x8b\x5a\x4a\x97\xc7\x3b\xe6\x5a\xba\xe4\xd4\xce\xbf\x86\xd1\x94\x1e\xef\x64\x35\x8c\xbe\x36\xc2\x68\x55\xe8\xb4\xd1\xd7\x66\x6d\xdc\x98\xa5\xd2\x59\x07\xa5\x93\x8d\x50\x63\xf2\x3d\x7a\xfd\x1d\xb4\xad\x98\x66\xb0\xc7\x9e\x6e\xcf\x6b\x53\xcb\x2f\x48\x28\xb5\xb2\x2a\x55\xf9\xd8\xa6\xe5\x36\xae\xda\x7e\xf4\x17\x5d\x2f\xe9\xea\x97\x2e\x3d\xdf\xa2\x64\x6e\xbb\x89\xbb\x4e\x47\x67\x75\xdb\x1e\xae\xed\xf5\xcd\xa2\xab\x6c\x6a\xa0\x54\x59\xfd\x1e\x52\x57\x62\x33\xd6\x61\xc7\xbc\x7b\x35\xda\x1b\xe3\x05\xd3\x95\x9c\x14\x59\x64\xee\xf3\x66\xf7\x1b\x6d\xf1\xa3\xbf\x18\x6d\xf7\x82\x90\x9d\x0f\x54\xaa\x54\xd1\xe6\xb7\x7a\x63\x91\xeb\x4c\x2d\x65\xdf\x17\xb1\x70\x3b\x91\xfb\xfd\xd7\xeb\x0c\xcd\x59\x4b\x75\x7f\x60\x1f\xfc\xe2\xce\xd7\x3a\xac\x0c\x9b\x61\xa1\x1e\xe8\xa5\xfc\x54\x8f\x6e\xe2\xa3\x35\x29\x69\x46\x5f\x3d\x13\x12\xe0\xfa\xd6\x00\x63\xac\x19\x69\x6d\x50\x4e\x30\x06\x5e\x96\x28\xb3\xe8\xeb\x37\x0f\x78\x7a\xb9\xce\x3d\x7b\x15\x8c\x31\x7a\x5f\xe9\x86\x4d\xb0\xb6\xd8\xc3\x11\xac\x5d\xdb\xbc\x5e\xc3\xce\x71\x39\x43\x9e\xa1\xf6\x9e\x92\x36\xe3\xf7\xbd\x83\x31\xfc\x71\xb3\xb2\x68\xd8\x51\xb5\x58\xb8\x9f\xae\x24\xaa\x59\x7c\x25\x4a\xfb\x9b\xa2\x57\xd1\x1e\xfa\x21\xe0\x2f\xf7\x53\x41\xe2\x59\x25\x5f\x67\xa1\xbf\x6e\x34\x93\x47\x57\x52\x0a\x79\x7b\x30\x6a\xd9\xf4\xb1\xc5\xeb\x70\x6f\xba\xfe\xd9\x13\xc5\xaf\xa5\xa8\x75\x5f\x3a\x90\xef\xad\x5b\x49\xaa\x24\xbd\xc4\xa8\xfe\xab\x43\xe2\xd3\x17\x0f\x3f\xca\xf6\xd5\x7b\x49\xe2\xd4\x3b\x73\xeb\xbf\xf1\x83\x0e\x51\x73\x76\x9f\xb3\x8b\x12\x65\xb7\xd8\x67\x5a\x3c\xa0\x66\x6e\xc5\x3e\xaa\x44\x9e\x7d\xae\x50\xaf\xea\x80\x9a\x9f\xa6\xbe\x05\xae\x17\x5f\xd3\x91\x9b\x96\x5b\x77\xbf\x5e\xcf\x5b\xcf\x41\x47\x44\xf2\x8a\x9d\xf5\x40\x9e\xc3\xff\x05\x00\x00\xff\xff\xaa\xfc\xbe\x6b\xf8\x11\x00\x00")

func templates_testSingletonMysql_main_testTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_main_testTpl,
		"templates_test/singleton/mysql_main_test.tpl",
	)
}

func templates_testSingletonMysql_main_testTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_main_testTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_main_test.tpl", size: 4600, mode: os.FileMode(420), modTime: time.Unix(1528388322, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonMysql_suites_testTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\xb1\x8e\x83\x30\x10\x44\x7b\xbe\x62\x84\x28\xe0\x04\xfe\x80\x93\xae\xba\xea\xae\x48\x11\x91\x0f\x70\xc2\x82\x2c\x39\x1b\x84\x17\x29\x92\xf1\xbf\x47\x18\x8b\x90\xce\xe3\x99\xb7\x3b\xdb\xcf\x7c\x43\x4b\x4e\x2e\xa3\xa3\x49\x4a\xc1\x97\x90\x13\xc3\x83\x6a\x2b\xf8\x0c\xf0\xbe\xc1\xa4\x79\x20\x14\x86\x3b\x7a\xd6\x28\x44\x5f\x2d\xe1\xfb\x07\xaa\x5d\x5f\x2e\x84\x94\x33\x7d\x32\xd5\x9f\xfb\x7f\x18\x8e\x36\x9a\xdd\x27\xeb\x8e\x72\xcb\x9e\xf4\x3d\x0e\x4b\x64\x94\x0b\x46\x3b\x4f\xda\x62\x81\x18\xb1\xf4\xab\x77\x50\xd4\x79\xe6\x32\xf7\xfe\x4d\x87\x90\xd7\x58\x6b\x7f\x7e\x6e\x27\x55\x71\x19\x71\x77\xec\x91\x54\xc8\x5e\x01\x00\x00\xff\xff\x2f\xea\xf2\xb5\x00\x01\x00\x00")

func templates_testSingletonMysql_suites_testTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_suites_testTpl,
		"templates_test/singleton/mysql_suites_test.tpl",
	)
}

func templates_testSingletonMysql_suites_testTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_suites_testTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_suites_test.tpl", size: 256, mode: os.FileMode(420), modTime: time.Unix(1528388322, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonMysql_upsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xdf\x8f\xd3\x30\x0c\xc7\x9f\x93\xbf\xc2\x54\x3a\x5d\x2b\x45\x3d\xee\x15\x69\x0f\x77\x6c\x9c\x0a\x63\xbf\x07\x42\x88\x87\x6c\x71\xb6\x48\x5d\x3a\x12\x67\x30\xa1\xfd\xef\x28\x5d\xb7\xf5\x8e\x21\xf1\xc0\x4b\xeb\xc4\xf6\x37\xfe\xd8\xbe\xbb\x83\x45\x30\xa5\x9a\x6f\x3d\x3a\x1a\x07\x74\xfb\x8f\xfb\xe9\xb8\x7f\xbc\xf5\x20\x21\x1e\x3c\x49\xc2\x0d\x5a\x02\x4f\xce\xd8\x15\x04\x1f\xbf\xb4\x46\x08\x75\x62\x57\x92\x84\xad\xab\x76\x46\xa1\xca\xb9\x0e\x76\x79\x5d\x37\x55\x46\x82\x72\x66\x87\xce\xe7\x5d\x23\x4b\x5c\x92\x00\x92\x8b\x12\x07\x72\x83\x8d\xbe\x80\xb0\x55\x92\x50\xc0\x8f\xb5\x21\x2c\x8d\x27\xf8\xfa\xed\xe8\xcb\x4e\x35\xfc\xe2\xec\xe2\xed\xc4\xdb\x8d\xb4\xab\x12\xf3\x42\xa1\xa5\x71\xa8\x08\xa7\xa5\x59\x62\x7c\x32\xef\x8f\x05\xc4\xff\x64\xdc\xd2\xcc\x38\x67\x8b\xa0\xe1\x4d\x3b\xfb\x09\xe9\x31\x68\x8d\x2e\xcd\x38\x53\xa8\xd1\xb5\x9c\xa3\x70\x72\x2e\x82\x8e\xe9\x3b\xe9\x60\x59\x95\x61\x63\x7d\x53\x17\x67\x46\x43\x89\x36\xbd\x3c\x03\xaf\x3a\xf0\x3a\xd6\xcb\x4e\xa1\x9d\x26\xd8\xe7\xef\x2b\xd3\x0a\x15\x90\x08\x48\x32\xce\x0e\xfc\xac\x73\x6c\x45\x06\x9d\x93\x88\xde\x50\xfe\x6e\xeb\x8c\x25\x9d\x72\xc6\x22\x82\x88\xff\xa4\x18\x4c\x7b\x93\x19\x14\x4f\x83\xe1\xa4\x07\xc5\x60\x36\x84\x1b\x0f\xe9\x8d\xcf\xe0\xd3\x43\x7f\xde\x9b\xd6\x76\x52\x07\x9f\x5b\x5e\x9f\x9a\xba\x6a\xbb\x45\x5b\xca\x25\xae\xab\x52\xa1\xf3\x75\x17\xe7\x1e\x0b\xab\xf0\x67\xdb\x21\x5e\xc0\x0a\xb8\x17\x70\x9f\x45\xa9\x8c\x33\xe6\x90\x82\xb3\xb0\x08\x3a\x9f\xd6\xc8\x69\x43\xf7\x82\xa2\x81\x38\x33\xfc\xa5\x78\x18\x0e\xa0\x3b\x1f\xf5\x8b\xb7\x0f\xb3\x1e\x7c\xe8\x7d\x81\xf9\xa8\x1b\xcd\x9a\xea\x19\x54\x8b\xe9\xbf\x21\xc5\x91\xeb\xca\x81\x11\xb0\x8b\x6b\xe3\xa4\x5d\x61\xb3\xac\xf5\x6c\x8c\x06\x73\x19\x77\xa4\xca\x3f\x3b\x43\xf8\xb8\x27\x4c\x6f\xc5\x6d\x6c\xc9\x81\x33\xf6\x3d\xae\xa7\x7a\xbe\x79\x97\xbd\xfd\x63\x65\x77\x19\x6f\x89\x35\x8d\x3c\x6a\x5c\xf3\x24\xd0\x69\x9a\x96\x26\xff\x98\x79\x2c\x30\xbb\x6d\xa6\x73\x6d\x6c\x07\xfe\x3b\x00\x00\xff\xff\xa9\x3a\x4a\xd3\x2e\x04\x00\x00")

func templates_testSingletonMysql_upsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_upsertTpl,
		"templates_test/singleton/mysql_upsert.tpl",
	)
}

func templates_testSingletonMysql_upsertTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_upsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_upsert.tpl", size: 1070, mode: os.FileMode(420), modTime: time.Unix(1528388322, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testUpsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\xcd\x6e\xdb\x3c\x10\x3c\x8b\x4f\xb1\x9f\xf1\xb5\x90\x0a\x85\x41\xaf\x29\x7c\x48\x9c\x1c\x82\xa2\x86\x11\xcb\x0f\xc0\x48\x2b\x87\x30\x4d\x0a\xe4\x2a\xb1\xcb\xf0\xdd\x0b\xca\x76\xfc\x0f\x18\x28\x7a\xc8\xc1\x06\xb8\x9e\x9d\x99\x1d\x2e\xed\xfd\x15\xfc\x4f\xe2\x59\xe1\x50\xcc\x71\x2c\xf5\xb4\x55\xc2\xc2\x4d\x1f\x78\x11\xab\x3c\x96\xe1\x1d\xdc\xe6\x97\x77\x20\x49\x0a\x07\xc2\x21\x5c\x85\xc0\xf6\x09\x46\xaa\xb5\x42\x1d\xb7\x37\xab\xfa\xc9\xe6\x57\x61\x2f\x6a\x2d\xc5\x1c\xd5\xc9\xd6\x8b\x6c\xef\xb7\xd7\xad\x2e\x81\xd0\x91\xf7\x87\xee\x43\x98\x34\x0e\x2d\xa5\x04\xdf\x22\x42\xea\x29\x2f\x32\xf0\x2c\x21\x3e\x12\x56\x28\x85\x2a\xcd\x18\x4b\x64\x0d\x0a\x75\xea\xfd\xa1\x8f\x10\x06\x46\xb5\x73\xed\x32\xe8\xf7\xcf\x62\x46\x56\xce\x85\x5d\xfe\xc4\xe5\x07\xda\xb3\x24\x21\x3e\x9e\xc9\x26\xed\xc5\xef\x46\xea\x29\x74\xf6\xe0\x4d\xd2\x0b\x18\xad\x96\xd0\xac\xfa\x60\x86\x4b\x28\x57\x9d\xbd\x8c\x25\x81\xb1\xc4\x21\x56\x31\x04\x2b\x74\x65\xe6\xf2\x37\xf2\x21\xbe\x8d\x11\xab\x34\x63\xc9\xab\xb0\x80\xb6\xfb\x18\xcb\x92\xeb\x6b\xb8\x25\xc2\x79\x43\x40\x2f\x08\x8f\xc3\xf1\xc3\x53\x01\x4e\x56\x08\xa6\x06\xa1\x61\x32\x8a\x15\x96\x98\xc8\xb8\x1b\xd3\x76\x04\x1f\xba\x14\x22\xeb\xae\xe8\x98\x6c\x5b\x52\x1a\xdd\xe4\xf0\xd5\xe4\x70\x6a\xfc\xfb\xbb\x62\xd9\xa0\xcb\x81\x6c\x8b\xd9\x8f\x8e\xe4\xbf\x3e\x68\xa9\xd6\x31\x3c\x44\x9f\x75\xda\x9b\xe8\x2e\x00\x32\x5b\x85\x33\x76\xc0\x75\xc2\x37\xf0\xc5\xf5\xf2\x48\xb8\x8e\xc5\x7b\x59\x83\x36\x04\x7c\x68\x06\x46\x13\x2e\x28\x84\x92\x16\x71\xb0\x72\x75\xe6\x77\xa2\x9c\x4d\xad\x69\x75\x95\x66\xde\xa3\xae\x42\x60\xc9\x0a\xf2\xab\x75\x54\x2c\xd2\x8e\x65\x97\xe1\xa8\xf0\x6c\xa4\xe2\x77\x38\x95\xba\xe3\x50\x0e\x77\x6b\xc5\x22\x2d\x69\x91\xc7\x09\x37\x0a\x17\x81\x32\x96\x54\x58\xa3\x05\x5a\xf0\x27\xa3\xd4\xb3\x28\x67\xf1\x3e\x3f\x82\x37\x7c\xbd\xb0\xe7\xe6\x8c\x17\x80\xba\x8a\x8b\x0f\xf1\xd4\xc9\x3d\xea\x1a\x6d\x9a\xed\x9f\x2e\xbb\x87\xb6\x93\x3b\x73\x09\x47\xe9\x97\xa6\xd5\xd4\x15\x0e\x17\x69\xf3\xde\xd2\x8c\x0f\x22\xe8\xc2\x01\xb6\xb3\x1f\xfb\x4c\x37\xba\x11\xd2\x29\x47\xd0\xf7\x3d\x48\xef\x4d\x68\x02\xa3\x11\x2c\x96\xc6\x56\x39\x4c\x0d\xdd\xf4\xf2\x15\x7e\xed\xfa\xe0\x79\x4c\x46\xf7\xb7\xc5\xc3\xa9\xe7\xf1\xd7\xfb\x5f\x0b\xe5\xf0\x34\xe6\xe8\x2f\x82\x73\xfe\x6f\x9f\xca\x27\x5c\xaa\xcf\xb2\x53\x81\xfd\x09\x00\x00\xff\xff\xf3\x21\xaa\x9c\x6d\x07\x00\x00")

func templates_testUpsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertTpl,
		"templates_test/upsert.tpl",
	)
}

func templates_testUpsertTpl() (*asset, error) {
	bytes, err := templates_testUpsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.tpl", size: 1901, mode: os.FileMode(420), modTime: time.Unix(1528388322, 0)}
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
	"templates/17_upsert.tpl": templates17_upsertTpl,
	"templates_test/singleton/mysql_main_test.tpl": templates_testSingletonMysql_main_testTpl,
	"templates_test/singleton/mysql_suites_test.tpl": templates_testSingletonMysql_suites_testTpl,
	"templates_test/singleton/mysql_upsert.tpl": templates_testSingletonMysql_upsertTpl,
	"templates_test/upsert.tpl": templates_testUpsertTpl,
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
		"17_upsert.tpl": &bintree{templates17_upsertTpl, map[string]*bintree{}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"mysql_main_test.tpl": &bintree{templates_testSingletonMysql_main_testTpl, map[string]*bintree{}},
			"mysql_suites_test.tpl": &bintree{templates_testSingletonMysql_suites_testTpl, map[string]*bintree{}},
			"mysql_upsert.tpl": &bintree{templates_testSingletonMysql_upsertTpl, map[string]*bintree{}},
		}},
		"upsert.tpl": &bintree{templates_testUpsertTpl, map[string]*bintree{}},
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

