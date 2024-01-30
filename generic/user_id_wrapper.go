package generic

type UserIdWrapper[T any] struct {
	UserId int32 `xorm:"pk 'user_id'"`
	Object *T    `xorm:"extends"`
}

type NonPkUserIdWrapper[T any] struct {
	UserId int32 `xorm:"'user_id'"`
	Object *T    `xorm:"extends"`
}
