package repository

import (
    "capel-meister/pkg/utils"
);


/*
IsEqualTo compares one (L) RepoMetadata to another (R).
*/
func (L *RepoMetadata) IsEqualTo (R *RepoMetadata) bool {
    urlEqual := L.URL == R.URL;
    branchEqual := L.BranchName == R.BranchName;
    return urlEqual && branchEqual;
}


/*
LoadMetadata loads the deployment metadata from the JSON file.
*/
func LoadMetadata (file string) []RepoMetadata {
    repoMetadataList := []RepoMetadata{};
    utils.LoadJSONFromFile(file, &repoMetadataList)
    return repoMetadataList
}


/*
WriteMetadata appends a new deployment metadata to the JSON file.
*/
func (repoMetadata *RepoMetadata) SaveMetadata (file string) {
    deployMetadata := LoadMetadata(file);
    for _, eachDeployMetadata := range deployMetadata {
        if eachDeployMetadata.IsEqualTo(repoMetadata) {
            panic("Deployment Exists");
        }
    };
    deployMetadata = append(deployMetadata, *repoMetadata);
    utils.WriteJSONToFile(file, deployMetadata);
}

