package server

import (
	"capel-meister/pkg/webhook"
	"github.com/gin-gonic/gin"
);

func BuildRouter () {
    gin.SetMode(gin.ReleaseMode);
    router := gin.New();
    router.SetTrustedProxies([]string{"127.0.0.1"});

    // endpoints
    router.GET("/", webhook.GitHubWebhook);

    router.Run("127.0.0.1:8080");
}
