package gamedata

import (
	"elichika/dictionary"
	"fmt"

	"xorm.io/xorm"
)

// as we implement more features, some asset pack break the original server's (and the public cdn) rule.
// therefore, we need to handle asset pack on a file-to-file basis.
// ideally, this is a part of the server and we will read from some server-sided database that automatically track the actual change
// but for now, it's a map for special package
// TOOD(assets): actually make this a database instead of hard coding 1 instance
type AssetPack struct {
	Name          string `xorm:"pk 'name'"`
	MasterVersion string `xorm:"'master_version'"`
}

func loadAssetPack(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading AssetPack")
	gamedata.AssetPack = make(map[string]*AssetPack)
	gamedata.AssetPack["ubias7"] = &AssetPack{
		Name:          "ubias7",
		MasterVersion: "b66ec2295e9a00aa",
	}
}

func init() {
	addLoadFunc(loadAssetPack)
}
