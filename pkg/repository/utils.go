package repository

import (
    "os"
    "net/url"
    "strings"
    "path/filepath"
    "capel-meister/pkg/logs"
);

func PathFromRepoURL (repoUrl string) string {
    parsedUrl, err := url.Parse(repoUrl);
    logs.Error(err);

    repoPath := filepath.Join(
        os.Getenv("GITHUB_PATH"), 
        strings.Split(parsedUrl.Path, "/")[2],
    );
    logs.Info(repoPath);

    return repoPath;
}

