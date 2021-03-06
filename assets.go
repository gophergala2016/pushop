// Code generated by go-bindata.
// sources:
// assets/templates/image.html
// assets/templates/index.html
// assets/templates/layout.html
// DO NOT EDIT!

package main

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

var _assetsTemplatesImageHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x91\xcf\x8e\x9b\x30\x18\xc4\xcf\xe4\x29\x2c\x1f\x7a\x4b\xac\x36\xb7\x16\x90\x5a\xf5\x52\xa9\x4a\xab\x95\xf6\x01\x8c\xf9\x02\x4e\xf0\x1f\xd9\x26\xbb\xc8\xe2\xdd\xd7\x18\x92\x98\xec\xde\xf6\x00\x8c\x3c\xe3\x9f\xc7\x1f\xde\xd7\x70\xe4\x12\x10\x66\x4a\x3a\x90\x0e\x8f\xe3\x26\xcb\x6b\x7e\x41\xac\xa3\xd6\x16\xd8\xa8\x17\x5c\x6e\xb2\xd5\x1a\x53\xdd\xd6\x8a\xed\xd7\x6f\x68\x52\xa2\x0e\x2a\x66\x56\x21\xd7\xf6\xa2\x92\x94\x77\x18\x71\x07\xc2\x32\xa5\x21\x2a\x37\x68\x28\x70\xeb\x9c\xfe\x4e\x88\x65\x2d\x08\xba\x53\xa6\x21\x7f\x04\x6d\xe0\x5f\x75\x02\xe6\x66\x5a\x96\x73\xd1\x20\x6b\x58\x81\xbd\xdf\x45\x7b\x7e\x3f\x3f\xfd\x1d\x47\x8c\x68\xe7\x12\xe7\x40\x05\x84\xd5\x65\x67\xda\x96\x6a\xc7\x95\x5c\x9c\x2c\x6f\xf7\xb1\x86\x36\x4a\x17\x58\x86\x5d\xb8\x7c\x80\xe4\xa4\xdd\x2f\x69\xef\x11\x3f\xa2\xc5\xfd\x0d\x96\x19\x1e\x69\x21\xa4\x13\x4c\x7d\x77\x12\xda\x3a\x4f\x74\x70\x40\xd6\xd3\x84\x13\xf4\xcf\xde\xb5\xca\x4c\xbc\xf2\xd7\x80\x72\xab\xa9\x4c\xc0\x34\xba\x91\x79\x0b\x92\x29\x53\x3e\xf2\x72\x12\xee\x3c\xb7\xf6\xfe\xde\xf9\xbf\x81\x0b\x57\xbd\x0d\xfb\xe8\x75\x22\x95\x93\x28\x3c\x5b\x6d\xb8\xa0\x66\x88\xba\x6b\x30\x32\xaa\x0b\x7f\xa6\xea\x9d\x0b\xd7\x40\x94\x31\xb0\xf6\x0c\x43\x81\x4f\x18\xb5\x06\x8e\xc9\xb4\xef\x58\x5c\x7e\xe9\xa8\x31\x3f\x72\x42\x57\x7d\xd2\x12\x07\x78\x75\x9f\x2a\x70\x7e\x57\x60\x46\x86\xc3\xcd\x07\x87\xdf\x66\x71\x15\xcb\xf7\x1a\x79\x0b\x00\x00\xff\xff\xc3\xfe\x95\xa5\xf7\x02\x00\x00")

func assetsTemplatesImageHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesImageHtml,
		"assets/templates/image.html",
	)
}

