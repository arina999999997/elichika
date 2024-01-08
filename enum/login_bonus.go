package enum

const (
	LoginBonusReceiveStatusReceived   int32 = 0x00000001
	LoginBonusReceiveStatusReceiving  int32 = 0x00000002
	LoginBonusReceiveStatusUnreceived int32 = 0x00000003
)

const (
	LoginBonusContentGradeNormal int32 = 0x00000001
	LoginBonusContentGradeRare   int32 = 0x00000002
)

const (
	LoginBonusTypeNormal   int32 = 0x00000001
	LoginBonusTypeBeginner int32 = 0x00000002
	LoginBonusTypeEvent2d  int32 = 0x00000003
	LoginBonusTypeEvent3d  int32 = 0x00000004
	LoginBonusTypeBirthday int32 = 0x00000005
	LoginBonusTypeComeback int32 = 0x00000006
)
