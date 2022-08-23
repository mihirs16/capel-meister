package server

import (
    "os"
	"capel-meister/pkg/actions"
	"capel-meister/pkg/repository"
	"github.com/gin-gonic/gin"
);


func NewDeployment (context *gin.Context) {
    repoMetadata := repository.RepoMetadata{};
    context.BindJSON(&repoMetadata);
    repoMetadata.SaveMetadata(os.Getenv("DEPLOYMENT_FILE_PATH"));
    actions.Deploy([]repository.RepoMetadata{repoMetadata});
    context.JSON(200, repoMetadata)
}


