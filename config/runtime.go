package config

import (
	"elichika/enum"
	"elichika/utils"

	"encoding/json"
)

type RuntimeConfig struct {
	CdnServer        *string `json:"cdn_server"`
	ServerAddress    *string `json:"server_address"`
	TapBondGain      *int32  `json:"tap_bond_gain"`
	AutoJudgeType    *int32  `json:"auto_judge_type"`
	Tutorial         *bool   `json:"tutorial"`
	LoginBonusSecond *int    `json:"login_bonus_second"` // the second from mid-night till login bonus
	TimeZone         *string `json:"timezone"`           // https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
}

func defaultConfigs() *RuntimeConfig {
	configs := RuntimeConfig{
		CdnServer:        new(string), // self-hosted
		ServerAddress:    new(string),
		TapBondGain:      new(int32),
		AutoJudgeType:    new(int32),
		Tutorial:         new(bool),
		LoginBonusSecond: new(int),
		TimeZone:         new(string),
	}
	*configs.CdnServer = "https://llsifas.catfolk.party/static/"
	*configs.ServerAddress = "0.0.0.0:8080"
	*configs.TapBondGain = 20
	*configs.AutoJudgeType = enum.JudgeTypeGreat
	*configs.Tutorial = true
	*configs.LoginBonusSecond = enum.HourSecondCount * 4
	*configs.TimeZone = "Asia/Tokyo"
	return &configs
}

func Load(p string) *RuntimeConfig {
	if !utils.PathExists(p) {
		c := defaultConfigs()
		c.Save(p)
		return c
	}

	c := RuntimeConfig{}
	err := json.Unmarshal([]byte(utils.ReadAllText(p)), &c)
	if err != nil {
		panic("config file is wrong, change/delete it and try again")
	}
	d := defaultConfigs()
	// TODO: rewrite this with reflect
	if c.CdnServer == nil {
		c.CdnServer = d.CdnServer
	}
	if c.ServerAddress == nil {
		c.ServerAddress = d.ServerAddress
	}
	if c.TapBondGain == nil {
		c.TapBondGain = d.TapBondGain
	}
	if c.AutoJudgeType == nil {
		c.AutoJudgeType = d.AutoJudgeType
	}
	if c.Tutorial == nil {
		c.Tutorial = d.Tutorial
	}
	if c.LoginBonusSecond == nil {
		c.LoginBonusSecond = d.LoginBonusSecond
	}
	if c.TimeZone == nil {
		c.TimeZone = d.TimeZone
	}

	return &c
}

func (c *RuntimeConfig) Save(p string) error {
	data, err := json.Marshal(c)
	utils.CheckErr(err)
	utils.WriteAllText(p, string(data))
	return nil
}
