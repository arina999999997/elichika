package client

type UserReferenceBook struct {
	ReferenceBookId int32 `xorm:"pk 'reference_book_id'" json:"reference_book_id"`
}
