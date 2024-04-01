package notice

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/router"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// TODO(notice): This is placeholder, we would probably want something like load from a database
func fetchNoticeDetail(ctx *gin.Context) {
	req := request.FetchNoticeDetailRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)
	// there is no request body

	common.JsonResponse(ctx, response.FetchNoticeDetailResponse{
		Notice: client.NoticeDetail{
			NoticeId: req.NoticeId,
			Category: enum.NoticeTagCategoryInformation,
			Title: client.LocalizedText{
				DotUnderText: "Game Operations update",
			},
			DetailText: client.LocalizedText{
				DotUnderText: "This is the Love Live! School Idol Festival All Stars game operations team.\n\nOnce again, we would like to thank everyone who has played SIFAS, whether you've been playing up until this point, or even have only played just once.\n\nWe received a variety of opinions following our operations news announcement on April 30.\nFirst of all, we would like to thank you all for sharing your opinions with us.\nThe Operations Team has heard your voices and shares your passion for SIFAS.\nWe have read all of the opinions that we received.\nAs such, we truly regret that we cannot respond to every request due to the high volume of player responses.\n\nWe would like to take this opportunity to deliver some responses from the Operations Team.\nWe heard from many people who hope that the Story and Bond Episodes will continue after this point.\nThe Operations Team is considering ways to allow you all to keep following the All Star Story.\nWe believe that someday, we will be able to announce this in a form that will please all of you and hope you will stay tuned to find out more. \n\nWe have also been working on plans to have SIFAS art and stories remain available to you moving forward.\nAs a first step, Love Live! School Idol Festival All Stars Event Memories will go on sale on June 26. This book will offer popular SIFAS event stories, as voted on by fans, in comic form.\nWe are also currently considering other projects, which we hope to announce to all of you soon.\n\nWe have received comments from many people wondering if there will be a way to continue to watch the 3D music videos.\nAs we announced previously, we will hold the Love Live! SIF 2023 Series Fan Festival All Stars Metaverse Live Show online event on July 8 and July 9 (JST).\nWe would like to continue to hold fun events like this that make use of the SIFAS 3D music videos.\nWe will also consider additional projects to convey the charm of school idols to even more people, so please stay tuned for more.\n\nStarting today, we will begin concentrating on Love Live! School Idol Festival 2: Miracle Live!, the new School Idol Festival series game app.\nWe hope to reward your love of Love Live! School Idol Festival and Love Live! School Idol Festival All Stars by bringing you even more heart-pounding fun that will elevate the Love Live! series to new heights of excitement. \n\nFinally, from the bottom of our hearts, we would like to thank all of you for loving Love Live! School Idol Festival All Stars.\nWe appreciate every single one of you for supporting and cheering on so many school idols as the president of the Nijigasaki High School Idol Club and as a fan of school idols.\nThe Love Live! series still holds a lot in store for the Nijigasaki High School Idol Club, ..'s, Aqours, Liella!, Hasunosora Girls' High School Idol Club, and School Idol Musical.\nAs such, we hope you will continue to cheer on the remarkable school idols of the Love Live! series.\n\nWe also hope that you will continue to enjoy the Love Live! and School Idol Festival series.\n\n",
			},
			Date: 1688104800, // 2023/06/30 15:00 Asia/Tokyo
		},
	})
}

func init() {
	router.AddHandler("/notice/fetchNoticeDetail", fetchNoticeDetail)
}
