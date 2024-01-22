package generic

type UserIdWrapper[T any] struct {
	UserId int32 `xorm:"pk 'user_id'"`
	Object *T    `xorm:"extends"`
}
