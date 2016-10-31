
# json
    import "github.com/coralproject/shelf/cmd/corald/fixtures/json"






## func Asset
``` go
func Asset(name string) ([]byte, error)
```
Asset loads and returns the asset for the given name.
It returns an error if the asset could not be found or
could not be loaded.


## func AssetDir
``` go
func AssetDir(name string) ([]string, error)
```
AssetDir returns the file names below a certain
directory embedded in the file by go-bindata.
For example if you run go-bindata on data/... and data contains the
following hierarchy:


	data/
	  foo.txt
	  img/
	    a.png
	    b.png

then AssetDir("data") would return []string{"foo.txt", "img"}
AssetDir("data/img") would return []string{"a.png", "b.png"}
AssetDir("foo.txt") and AssetDir("notexist") would return an error
AssetDir("") will return []string{"data"}.


## func AssetInfo
``` go
func AssetInfo(name string) (os.FileInfo, error)
```
AssetInfo loads and returns the asset info for the given name.
It returns an error if the asset could not be found or
could not be loaded.


## func AssetNames
``` go
func AssetNames() []string
```
AssetNames returns the names of the assets.


## func MustAsset
``` go
func MustAsset(name string) []byte
```
MustAsset is like Asset but panics when Asset would return an error.
It simplifies safe initialization of global variables.


## func RestoreAsset
``` go
func RestoreAsset(dir, name string) error
```
RestoreAsset restores an asset under the given directory


## func RestoreAssets
``` go
func RestoreAssets(dir, name string) error
```
RestoreAssets restores an asset under the given directory recursively









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)