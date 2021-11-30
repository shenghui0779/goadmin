package assets

import (
	"os"

	rice "github.com/GeertJohan/go.rice"
)

var AssetBox *rice.Box

// LoadAssets load asset box
func LoadAssets() {
	path := os.Getenv("aaa")
	// AssetBox = rice.MustFindBox("./assets")
	AssetBox = rice.MustFindBox(path)
}
