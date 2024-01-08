package client

type Content struct {
	ContentType   int32 `xorm:"'content_type'" json:"content_type" enum:"ContentType"`
	ContentId     int32 `xorm:"'content_id'" json:"content_id"`
	ContentAmount int32 `xorm:"'content_amount'" json:"content_amount"`
}

func (c *Content) Amount(amount int32) Content {
	return Content{
		ContentType:   c.ContentType,
		ContentId:     c.ContentId,
		ContentAmount: amount,
	}
}
