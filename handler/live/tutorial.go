package live

import (
	"github.com/gin-gonic/gin"
)

// doesn't really need to do anything here?
// TODO(refactor): Change to use request and response types
func FinishTutorial(ctx *gin.Context) {
	LiveFinish(ctx)
}
