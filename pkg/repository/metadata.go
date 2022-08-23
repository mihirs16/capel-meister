package repository

import (
    "capel-meister/pkg/utils"
);


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
        if eachDeployMetadata == *repoMetadata {
            panic("Deployment Exists");
        }
    };
    deployMetadata = append(deployMetadata, *repoMetadata);
    utils.WriteJSONToFile(file, deployMetadata);
}

