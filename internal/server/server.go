package server

import (
	"github.com/gin-gonic/gin"
);


func BuildRouter (mode string) {
    gin.SetMode(mode);
    router := gin.New();
    router.SetTrustedProxies([]string{"127.0.0.1"});
    router.Use(gin.Recovery());
    if gin.Mode() == gin.DebugMode {
        router.Use(gin.Logger());
    };
  
    deployRouter := router.Group("/deploy") 
    {
        deployRouter.POST("/new", NewDeployment);
    };

    router.Run(":8080");
}
