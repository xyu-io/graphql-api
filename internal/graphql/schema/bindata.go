// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../module/demo/graphql/demo_sdl/demo.graphql
// ../sdl/demo.graphql
// ../sdl/schema.graphql
package schema

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

var _ModuleDemoGraphqlDemo_sdlDemoGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x70\x49\xcd\xcd\x57\x28\xc9\x2c\xc9\x49\x55\x78\x36\x7f\xe9\x8b\xf5\x8b\x9e\xf6\x37\x3d\x9b\xba\x81\x97\x2b\x33\xaf\xa0\xb4\x04\x2c\x1d\x58\x9a\x5a\x54\x19\x52\x59\x90\xaa\x50\xcd\xcb\xa5\xa0\xa0\xa0\xa0\xac\xf0\x6c\x41\xfb\xcb\x45\x33\x20\xbc\xcc\x14\x2b\x85\xe0\x92\xa2\xcc\xbc\x74\x45\x5e\xae\x5a\x5e\x2e\x5e\x2e\x14\x53\x9f\xec\x5f\xf7\x6c\xca\x4e\x4c\x53\x7d\x4b\x4b\x12\x4b\x32\xf3\xf3\x70\x1b\x0c\xd6\x0f\x33\x1b\x9b\xd1\x2f\xf6\x4f\x79\x3a\x7b\xde\xd3\x86\x3d\xcf\x37\xee\x7e\x3a\xaf\x9b\x97\xab\x04\x64\x16\x48\x41\x50\x6a\x71\x41\x7e\x5e\x71\x2a\x91\x86\x83\x1c\x0e\x08\x00\x00\xff\xff\x3f\xa3\x97\xe4\x0b\x01\x00\x00")

func ModuleDemoGraphqlDemo_sdlDemoGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_ModuleDemoGraphqlDemo_sdlDemoGraphql,
		"../../../module/demo/graphql/demo_sdl/demo.graphql",
	)
}

func ModuleDemoGraphqlDemo_sdlDemoGraphql() (*asset, error) {
	bytes, err := ModuleDemoGraphqlDemo_sdlDemoGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../../module/demo/graphql/demo_sdl/demo.graphql", size: 267, mode: os.FileMode(438), modTime: time.Unix(1707201520, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _SdlDemoGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x70\x49\xcd\xcd\x0f\x2c\x4d\x2d\xaa\xf4\x4f\xca\x4a\x4d\x2e\x09\xa9\x2c\x48\xad\xe6\xe5\x52\x50\x50\x50\x48\x4f\x2d\x09\xc9\x2c\xc9\x49\xb5\x02\xab\x09\x4a\x2d\x2e\xc8\xcf\x2b\x4e\x05\x29\xe0\xe5\xaa\xe5\xe5\xe2\xe5\x82\xeb\xf7\x2d\x2d\x49\x2c\xc9\xcc\xcf\xc3\x30\x22\x39\x23\x31\x2f\x3d\x15\x6c\x8a\x46\x09\xc4\xac\xe0\x92\xa2\xcc\xbc\x74\x45\x4d\x1c\xa6\x02\x02\x00\x00\xff\xff\x1e\x1f\x15\xc4\x94\x00\x00\x00")

func SdlDemoGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_SdlDemoGraphql,
		"../sdl/demo.graphql",
	)
}

func SdlDemoGraphql() (*asset, error) {
	bytes, err := SdlDemoGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../sdl/demo.graphql", size: 148, mode: os.FileMode(438), modTime: time.Unix(1707201500, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _SdlSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\xe5\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x10\x91\xdc\xd2\x92\xc4\x92\xcc\xfc\x3c\x2b\x05\x5f\x28\x8b\x97\xab\x96\x97\x4b\x59\x21\xbd\x28\xb1\x20\xa3\x30\xe7\x59\x67\xc3\xb3\x39\x9d\xcf\xe6\x2f\x7d\xb1\x7e\xd1\xb3\xbe\xa5\x4f\xfb\x17\xf3\x72\x95\x54\x16\xa4\x42\xcc\x80\x99\x9b\x92\x9a\x9b\x6f\xa5\xe0\x92\x9a\x9b\x0f\x16\xf6\x4f\xca\x4a\x4d\x2e\x09\xa9\x2c\x48\xc5\x66\xd8\x93\xfd\xeb\x9e\x4d\xd9\x89\x62\x18\xcc\x6e\x4c\xf3\x60\x32\x68\x46\xf2\x72\x01\x02\x00\x00\xff\xff\xc4\x80\x3c\x7a\xd9\x00\x00\x00")

func SdlSchemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_SdlSchemaGraphql,
		"../sdl/schema.graphql",
	)
}

func SdlSchemaGraphql() (*asset, error) {
	bytes, err := SdlSchemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../sdl/schema.graphql", size: 217, mode: os.FileMode(438), modTime: time.Unix(1707199749, 0)}
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
	"../../../module/demo/graphql/demo_sdl/demo.graphql": ModuleDemoGraphqlDemo_sdlDemoGraphql,
	"../sdl/demo.graphql":                                SdlDemoGraphql,
	"../sdl/schema.graphql":                              SdlSchemaGraphql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"..": &bintree{nil, map[string]*bintree{
				"module": &bintree{nil, map[string]*bintree{
					"demo": &bintree{nil, map[string]*bintree{
						"graphql": &bintree{nil, map[string]*bintree{
							"demo_sdl": &bintree{nil, map[string]*bintree{
								"demo.graphql": &bintree{ModuleDemoGraphqlDemo_sdlDemoGraphql, map[string]*bintree{}},
							}},
						}},
					}},
				}},
			}},
		}},
		"sdl": &bintree{nil, map[string]*bintree{
			"demo.graphql":   &bintree{SdlDemoGraphql, map[string]*bintree{}},
			"schema.graphql": &bintree{SdlSchemaGraphql, map[string]*bintree{}},
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
