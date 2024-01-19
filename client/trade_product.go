package client

import (
	"elichika/generic"
	// "reflect"
	// "encoding/json"
)

type TradeProduct struct {
	ProductId    int32                   `json:"product_id"`
	TradeId      int32                   `json:"trade_id"`
	SourceAmount int32                   `json:"source_amount"`
	StockAmount  generic.Nullable[int32] `xorm:"json" json:"stock_amount"`
	TradedCount  int32                   `xorm:"-" json:"traded_count"` // this is per user data
	Contents     generic.Array[Content]  `xorm:"json" json:"contents"`  // this can be in another table but this is easier
}

// for official server, trade_id = product_id for for whatever reason
// there doesn't seem to be any adverse effect of not doing that, but if it's necessary then we can use the custom mashaller
// func (t TradeProduct) MarshalJSON() ([]byte, error) {
// 	res := []byte{}
// 	res = append(res, []byte("{")...)

// 	var bytes []byte
// 	var err error
// 	for i := 0; i < reflect.TypeOf(t).NumField(); i++{
// 		if i > 0 {
// 			res = append(res, []byte(",")...)
// 		}
// 		res = append(res, []byte("\"")...)
// 		res = append(res, []byte(reflect.TypeOf(t).Field(i).Tag.Get("json"))...)
// 		res = append(res, []byte("\":")...)

// 		if i <= 1 {
// 			bytes, err = json.Marshal(t.ProductId)
// 		} else {
// 			bytes, err = json.Marshal(reflect.ValueOf(t).Field(i).Interface())
// 		}
// 		if err != nil {
// 			return res, err
// 		}
// 		res = append(res, bytes...)
// 	}
// 	res = append(res, []byte("}")...)
// 	return res, err
// }
