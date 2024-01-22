package client

import (
	"elichika/enum"

	"fmt"
)

type UserTrainingMaterial struct {
	TrainingMaterialMasterId int32 `json:"training_material_master_id"`
	Amount                   int32 `json:"amount"`
}

func (utm *UserTrainingMaterial) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeTrainingMaterial { // 12
		panic(fmt.Sprintln("Wrong content for TrainingMaterial: ", content))
	}
	utm.TrainingMaterialMasterId = content.ContentId
	utm.Amount = content.ContentAmount
}
func (utm *UserTrainingMaterial) ToContent(contentId int32) Content {
	return Content{
		ContentType:   enum.ContentTypeTrainingMaterial,
		ContentId:     contentId,
		ContentAmount: utm.Amount,
	}
}
