package config

import (
	"elichika/utils"
	"encoding/json"
	"os"
	"strconv"
	"time"
)

type AppConfigs struct {
	CdnServer  string   `json:"cdn_server"`
	ServerAddress string `json:"server_address"`
	TapBondGain int `json:"tap_bond_gain"`
}

func DefaultConfigs() *AppConfigs {
	return &AppConfigs{
		CdnServer: "http://127.0.0.1:8080/static", // self-hosted 
		ServerAddress: "0.0.0.0:8080",
		TapBondGain: 20,
	}
}

func Load(p string) *AppConfigs {
	if !utils.PathExists(p) {
		_ = DefaultConfigs().Save(p)
	}
	c := AppConfigs{}
	err := json.Unmarshal([]byte(utils.ReadAllText(p)), &c)
	if err != nil {
		_ = os.Rename(p, p+".backup"+strconv.FormatInt(time.Now().Unix(), 10))
		_ = DefaultConfigs().Save(p)
	}
	c = AppConfigs{}
	_ = json.Unmarshal([]byte(utils.ReadAllText(p)), &c)
	return &c
}

func (c *AppConfigs) Save(p string) error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	utils.WriteAllText(p, string(data))
	return nil
}