func assetsTemplatesImageHtml() (*asset, error) {
	bytes, err := assetsTemplatesImageHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/image.html", size: 759, mode: os.FileMode(436), modTime: time.Unix(1453647100, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8e\xcf\x4e\xc4\x20\x10\xc6\xcf\xdd\xa7\x98\x10\x8f\x76\x7b\xd0\x78\x30\x94\xb3\x26\xc6\x83\xd1\x07\x18\x81\xdd\x92\xf0\xc7\x00\xea\x61\xc2\xbb\x0b\x4b\x9b\x74\x3d\x34\xf9\x4d\xbf\xef\x37\x0c\x91\xd2\x27\xe3\x35\x30\x19\x7c\xd6\x3e\xb3\x52\x0e\x03\x57\xe6\x07\xa4\xc5\x94\x66\x16\xc3\x2f\x13\x44\x11\xfd\x59\xc3\x8d\xb9\xad\x9f\xc3\x8a\x8f\x33\x1c\x9f\x1b\xa5\x66\x5c\x29\x32\xd8\x31\xb9\xf1\x01\x1a\x38\x35\xde\x33\x51\x1b\x57\x95\xbc\x7c\xbb\x4f\x8f\xc6\xf6\x68\xe0\x08\x4b\xd4\xa7\x99\x11\xf5\xfd\xc7\xa7\xec\xec\xc7\xdb\x4b\x29\x4c\x70\xe3\xce\x90\xa2\xdc\xa5\xef\x9b\xdf\x2b\x80\x36\xef\xd2\x57\x74\xfa\x22\x4e\xb8\xae\xdf\x5f\x87\x5f\xd9\x04\xbf\x3e\x3c\xf0\xe5\x4e\xfc\x13\xf9\x54\xff\x75\x6f\xaa\x62\x3f\x7e\xa3\x0d\x88\xb4\x57\xad\x7b\x19\xd7\xe9\xf0\x17\x00\x00\xff\xff\x2c\x8c\x7b\xa7\x50\x01\x00\x00")

func assetsTemplatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesIndexHtml,
		"assets/templates/index.html",
	)
}

