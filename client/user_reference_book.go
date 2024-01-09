package client

type UserReferenceBook struct {
	ReferenceBookId int32 `xorm:"pk 'reference_book_id'" json:"reference_book_id"`
}

func (urb *UserReferenceBook) Id() int64 {
	return int64(urb.ReferenceBookId)
}
