package config

import (
	"elichika/enum"
	"elichika/utils"

	"encoding/json"
	"fmt"
	"reflect"
)

type RuntimeConfig struct {
	ServerAddress            *string `json:"server_address" of_label:"Server's address"`
	CdnServer                *string `json:"cdn_server" of_label:"CDN server's address"`
	CdnPartialFileCapability *string `json:"cdn_partial_file_capability" of_type:"select" of_options:"Has static partial file\nstatic_file\nHas partial file mapping\nmapped_file\nHas range API\nhas_range_api\nNothing\nnothing" of_label:"CDN server's partial file capability" `
	AdminPassword            *string `json:"admin_password" of_label:"Admin password" of_type:"password""`
	TapBondGain              *int32  `json:"tap_bond_gain" of_label:"Partner tap bond reward" of_attrs:"min=\"0\" max=\"20000000\"`
	AutoJudgeType            *int32  `json:"auto_judge_type" of_type:"select" of_options:"None\n1\nMiss\n10\nBad\n12\nGood\n14\nGreat\n20\nPerfect\n30" of_label:"Autoplay judge type"`
	Tutorial                 *bool   `json:"tutorial" of_label:"Enable tutorial"`                                                          // whether to turn on tutorial when starting a new account
	LoginBonusSecond         *int32  `json:"login_bonus_second" of_type:"time" of_label:"Login bonus reset time"`                          // the second from mid-night till login bonus
	TimeZone                 *string `json:"timezone" of_label:"Timezone (from tz database)"`                                              // https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	DefaultContentAmount     *int32  `json:"default_content_amount" of_label:"Default item count" of_attrs:"min=\"0\" max=\"1000000000\""` // the amount of items to give an user if they don't have that item
	MissionMultiplier        *int32  `json:"mission_multiplier" of_label:"Mission progress multiplier" of_attrs:"min=\"0\" max=\"10000\""` // multiply the progress of missions. Only work for do "x" of things, not for "get x different thing or reach x level"
	ResourceConfigType       *string `json:"resource_config_type" of_type:"select" of_options:"Original\noriginal\nComfortable\ncomfortable\nFree\nfree" of_label:"Resource config"`
	EventAutoSchedulerPeriod *string `json:"event_auto_scheduler_period" of_type:"select" of_options:"once per day\n1_day\nonce per week\n7_days" of_label:"Event frequency"`
	MaintenanceUrl           *string `json:"maintenance_url" of_label:"Maintenance Url"`
}

func defaultConfigs() *RuntimeConfig {
	configs := RuntimeConfig{
		ServerAddress:            new(string),
		CdnServer:                new(string),
		CdnPartialFileCapability: new(string),
		AdminPassword:            new(string),
		TapBondGain:              new(int32),
		AutoJudgeType:            new(int32),
		Tutorial:                 new(bool),
		LoginBonusSecond:         new(int32),
		TimeZone:                 new(string),
		DefaultContentAmount:     new(int32),
		MissionMultiplier:        new(int32),
		ResourceConfigType:       new(string),
		EventAutoSchedulerPeriod: new(string),
		MaintenanceUrl:           new(string),
	}
	*configs.ServerAddress = "0.0.0.0:8080"
	*configs.CdnServer = "https://llsifas.imsofucking.gay/static/"
	*configs.CdnPartialFileCapability = "nothing"
	*configs.AdminPassword = ""
	*configs.TapBondGain = 20
	*configs.AutoJudgeType = enum.JudgeTypeGreat
	*configs.Tutorial = true
	*configs.LoginBonusSecond = enum.HourSecondCount * 4
	*configs.TimeZone = "Asia/Tokyo"
	*configs.DefaultContentAmount = 10000000
	*configs.MissionMultiplier = 1
	*configs.ResourceConfigType = "comfortable"
	*configs.EventAutoSchedulerPeriod = "7_days"
	*configs.MaintenanceUrl = "http://127.0.0.1:8080/webui/admin/"
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
	if *c.CdnServer == "https://llsifas.catfolk.party/static/" {
		*c.CdnServer = "https://llsifas.imsofucking.gay/static/"
	}
	return &c
}

func (c *RuntimeConfig) Save(p string) error {
	data, err := json.Marshal(c)
	utils.CheckErr(err)
	utils.WriteAllText(p, string(data))
	return nil
}

var confs = []*RuntimeConfig{}

func UpdateConfig(newConfig *RuntimeConfig) {
	confs = append(confs, Conf)
	newConfig.Save("./config.json") // this has lock so the file can't be corrupted
	// this should be safe because we overwrite the pointer, not the object
	// if someone had an old version of confs then they would just have an access to the old config until they discard it
	// this also assume the pointer size is less than machine word, which it should be
	Conf = newConfig
}

func (r RuntimeConfig) ResourceConfig() *ResourceConfig {
	return resourceConfigs[*r.ResourceConfigType]
}
