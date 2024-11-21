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
	response.MessageJa = "サーバーはメンテナンス中です。後ほどもう一度確認してください。\n独自のサーバーを実行している場合は、管理 WebUI に移動して elichika を再起動してください。"
	response.MessageEn = "The server is in maintenance, check again later.\nIf you are running your own server, go to the admin webui and restart elichika."
	response.MessageZh = "伺服器正在維護中，請稍後再查看。\n如果您正在運行自己的伺服器，請前往管理 WebUI 並重新啟動 elichika。"
	response.MessageKo = "서버가 유지 관리 중입니다. 나중에 다시 확인하세요.\n직접 서버를 실행하고 있다면 관리자 webui로 가서 elichika를 다시 시작하세요."
	response.MessageTh = "เซิร์ฟเวอร์กำลังอยู่ในระหว่างการบำรุงรักษา โปรดตรวจสอบอีกครั้งในภายหลัง\nหากคุณใช้งานเซิร์ฟเวอร์ของคุณเอง ให้ไปที่เว็บ UI ของผู้ดูแลระบบและรีสตาร์ท elichika"

	response.UrlJa = *config.Conf.MaintenanceUrl
	response.UrlEn = *config.Conf.MaintenanceUrl
	response.UrlZh = *config.Conf.MaintenanceUrl
	response.UrlKo = *config.Conf.MaintenanceUrl
	response.UrlTh = *config.Conf.MaintenanceUrl
}

func maintenanceClient(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, response)
}

var webuiResponse = ""

func init() {
	webuiResponse += "<html>\n"
	webuiResponse += "<div>" + response.MessageEn + "</div>\n"
	webuiResponse += "<div>" + response.MessageJa + "</div>\n"
	webuiResponse += "<div>" + response.MessageZh + "</div>\n"
	webuiResponse += "<div>" + response.MessageKo + "</div>\n"
	webuiResponse += "<div>" + response.MessageTh + "</div>\n"
	webuiResponse += "</html>\n"
}

func maintenanceWebUi(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusServiceUnavailable)
	ctx.Writer.Write([]byte(webuiResponse))
}
