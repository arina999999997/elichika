package config

// with this config, we can control how resource are handled by the server
// While it's correctly implemented, only 3 settings will be exposed through the admin webui:
// - Original:
//   - The original settings, every resources are consumed when necessary.
//   - There is no special behaviour to make things easier (i.e. free gems when trying to buy it).
//   - Use this setting if you want to experience the true f2p life.
// - Comfortable:
//   - Like original but with QOL improvements.
//   - Think of this like what a whale can afford to do.
//   - Things like LP and AP and Gem are unlimited, as it can be bought.
//   - But other things earned ingame like shop currencies are limited, so you still have to play to get them.
//   - This is the default settings.
// - Free:
//   - Like comfortable but even more unrestrictive.
//   - Generally, numbers can only go up, not down.
//   - Not recommended if there are random users on the server, as the storage can really get massive if some people misbehave:
//     - For example, cheering in member guild always give 1 present box reward per magaphone
//     - With comfortable setting, they are rate limited by the training speed, because megaphone are still consumed.
//     - With free setting, the megaphone isn't consumed, so they can build up 1000s of megaphone and just send the same thing over and over again.

type ResourceConfig struct {
	// the field are listed in roughly in the order on how free it is:
	// - free mode set everything to false.
	// - comfortable mode set a prefix to false.
	// - original mode set everything to true
	// the name is more related to the menu name, not the item itself:
	// - some items used in multiple contextes like Gold can be consumed in once place but not the others.
	ConsumeLp             bool
	ConsumeAp             bool
	ConsumeGachaCurrency  bool
	ConsumeNaviTap        bool
	ConsumeDailyLiveLimit bool

	ConsumeSkipTicket       bool
	ConsumePracticeItems    bool // practice items, grade up items, accessory improvement items
	ConsumeExchangeCurrency bool // whether currency are consumed WHILE EXCHANGE at the exchange shop
	ConsumeMemberCheerItem  bool // whether to consume the cheer item when cherring. the free cheer is always consumed.
	ConsumeMiscItems        bool // whether to consume random items like memory key or water bottle.
}

var resourceConfigs = map[string]*ResourceConfig{}

func init() {
	resourceConfigs["free"] = &ResourceConfig{}
	resourceConfigs["comfortable"] = &ResourceConfig{
		ConsumeSkipTicket:       true,
		ConsumePracticeItems:    true,
		ConsumeExchangeCurrency: true,
		ConsumeMemberCheerItem:  true,
		ConsumeMiscItems:        true,
	}
	resourceConfigs["original"] = &ResourceConfig{
		ConsumeLp:             true,
		ConsumeAp:             true,
		ConsumeGachaCurrency:  true,
		ConsumeNaviTap:        true,
		ConsumeDailyLiveLimit: true,

		ConsumeSkipTicket:       true,
		ConsumePracticeItems:    true,
		ConsumeExchangeCurrency: true,
		ConsumeMemberCheerItem:  true,
		ConsumeMiscItems:        true,
	}
}
