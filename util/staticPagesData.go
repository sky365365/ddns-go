// Code generated for package util by go-bindata DO NOT EDIT. (@generated)
// sources:
// static/pages/writing.html
package util

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _staticPagesWritingHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5a\xff\x6f\x1c\xc5\x15\xff\xdd\x7f\xc5\x30\x58\xf2\x5d\xe5\xdd\xb5\xcf\xce\x89\x9a\x5b\x57\x90\xa3\xc5\x6a\x08\x51\x9c\x20\x7e\xb3\xe6\x76\xe7\x6e\xc7\x9e\x9d\xd9\xcc\xcc\x9e\x7d\x39\x9d\x14\x68\x43\x29\x02\x29\xa8\x20\x51\x0a\x0d\x54\x80\x50\x55\x92\xfe\xd4\x22\x68\xc8\x1f\x83\xcf\x76\x7e\xea\xbf\x50\xcd\xec\xed\x97\xbb\x5b\xc7\x07\x71\x82\x8b\x7f\xf0\xcd\x97\x37\xef\xbd\x79\x9f\x37\xef\xcd\xee\xdb\x46\xa0\x42\x0a\x28\x62\x1d\x17\x5e\x0f\xe0\xfa\xdc\x5c\x23\xc0\xc8\x5f\x9f\x03\xa0\x11\x62\x85\x80\x17\x20\x21\xb1\x72\x61\xac\xda\xd6\x33\x30\x9f\x08\x94\x8a\x2c\x7c\x2d\x26\x5d\x17\xbe\x6a\x5d\x7d\xce\x3a\xcf\xc3\x08\x29\xd2\xa2\x18\x02\x8f\x33\x85\x99\x72\xe1\xc6\x0b\x2e\xf6\x3b\xb8\xb0\x8e\xa1\x10\xbb\xb0\x4b\xf0\x6e\xc4\x85\x2a\x90\xee\x12\x5f\x05\xae\x8f\xbb\xc4\xc3\x96\xe9\x2c\x02\xc2\x88\x22\x88\x5a\xd2\x43\x14\xbb\xcb\x53\x6c\x50\xac\x02\x2e\x0a\x4c\xb6\xc9\x48\x96\x22\x8a\xe2\xf5\x66\xf3\xe2\xa6\xf5\x9b\x97\x1b\x4e\xd2\xd5\x13\x4f\x59\x16\x78\x9e\x73\x25\x95\x40\x11\x38\xbf\xb9\x09\x2c\xcb\x4c\x50\xc2\x76\x80\xc0\xd4\x85\x52\xf5\x28\x96\x01\xc6\x0a\x82\x40\xe0\xb6\x0b\x1d\xa9\x90\x22\x9e\xd3\x4a\x17\xda\x21\x61\xb6\x27\x25\x9c\x79\xa9\xc7\xc3\x90\x17\xd6\x48\x4f\x90\x48\x01\x29\xbc\x9c\x66\xfb\x5a\x8c\x45\xcf\x5a\xb1\xcf\xd9\xcb\x46\xc2\xb6\x84\xeb\x0d\x27\x21\x5d\x9f\x6b\x38\x09\x36\x73\x8d\x16\xf7\x7b\x86\x8b\x1e\xc0\x42\x37\x01\x68\xf8\xa4\x0b\x3c\x8a\xa4\x74\x21\x43\xdd\x16\x12\x20\xf9\xb1\x7c\x24\x76\x40\xab\x93\xfc\xca\x00\xf9\x7c\xd7\x92\x21\x4c\x96\x8d\x2f\xd4\x96\x44\x84\x61\x01\x7c\xab\x4d\xf1\x1e\xd8\x8e\xa5\x22\xed\x9e\x35\x32\xb1\xd5\xc2\x6a\x17\x63\x96\x2d\x06\xa0\x81\xd2\xad\xc2\x71\xf1\x56\x4b\x20\xe6\xa7\x8c\x10\x25\x1d\x66\x11\x85\x43\x69\x79\x98\x29\x2c\x0a\x3c\xb4\x41\xba\x1d\xb0\x17\x52\x26\x5d\xa8\x7d\x6b\xcd\x71\x76\x77\x77\xed\xdd\x15\x9b\x8b\x8e\x53\x5b\x5a\x5a\x72\x64\xb7\x03\x41\xe2\x26\xb0\xb6\x04\x41\x80\x49\x27\x50\x49\xbb\x4d\x28\x75\x21\xe3\x0c\x43\x20\x95\xe0\x3b\xd8\x85\x5e\x2c\x04\x66\xea\x3c\xa7\x5c\xc0\x82\x28\x30\xa2\xb0\x28\x61\xd8\x43\x91\x0b\x05\x8f\x99\x0f\x8b\xc3\xdb\x9c\xb0\xc9\xf1\x54\x34\x04\x48\x10\x64\x05\xc4\xf7\x31\x73\xa1\x12\x31\xce\x76\x1e\x0a\xab\x36\x2e\x4b\xbb\xfa\xf3\x7c\xcf\x85\x4b\x60\x09\xd4\x56\x41\x6d\x15\x82\x36\xf7\x62\x89\x5a\x14\xbb\xb0\x8d\xa8\xc4\x63\x96\x00\xa0\x11\x21\x15\x00\xdf\x85\x2f\xd5\x56\xc0\xf2\x2f\x51\x0d\xd4\x80\x5e\xbd\x6c\xd5\x40\xed\xc5\x95\x62\xdf\xaa\xbd\xf2\x4c\xde\x07\x35\xab\x16\xac\xd2\x9a\xb5\x12\xd4\x69\x0d\xac\x04\xab\xc5\x39\x50\xbb\x0e\x81\x33\x21\xca\x23\xc2\xa3\x18\x78\x7b\x2e\x5c\xae\x41\xe0\xf5\x5c\xb8\xbc\x02\x81\x70\xe1\xaa\x26\x6e\x68\xab\x8f\xe3\xa4\x04\x67\x9d\xfc\x68\x8d\xfa\xb9\x3b\x38\x28\x73\x2c\xc7\x27\xdd\x91\x73\xa6\xcd\xc4\x89\x13\x9f\x35\xa7\x19\x11\x06\x04\xd7\xa6\xd0\x4d\x6d\xec\x5e\xd2\x11\x1d\xc2\x2c\xc5\xa3\x35\xb0\x7c\x2e\xda\x83\xd3\x4e\x2e\xf8\xee\x31\x3e\x4c\xad\xd0\xb7\xea\x80\xb7\xdb\x12\x2b\xdd\x5e\x29\xfa\x6b\x9b\x8b\x70\x7d\xae\xb8\xa7\x56\xac\x14\x67\xe9\xfa\x96\x62\xa0\xa5\x98\x15\x09\x12\x22\xd1\x03\x32\x6e\x85\x44\x6d\xb5\xd4\x94\x76\x2d\xae\x14\x0f\x13\x05\x9f\x85\xeb\x9b\xa8\x8b\x1b\x4e\xc2\x6c\x5c\x40\x41\x3b\x44\xb1\x50\xc0\xfc\xb7\x64\xec\x79\x58\xca\x8c\xad\x4f\x64\x44\x51\x6f\x0d\x68\x47\x7e\x76\xd2\x2b\x12\x4b\x03\xe2\xbb\x50\x60\x19\x53\xf5\x92\xec\xc0\xf5\xfd\xfb\x1f\x0f\xbf\xfa\xe0\xe0\xcd\x5b\xc3\xb7\x6e\x4f\xa3\x91\x59\xfe\x18\x6d\x74\x10\xa6\x58\x4d\x8a\x0a\xce\x4d\x10\x6c\x6d\x69\xd8\xe0\x7a\xf3\xe2\xe6\xc1\x47\xef\x0c\xdf\xfa\x74\xf8\xfe\x1b\x0d\x27\x38\x37\xb1\x6e\x9a\xf3\xd6\x96\x0e\x56\x70\x4c\x81\x49\x52\x0d\x88\xd5\x11\x3c\x8e\x40\x11\xd3\x02\x31\x45\x2d\x4c\x8b\xf8\xca\xd0\xaa\x01\xdd\x30\x6b\xcd\xb4\x0e\x95\xa6\x51\xb2\x7e\xc2\x3b\x64\x68\x2d\x2f\x95\xc8\x29\x51\xcb\x0b\xb0\xb7\x03\xf2\xa6\x45\x98\x8e\x0f\x53\xb2\x4b\x78\x01\xd0\x20\x2c\x8a\xd5\x34\x3f\xcb\x8c\x43\xa0\x7a\x11\x76\xa1\x40\x3e\xe1\x70\x94\xd2\x9a\x4c\x5e\x44\x21\x86\x06\x66\x44\x89\xcf\x24\x04\x5d\x44\x63\x9c\x77\x39\xf3\x28\xf1\x76\xd2\x81\xf3\x9a\x25\xf6\x7f\x1d\xb3\x4a\x15\x82\x7e\x9f\xb4\x01\xbe\x06\xe6\xed\xe6\xc5\x4d\x5b\xf3\x02\xe9\xc2\xc1\xc0\x4b\x48\xfb\x7d\xcc\xfc\xc1\xe0\x18\xa5\xc7\x8c\x5d\x50\x3a\xd9\xa9\xb6\x45\xa6\x4a\x39\x07\x00\x9e\x33\xd3\x95\x07\x1f\xdc\x7f\xf0\x87\xb7\xf7\xbf\x79\xb7\x5a\x2e\xe9\x38\xbc\xc6\x42\xc7\x59\x06\xc8\x67\x32\xe2\x7e\x06\x50\xda\xcd\x00\x4a\x06\x4e\x06\x68\xb4\xf0\x14\x01\x1a\x71\x3c\x0e\xa0\xa6\x99\xae\x1c\xdd\xfc\xee\xe8\xce\xdd\x9f\x31\x40\x1e\xe5\xb1\xdf\xa6\x48\xe0\x0c\xa4\xe2\x50\x06\x54\x3e\x78\x32\x58\x05\x06\xa7\x08\x58\x81\xeb\x71\xa0\x9d\xcf\x48\x4e\x0f\x2d\x19\x22\x4a\x53\x5f\xde\x0a\x30\x8d\xe0\x98\xa2\x0a\xef\x29\xa0\xff\x59\x61\xac\xb0\x6f\xee\xa3\x7a\x49\x49\x90\x2d\x93\x31\x9d\x7d\xc0\x8f\x0c\xfe\xc6\x48\x4d\x26\x37\x9a\xd9\xd1\xdb\xf0\x2f\x24\x06\x3c\x29\x31\x6c\x34\x4f\x23\x35\x94\xb8\x22\x67\x4a\x70\x5a\xf0\xbc\x54\xbb\x51\x73\xe4\x72\xfd\xbe\x71\x9f\x8d\xe6\x60\x50\xb6\xc1\x27\x63\xb8\x4d\xec\x09\xfd\x5c\x32\x32\x5e\xd2\x9d\xd1\x80\x09\xf1\x93\x32\x62\x51\xd3\x42\x77\xdc\x98\xc9\xe8\xa3\x19\x74\x8a\xf0\x74\x2f\x4b\x1b\x97\x5e\x59\x7d\xe2\x97\x24\x03\x37\x89\xba\xab\x5b\x98\xa1\xe4\x69\x7c\x1c\x5c\xb8\x7e\xf0\xc1\xdd\xe1\xad\x2f\x86\xb7\xee\x1e\xbe\xf7\xe5\xe9\x61\x9a\x84\x62\x13\xd8\x5a\x7c\x0f\x96\xc6\x6c\x1d\xf8\x4b\xaf\xf8\xe6\x02\x6d\x00\x1f\x53\x3d\x71\x89\x8d\xa8\xbb\xfa\xc2\x68\x24\x0f\xc8\x7a\xd4\x4e\x86\x81\x7e\x06\x3b\x39\x16\x3f\x89\x83\x66\xd4\x8f\xc5\x0c\x67\xea\xe0\xf6\xe7\x47\x77\xff\x76\xf5\xf2\x85\xd3\x86\xa0\x28\xbd\xec\x84\x69\xc3\x5d\xd5\x34\x99\xb9\xcd\x0a\xf3\x50\xeb\x63\xe9\x09\xd2\xc2\x7e\xab\x97\xcf\x8d\xd2\x42\x7e\xfe\x8c\xe9\xaf\x5e\xbe\x50\x7a\xfa\xc6\x92\xca\x04\x8b\x87\x66\x96\x83\xaf\x3e\xbb\x7a\xf9\xc2\x83\x8f\x6e\x1c\xdd\x7f\x6f\xf8\x97\xbf\xee\xdf\xfb\xe4\xf0\xc3\xdf\x0f\x6f\xfe\xe3\xf0\xde\xbb\xfa\x24\x0d\x3f\xfa\xe7\xf0\xe3\x1b\xdf\xdf\x78\x7d\xf8\xc5\xeb\xff\xfd\xcf\x87\x81\x52\x91\x5c\x73\x1c\x14\x11\x4b\x4b\xb1\x49\x64\xcb\x96\x43\xa2\xef\x6f\xbc\x96\xce\x85\x3d\x12\xd9\x24\x22\x91\xcd\xb0\x2a\x8c\x47\x5d\x5b\xf2\x20\xb6\x3d\x1e\x3a\x1e\x51\xbd\x6d\xc9\xd9\xaf\x08\x76\xcd\x4b\xae\x9f\x34\xc1\x19\x7b\xf9\x5c\x3f\x0a\xcb\x93\x5d\xa8\x99\x10\x9e\x86\x03\x69\x20\x90\xc0\xa8\xdc\x71\x32\x28\x33\xd5\x72\x57\x6a\xa6\x43\x82\xef\x4a\x17\xae\x1c\xeb\x49\xa3\xb5\x89\x2b\xac\xcf\xf5\xfb\x16\x10\x88\x75\x30\x98\x27\x8b\x60\xbe\x0b\xd6\x5c\x90\x78\xd6\x88\xe3\x60\x30\xd7\xef\xcf\x77\xcd\x8f\x05\x30\xf3\x81\x35\x18\x94\x5e\x6e\x52\xe5\x67\x71\xc6\x31\x2d\x1e\xee\x90\xfb\x5f\xdf\x38\xfa\xf4\xed\xfd\xaf\x6f\xec\x7f\xfd\xf7\xe1\xed\xdb\xc3\x5b\xef\x3c\xba\x6f\x3c\xfe\x94\x53\xff\x61\x29\x67\x5c\xe5\x53\xf4\xe3\xfa\xff\x6f\x02\xaa\x4f\x27\xa0\x7a\x69\x02\xaa\x9f\xc9\x04\x54\x3f\xb3\x09\x28\xb3\xaf\x21\xc9\x8d\x7b\xf5\xd8\x04\x54\x9f\x4e\x40\x13\x91\xbf\x9e\x45\xfe\x19\x92\x51\xfd\x74\x92\x51\xfd\xe1\xc9\xa8\x5e\x96\x8c\xba\x75\xdb\xe4\x23\x8a\x1c\x9d\x6f\x0a\x33\x32\xc2\xd8\xb7\x19\x8e\xeb\x36\xf6\x63\xdb\x63\x4e\x07\xab\x8d\x4b\x76\x14\x44\x3f\x75\x32\xaa\x9f\xdd\x64\x54\x2f\x49\x46\xf5\x19\x93\x51\x7d\xd6\x64\x54\x7f\x9c\xc9\xa8\xfe\x33\x4f\x46\xc3\x9b\xff\x1a\xde\xf9\xdd\x83\x9b\xef\x1c\xde\xbb\x73\x16\x92\xd2\x55\x89\x05\x33\xaf\x86\x4e\xf2\xe5\xc3\x3f\x7f\x3b\xbc\xf7\xfe\xe1\x7b\x5f\x1e\xbc\xf9\x6f\x63\xe9\x27\xf0\xe0\x9b\x6b\xa7\xdd\x23\xef\xe5\xd7\xee\x74\x6c\x30\x28\x73\xeb\x74\x36\x75\xe9\x12\x1d\x72\xef\x1b\x27\x3e\xc9\xf3\xbe\xd9\xbf\xff\xf1\xc1\x5b\x9f\x25\x41\x70\xff\xfe\xa7\x07\xaf\xdd\x1d\xde\xf9\xe3\xf0\xe6\x97\x8b\x60\xf8\xed\x37\x47\x77\xee\x1c\x7d\xf7\xa7\xe1\xcd\xcf\x7f\xd2\x60\x75\x09\x49\xb9\xcb\x85\x3f\x2b\xb8\xc3\xbb\x6f\x1c\x7e\xf2\xda\x63\x46\x36\xc9\x87\x51\xa6\x5a\x82\x74\xae\xaa\x86\x22\xef\xe5\x48\xa7\x63\xe5\x48\xa7\xfc\x66\x42\x7a\x9c\xf8\xcc\x20\x3d\x4b\xe8\x79\xbc\x65\xb9\x86\x93\x14\x02\x4b\xf4\x29\xa9\x26\x8e\x55\x10\xa3\xdc\x8c\x4c\x59\xbb\xa6\x0e\x6d\x51\xfd\x3f\x31\x67\x4b\x60\xb4\x53\x5e\xca\x34\x3a\x99\x65\x92\x5c\xc7\x6b\x60\x79\x25\xbb\x78\x52\xde\x31\xf5\xfe\xe8\xe1\xf5\x53\x9d\x31\xcc\x36\x1a\xe9\x87\x01\x9a\x60\xbe\xd2\x8e\x99\xa7\x08\x67\x95\x6a\x7f\xb4\x7e\xbe\x02\xed\x82\x95\xaa\x36\x67\x95\x05\xf3\xc2\x7b\x61\x31\xa3\xc6\x55\xd0\xcf\x76\x86\xed\x48\xe0\x2e\x66\xaa\x89\xdb\x28\xa6\xaa\x52\x7d\x36\x9b\x9b\xaf\x2c\xe8\xe8\xbc\x50\xb5\x11\x23\x21\x52\xb8\xd2\x07\xd2\x13\x9c\xd2\x2b\x7a\x67\x4b\x60\xb0\x08\x56\x96\x96\x8a\x2b\x6c\xb4\x8d\xf6\x2a\xfd\x02\xa6\x21\x56\x01\xf7\xd7\x00\xbc\xf4\xf2\xe6\x15\xb8\x58\x98\x89\x05\x5d\x03\xd0\x91\xa8\x8b\xc7\xc6\x7d\xa4\xd0\x9a\x16\xae\xd1\x5a\xa8\xda\x12\x0b\x82\x28\xb9\x8e\x2b\xd5\x22\xd9\xa8\x94\xba\x06\xd2\x8d\x81\x4a\x52\x24\x2d\x6e\x6f\xb4\x0d\xdb\x54\x5f\x17\xaa\xb6\x27\x65\x25\x2d\xba\xc2\x45\x00\x5b\x94\x7b\x3b\xb0\xb0\x03\xfd\x47\xda\x29\x2b\xf0\x94\xeb\x02\xc8\x77\xe0\x24\xcf\x71\xae\xc8\xf7\xcf\x6b\xf7\xa8\x24\xb5\x5e\xcb\xd7\xf7\x0a\x01\x27\x4b\x2c\xf3\x95\x85\xa7\xb3\x3a\xee\x42\xd5\x0e\x54\x48\x53\x9d\xc7\x48\x07\x00\x53\x89\xa7\x44\x3a\x0e\xe0\x3b\x13\x63\x12\xab\x2b\x24\xc4\x3c\x56\x25\xee\x30\x9b\x0d\xcc\x37\x14\x13\x26\x00\x23\x6c\x97\x26\x14\x9b\x2b\x12\x14\x3a\x58\x08\x2e\x72\x28\x2a\xdb\xd7\x5e\x7d\xf1\xf2\xa4\xd1\x8c\x02\xc9\x94\x2d\x15\x52\xb1\xbc\x82\xf7\xd4\x98\xe4\x5c\xc0\x20\x95\x3c\xa8\x26\x27\x38\xfd\xed\x22\x01\x5a\xb8\xcd\x05\x36\xaf\xdc\x5d\x08\x93\xf1\xcc\x0d\xa6\x6b\xa4\x99\x1e\x3e\xf7\xe2\x10\x33\x65\x77\xb0\x7a\x81\x62\xdd\x7c\xbe\xb7\xe1\x57\x46\x6f\xef\xab\xb6\x4f\xcc\x57\x1a\xbe\x0b\xcc\x67\x1a\xa3\x65\xda\x21\x0a\x22\x73\x9b\x9c\xc8\x2f\x09\xf0\x45\x7d\xe7\x4e\x52\xa5\x50\xe6\xa8\xda\x84\x31\x2c\x5e\xbc\xf2\xd2\x05\xe0\x02\xf8\x9c\x71\xf8\xdf\xe2\x1e\xd8\x68\xc2\x19\xd8\x14\x5f\xf8\x1f\xcb\x6a\xf4\xae\x7d\x06\x76\x49\x3e\x99\x60\x94\x3e\xd1\x08\x14\xda\x1e\x67\x92\x53\x6c\x23\x4a\x7a\x31\x33\xef\xda\x42\xc4\x50\x07\x3b\x68\x27\x11\x30\x98\x80\x6a\xba\x5a\xfa\x33\x81\xea\x91\x01\xba\xc2\x77\x30\x7b\x64\x54\x52\x44\x12\x3b\xeb\xe7\x4c\xe4\x79\x3c\x66\xca\x51\x39\xff\x49\x50\xca\x2b\xa3\x27\x03\x73\xbc\x35\xd2\x7d\x14\x6c\x0b\xdc\x19\x11\xf9\xe1\xee\xa0\xc4\xec\xab\x46\xa0\xc3\x33\x81\x96\x8f\x64\x60\xe7\xd6\x37\x07\x28\x12\xbc\x4d\x28\x36\xef\x19\x0c\x66\x72\x0c\x34\x1d\x0c\xfd\xa4\xfc\x0d\x5c\xb0\xd0\xef\xe7\x65\xeb\xc1\x60\x21\xa1\x91\xbb\x44\x79\x41\x65\x44\x96\x25\x07\x0f\xc9\xfc\x43\x91\xb5\x42\xa8\x9e\x8e\x9f\xd9\x94\xb9\xe6\xa4\xf1\x7a\x30\xc6\x68\xf4\xf9\x41\x91\xd1\xf4\xe9\x9e\x85\x51\xa1\x2c\x5e\x64\x56\xee\x95\x0f\x67\xe8\x27\x77\x9a\x1f\xbf\xb9\x81\xb9\x7f\xa5\xb7\xae\xb9\xc2\xfd\x2b\x3b\x2d\x1d\xac\x2e\xf0\x8e\xcc\x0e\xc8\xbc\x86\xbd\x02\x1d\x73\xb7\x5b\xcc\x33\xe2\x28\xcf\x17\x6e\x6a\x4f\x1b\x92\x92\x6b\x80\x49\x7d\x5a\x74\xc6\x7b\xce\x24\xfa\x0d\xa6\xb0\xe8\x22\x5a\x19\x8d\x2f\x82\x73\xe0\x17\x60\xd9\x64\xe9\xa2\x96\x8e\xe6\xb8\xfe\xbf\x00\x00\x00\xff\xff\x1e\xa5\x12\x3c\xea\x2b\x00\x00")

func staticPagesWritingHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticPagesWritingHtml,
		"static/pages/writing.html",
	)
}

func staticPagesWritingHtml() (*asset, error) {
	bytes, err := staticPagesWritingHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/pages/writing.html", size: 11242, mode: os.FileMode(420), modTime: time.Unix(1599812853, 0)}
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
	"static/pages/writing.html": staticPagesWritingHtml,
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
	"static": &bintree{nil, map[string]*bintree{
		"pages": &bintree{nil, map[string]*bintree{
			"writing.html": &bintree{staticPagesWritingHtml, map[string]*bintree{}},
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
