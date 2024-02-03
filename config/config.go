package config

import (
	"elichika/utils"

	"os"
	"time"

	_ "modernc.org/sqlite"
)

var (
	// merged from:
	// https://github.com/RayFirefist/SukuStar_Datamine/blob/master/lib/sifas_api/sifas.py#L120
	// https://github.com/RayFirefist/SukuStar_Datamine/blob/master/lib/sifas_api/sifas.py#L400
	ServerEventReceiverKey = "4924c4421e9e3a287dc31e2ff241a8fb46389c7f30bafee791bb06c9ae3b6c82"

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