func assetsTemplatesIndexHtml() (*asset, error) {
	bytes, err := assetsTemplatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/index.html", size: 336, mode: os.FileMode(436), modTime: time.Unix(1453646843, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesLayoutHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x54\x4d\x6f\xe3\x36\x10\x3d\x67\x7f\x05\xab\x8b\xec\x56\xa6\xd6\x4d\x3f\x8c\x5a\x0a\x90\x85\xd3\x05\x8a\x76\x37\x40\x83\x0d\xd0\x20\x87\x31\x35\x92\xe8\x50\xa4\x96\xa4\xa2\xb8\x8e\xfe\x7b\x87\x52\x62\xef\xf6\xb0\x17\x91\x9c\x79\xf3\xc8\x37\x1f\xca\xbe\x2b\x8c\xf0\xfb\x16\x59\xed\x1b\x75\xf1\x26\x7b\x5d\x10\x8a\x8b\x37\x67\x59\x83\x1e\x98\xa8\xc1\x3a\xf4\x79\xd4\xf9\x72\xb1\x8a\x8e\x76\x0d\x0d\xe6\xd1\xa3\xc4\xbe\x35\xd6\x47\x4c\x18\xed\x51\x13\xae\x97\x85\xaf\xf3\x02\x1f\xa5\xc0\xc5\x78\x48\x98\xd4\xd2\x4b\x50\x0b\x27\x40\x61\xbe\xe4\x6f\x47\x1e\x2f\xbd\xc2\x8b\xc3\x81\xdf\x84\xcd\x30\x64\xe9\x64\xf9\xfa\x8a\x02\x9d\xb0\xb2\xf5\xd2\xe8\x2f\x6e\xa1\xa8\xcd\xc9\x31\x0c\x23\xa1\x92\xfa\x81\x59\x54\x79\xe4\xfc\x5e\xa1\xab\x11\xe9\x65\xb5\xc5\x32\x8f\x6a\xef\x5b\xf7\x5b\x9a\x36\xf0\x24\x0a\xcd\xb7\xc6\x78\xe7\x2d\xb4\xe1\x20\x4c\x93\x1e\x0d\xe9\x39\x3f\xe7\xbf\xa4\xc2\xb9\x93\x8d\x37\x92\x50\xce\x45\xa4\xc4\x63\x65\xa5\xdf\xd3\x1d\x35\x9c\xaf\x7e\x5a\x2c\x3f\xaf\x9a\x9b\x3f\x3e\x5e\xfe\xfd\xb4\xda\x2d\x2f\xbb\x1f\xe0\xe7\xdb\xcd\x27\x7d\x2d\x7f\x54\x0f\xbf\x97\x7d\x7f\x75\x09\xab\x7a\xb3\x29\x76\xff\xa8\xf6\x4f\xac\x9e\xea\xdd\xa7\xbf\xae\x96\x65\xb5\xbb\xbd\x7e\xdf\x3c\xfc\xeb\x7e\x25\x51\xd6\x38\x67\xac\xac\xa4\xce\x23\xd0\x46\xef\x1b\xd3\x39\x52\x94\xa5\x53\x29\xb2\xad\x29\xf6\x4c\x7a\x6c\x9c\x30\x2d\xe5\x24\x1a\x0f\xa1\x74\x93\x30\xd2\xe5\x44\x8d\x0d\x70\x63\xab\xf4\x16\xb7\xd7\x50\x61\x48\xc9\xe1\xc0\x64\xc9\xf8\xa5\x06\xb5\xf7\x52\xb8\x61\xa0\x34\x4d\x69\x23\xef\xd9\xac\xec\xb4\x08\x09\x9c\xc9\xc4\x25\x26\xa9\x12\x9b\x40\xd2\xcc\x0f\xf2\x2e\x7e\x6f\x4c\xa5\xf0\x18\xf9\x71\xbb\x43\xe1\xe3\xfb\xdc\xae\xe5\x9d\xbd\xcf\xc3\xe7\xf9\xf9\x18\x3f\x3f\x04\xba\x60\xe4\x9f\xf3\x69\x79\x7e\xbe\xbb\x9f\xf3\xb6\x73\xf5\x0c\x6c\xd5\x35\x54\x37\x37\x1f\x92\xd1\xa9\xf2\xe5\xf7\x1a\x7b\xb6\x01\x8f\xb3\xf9\x1a\x72\xc7\x85\x45\x3a\x5c\x29\x0c\xc0\x99\x99\x27\x44\xd8\x90\xbd\x42\xff\x62\x74\xef\xf6\x37\x50\x7d\xa0\xae\x20\xf7\xdd\xdb\xfb\x35\x70\x70\x7b\x2d\xf2\x25\xed\x9c\x15\x79\xb5\x6e\x78\x0b\x96\xa0\x1f\x4c\x81\x5c\x6a\x87\xd6\xbf\xc3\xd2\x58\x9c\x05\x59\xc4\x38\xcc\x67\xbd\xd4\x85\xe9\x13\xea\xfd\xf1\x4d\x49\x3c\xe5\x23\x4e\xe2\x34\xed\xfb\x9e\x57\xa3\xf0\x05\xbc\x2a\x1f\xdb\xe3\x74\xda\x39\x42\x56\x10\xcf\xd7\x44\x57\xc1\x2c\x9e\x1e\x1e\x27\x2c\xa6\xa6\xfc\x22\xd3\xc1\x02\x9d\x37\x27\xa4\x43\x5d\x04\x6b\x4b\xd5\x09\x93\x33\x7a\xb2\xf4\x58\x0f\xaa\x16\x21\xd8\x58\x24\x0d\x8f\x4c\x28\x70\x2e\x8f\x68\xbb\x05\xcb\xa6\x65\x51\x60\x09\x9d\xf2\xa1\xba\x67\x59\x21\x8f\xa8\x30\x1a\x20\x35\xda\x45\xa9\x3a\x59\x8c\xfe\xaf\x00\x2f\xf1\xa1\xa5\xd0\x4e\xee\xb3\x0c\xfe\xe7\xdd\x5a\xd0\xc5\xeb\xcc\xa4\x11\x4d\x27\x9b\xc6\x93\x85\xf9\x84\x89\x34\x25\xd6\xf1\xfa\x97\x4d\x96\x52\x74\x58\xbf\xf9\x9c\xc3\x81\x7a\xb6\x55\x94\x2b\x16\xbd\xcc\x71\xc4\xf8\x28\x76\xe2\xc9\xd2\xd0\xe7\x63\xdb\x8f\x3f\xa2\xff\x02\x00\x00\xff\xff\xf6\xa4\x90\xfa\xa0\x04\x00\x00")

func assetsTemplatesLayoutHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesLayoutHtml,
		"assets/templates/layout.html",
	)
}

func assetsTemplatesLayoutHtml() (*asset, error) {
	bytes, err := assetsTemplatesLayoutHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/layout.html", size: 1184, mode: os.FileMode(436), modTime: time.Unix(1453646997, 0)}
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
	"assets/templates/image.html": assetsTemplatesImageHtml,
	"assets/templates/index.html": assetsTemplatesIndexHtml,
	"assets/templates/layout.html": assetsTemplatesLayoutHtml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"image.html": &bintree{assetsTemplatesImageHtml, map[string]*bintree{}},
			"index.html": &bintree{assetsTemplatesIndexHtml, map[string]*bintree{}},
			"layout.html": &bintree{assetsTemplatesLayoutHtml, map[string]*bintree{}},
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

