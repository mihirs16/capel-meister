package webhook

import "github.com/gin-gonic/gin"

func GitHubWebhook (context *gin.Context) {
    context.JSON(200, gin.H {
        "success": "true",
    });
}
