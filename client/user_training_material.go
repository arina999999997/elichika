package client

import (
	"elichika/enum"

	"fmt"
)

type UserTrainingMaterial struct {
	TrainingMaterialMasterId int32 `json:"training_material_master_id"`
	Amount                   int32 `json:"amount"`
}

func (utm *UserTrainingMaterial) Id() int64 {
	return int64(utm.TrainingMaterialMasterId)
}
func (utm *UserTrainingMaterial) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeTrainingMaterial { // 12
		panic(fmt.Sprintln("Wrong content for TrainingMaterial: ", content))
	}
	utm.TrainingMaterialMasterId = content.ContentId
	utm.Amount = content.ContentAmount
}
func (utm *UserTrainingMaterial) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeTrainingMaterial,
		ContentId:     utm.TrainingMaterialMasterId,
		ContentAmount: utm.Amount,
	}
}
