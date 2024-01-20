package enum

const (
	NoteGimmickTypeSuccess                     int32 = 0x00000001
	NoteGimmickTypeFailure                     int32 = 0x00000002
	NoteGimmickTypeAlways                      int32 = 0x00000003
	NoteGimmickTypeSuccessAndCardRoleIsBoost   int32 = 0x00000004
	NoteGimmickTypeSuccessAndCardRoleIsCharge  int32 = 0x00000005
	NoteGimmickTypeSuccessAndCardRoleIsHeal    int32 = 0x00000006
	NoteGimmickTypeSuccessAndCardRoleIsSupport int32 = 0x00000007
	NoteGimmickTypeFailureOrCardRoleIsBoost    int32 = 0x00000008
	NoteGimmickTypeFailureOrCardRoleIsCharge   int32 = 0x00000009
	NoteGimmickTypeFailureOrCardRoleIsHeal     int32 = 0x0000000a
	NoteGimmickTypeFailureOrCardRoleIsSupport  int32 = 0x0000000b
)
