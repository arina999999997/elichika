package user_social

import (
	"elichika/client"
	"elichika/subsystem/cache"
	"elichika/userdata"
	"elichika/utils"
)

var (
	getLivePartnerCache = cache.UniquePointerMap[int64, cache.CachedObject[client.LivePartner]]{}
)

func GetLivePartner(session *userdata.Session, otherUserId int32) client.LivePartner {
	key := (int64(otherUserId) << 32)
	cacher := getLivePartnerCache.Get(key)
	cacher.Acquire()
	defer cacher.Release()
	if cacher.ExpireAt <= session.Time.Unix() {
		cacher.ExpireAt = session.Time.Unix() + 60
		cacher.Value = getLivePartnerNoCache(session, otherUserId)
	}
	return *cacher.Value
}

func getLivePartnerNoCache(session *userdata.Session, otherUserId int32) *client.LivePartner {
	partner := client.LivePartner{}

	exist, err := session.Db.Table("u_status").Where("user_id = ?", otherUserId).Get(&partner)
	utils.CheckErrMustExist(err, exist)

	partner.IsFriend = IsFriend(session, otherUserId)
	for i := int32(1); i <= 7; i++ {
		partnerCard := GetOtherUserProfileLivePartnerCard(session, otherUserId, i)
		if partnerCard.PartnerCard.CardMasterId != 0 {
			partner.CardByCategory.Set(i, partnerCard.PartnerCard)
		}
	}
	return &partner
}
