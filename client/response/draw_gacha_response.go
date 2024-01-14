package response

import (
	"elichika/client"
	"elichika/generic"
)

type DrawGachaResponse struct {
	Gacha          client.Gacha                                   `json:"gacha"`
	ResultCards    generic.List[client.AddedGachaCardResult]      `json:"result_cards"`
	ResultBonuses  generic.Nullable[generic.List[client.Content]] `json:"result_bonuses"`
	RetryGacha     generic.Nullable[client.RetryGacha]            `json:"retry_gacha"`      //pointer
	StepupNextStep generic.Nullable[client.GachaDrawStepupNext]   `json:"stepup_next_step"` //pointer
	UserModelDiff  *client.UserModel                              `json:"user_model_diff"`
}
