package asset_manager

// raw asset represent an asset file from a specific locale

type RawAsset interface {
	GetLocale() string
	GetAssetPath() string
	GetUpdateSQL(*AssetDatabase) string // return the SQL to update the database file
	GetPack() *Pack                     // return the asset pack object that contain this RawAsset
	// note that for sound specifically, it has 2 files associated with it, so we need a secondary pack
	// Pack() return acb pack, while AwbPack return the awb pack, other assets should return nil here
	GetAwbPack() *Pack
	Info() string      // return a short string that contain enough information to fully describe the asset
	IsAvailable() bool // return if the pack can be found (i.e. not missing)

	// directly insert into the database
	InsertIntoAssetDatabase(*AssetDatabase)
}
