package repository

import (
	"capel-meister/pkg/logs"
	"github.com/go-git/go-git/v5"
);

/*
RepoMetadata has metadata 
for the repository, worktree and submodules 
*/
type RepoMetadata struct {
    repository  *git.Repository
    worktree    *git.Worktree
    submodules  git.Submodules
};

func LoadRepository (repoPath string) *RepoMetadata {
    repo, err := git.PlainOpen(repoPath);
    logs.Error(err);

    worktree, err := repo.Worktree();
    logs.Error(err);

    submodules, err := worktree.Submodules();
    logs.Error(err);

    repoMetadata := RepoMetadata{repo, worktree, submodules};
    return &repoMetadata;
}

func (repoMetadata *RepoMetadata) FetchRepository () {
    err := repoMetadata.repository.Fetch(&git.FetchOptions{});
    logs.Error(err);
}

// func () PullRepository () {}
