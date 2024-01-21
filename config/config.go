package config

import (
	"elichika/utils"

	"os"
	"time"

	_ "modernc.org/sqlite"
)

var (
	// Taken from:
	// https://github.com/RayFirefist/SukuStar_Datamine/blob/master/lib/sifas_api/sifas.py#L120
	// https://github.com/RayFirefist/SukuStar_Datamine/blob/master/lib/sifas_api/sifas.py#L400
	ServerEventReceiverKey = "31f1f9dc7ac4392d1de26acf99d970e425b63335b461e720c73d6914020d6014"
	JaKey                  = "78d53d9e645a0305602174e06b98d81f638eaf4a84db19c756866fddac360c96"
	// TODO: can merge JaKey

	SessionKey = "12345678123456781234567812345678"

	AssetPath = "assets/"

	GlMasterdataPath = AssetPath + "db/gl/"
	JpMasterdataPath = AssetPath + "db/jp/"

	ServerdataPath = "serverdata.db"
	UserdataPath   = "userdata.db"

	UserDataBackupPath = "backup/"

	MasterVersionGl = "2d61e7b4e89961c7" // read from GL database, so user can update db just by changing that
	MasterVersionJp = "b66ec2295e9a00aa" // ditto

	GlStartupKey = "TxQFwgNcKDlesb93"
	JpStartupKey = "5f7IZY1QrAX0D49g"

	StaticDataPath = "static/"

	Platforms   = []string{"a", "i"}
	GlLanguages = []string{"en", "ko", "zh"}
	JpLanguages = []string{"ja"}

	ServerInitJsons = "server init jsons/"
	Conf            = &RuntimeConfig{}

	GenerateStageFromScratch = false
)

func init() {
	os.MkdirAll(UserDataBackupPath, 0755)
	Conf = Load("./config.json")
	loc, err := time.LoadLocation(*Conf.TimeZone)
	utils.CheckErr(err)
	time.Local = loc
}
