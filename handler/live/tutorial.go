package live

import (
	"github.com/gin-gonic/gin"
)

// doesn't really need to do anything here?
func FinishTutorial(ctx *gin.Context) {
	LiveFinish(ctx)
}
