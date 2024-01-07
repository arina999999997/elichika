package generic

type UserIdWrapper[T any] struct {
	UserId int `xorm:"pk 'user_id'"`
	Object *T  `xorm:"extends"` // slice of items
}
