package client

type DoRarityUp struct {
	BeforeAccessoryRarity int32 `json:"before_accessory_rarity" enum:"AccessoryRarity"`
	AfterAccessoryRarity  int32 `json:"after_accessory_rarity" enum:"AccessoryRarity"`
	DoRarityUpAddSkill    bool  `json:"do_rarity_up_add_skill"`
}
