package repository

import (
	"capel-meister/pkg/logs"
	"github.com/go-git/go-git/v5"
)

/*
RepoMetadata has metadata
for the repository, worktree and submodules
*/
type RepoMetadata struct {
    repository  *git.Repository
    worktree    *git.Worktree
};

/* 
CloneRepository clones the repo 
from the `repoUrl` specified 
to the repoPath specified
*/
func CloneRepository (repoUrl string) *RepoMetadata {
    repoPath := PathFromRepoURL(repoUrl);
    repo, err := git.PlainClone(repoPath, false, &git.CloneOptions{ 
        URL: repoUrl,
        RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
    });
    logs.Error(err);

    worktree, err := repo.Worktree();
    logs.Error(err);

    return &RepoMetadata{repo, worktree};
}

/*
LoadRepository loads the repo if `.git` dir
if found in the `repoPath`
*/
func LoadRepository (repoUrl string) *RepoMetadata {
    repoPath := PathFromRepoURL(repoUrl);
    repo, err := git.PlainOpen(repoPath);
    logs.Error(err);

    worktree, err := repo.Worktree();
    logs.Error(err);

    return &RepoMetadata{repo, worktree};
}

/* 
PullRepository pulls the repository represented by RepoMetadata
*/
func (repoMetadata *RepoMetadata) PullRepository () {
    err := repoMetadata.worktree.Pull(&git.PullOptions{
        RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
    });
    if err == git.NoErrAlreadyUpToDate {
        logs.Info("Already up to date.");
    } else {
        logs.Error(err);
    }
}

/*
CheckoutBranch checks out the branch for the repo
*/
func (repoMetadata *RepoMetadata) CheckoutBranch (branchName string) {
    branch, err := repoMetadata.repository.Branch(branchName);
    logs.Error(err);

    err = repoMetadata.worktree.Checkout(&git.CheckoutOptions{
        Branch: branch.Merge,
    });
    logs.Error(err);
}

