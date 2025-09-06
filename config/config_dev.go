//go:build dev

package config

const (
	// Running in developer mode will likely corrupt the database in one way or another, back them up before doing anything
	// DeveloperMode = 0: off
	// DeveloperMode = 1: developing marathon event
	// DeveloperMode > 1: reserved for future
	// The developer mode functionality will be developped as need arise, and will not exists on the final product
	// Note that even with build tag, developer mode still need to be set to differentiate the use cases
	DeveloperMode                 = DeveloperModeEventMarathonDev
	DeveloperModeEventMarathonDev = 1
)
