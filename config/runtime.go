package config

import (
	"elichika/enum"
	"elichika/utils"

	"encoding/json"
	"fmt"
	"reflect"
)

type RuntimeConfig struct {
	CdnServer            *string `json:"cdn_server"`
	ServerAddress        *string `json:"server_address"`
	TapBondGain          *int32  `json:"tap_bond_gain"`
	AutoJudgeType        *int32  `json:"auto_judge_type"`
	Tutorial             *bool   `json:"tutorial"`               // whether to turn on tutorial when starting a new account
	LoginBonusSecond     *int    `json:"login_bonus_second"`     // the second from mid-night till login bonus
	TimeZone             *string `json:"timezone"`               // https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	DefaultContentAmount *int32  `json:"default_content_amount"` // the amount of items to give an user if they don't have that item
	MissionMultiplier    *int32  `json:"mission_multiplier"`     // multiply the progress of missions. Only work for do "x" of things, not for "get x different thing or reach x level"
}

func defaultConfigs() *RuntimeConfig {
	configs := RuntimeConfig{
		CdnServer:            new(string), // self-hosted
		ServerAddress:        new(string),
		TapBondGain:          new(int32),
		AutoJudgeType:        new(int32),
		Tutorial:             new(bool),
		LoginBonusSecond:     new(int),
		TimeZone:             new(string),
		DefaultContentAmount: new(int32),
		MissionMultiplier:    new(int32),
	}
	*configs.CdnServer = "https://llsifas.catfolk.party/static/"
	*configs.ServerAddress = "0.0.0.0:8080"
	*configs.TapBondGain = 20
	*configs.AutoJudgeType = enum.JudgeTypeGreat
	*configs.Tutorial = true
	*configs.LoginBonusSecond = enum.HourSecondCount * 4
	*configs.TimeZone = "Asia/Tokyo"
	*configs.DefaultContentAmount = 10000000
	*configs.MissionMultiplier = 1
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
	for i := 0; i < reflect.TypeOf(c).NumField(); i++ {
		f := reflect.ValueOf(&c).Elem().Field(i)
		if f.IsNil() {
			fmt.Println("Use default setting: ", reflect.TypeOf(c).Field(i).Name)
			f.Set(reflect.ValueOf(d).Elem().Field(i))
		}
		fmt.Println(reflect.TypeOf(c).Field(i).Name, ": ", f.Elem())
	}
	return &c
}

func (c *RuntimeConfig) Save(p string) error {
	data, err := json.Marshal(c)
	utils.CheckErr(err)
	utils.WriteAllText(p, string(data))
	return nil
}
