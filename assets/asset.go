package assets

import rice "github.com/GeertJohan/go.rice"

var AssetBox *rice.Box

// LoadAssets load asset box
func LoadAssets() {
	AssetBox = rice.MustFindBox("./assets")
	// AssetBox = rice.MustFindBox("assets/")
}
