package main

import (
	"elichika/config"

	"net/http"

	"github.com/gin-gonic/gin"
)

type MaintenanceResponse struct {
	MessageJa string `json:"message_ja"`
	MessageEn string `json:"message_en"`
	MessageZh string `json:"message_zh"`
	MessageKo string `json:"message_ko"`
	MessageTh string `json:"message_th"`
	UrlJa     string `json:"url_ja"`
	UrlEn     string `json:"url_en"`
	UrlZh     string `json:"url_zh"`
	UrlKo     string `json:"url_ko"`
	UrlTh     string `json:"url_th"`
}

var response MaintenanceResponse

func init() {
	// TODO(extra): add a way to modify these without rebuilding the code
	response.MessageJa = "サーバーはメンテナンス中です。後ほどもう一度ご確認ください。\n独自のサーバーを実行している場合は、WebUI に移動してそこから elichika を実行します。"
	response.MessageEn = "The server is in maintenance, check again later.\nIf you are running your own server, go to the webui and run elichika from there."
	response.MessageZh = "伺服器正在維護中，請稍後再查看。\n如果您正在運行自己的伺服器，請轉到 webui 並從那裡運行 elichika。"
	response.MessageKo = "서버가 유지 관리 중이므로 나중에 다시 확인하세요.자신의 서버를 운영하고 있다면 webui로 가서 거기에서 elichika를 실행하세요."
	response.MessageTh = "เซิร์ฟเวอร์อยู่ระหว่างการบำรุงรักษา กรุณาตรวจสอบอีกครั้งในภายหลัง\nหากคุณกำลังรันเซิร์ฟเวอร์ของตนเอง ให้ไปที่ webui และรัน elichika จากที่นั่น"

	response.UrlJa = *config.Conf.MaintenanceUrl
	response.UrlEn = *config.Conf.MaintenanceUrl
	response.UrlZh = *config.Conf.MaintenanceUrl
	response.UrlKo = *config.Conf.MaintenanceUrl
	response.UrlTh = *config.Conf.MaintenanceUrl
}

func maintenance(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, response)
}
