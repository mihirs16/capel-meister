package repository


/*
LoadMetadata loads the deployment metadata from the JSON file.
*/
func LoadMetadata (file string) []RepoMetadata {
    repoMetadataList := []RepoMetadata{};
    LoadJSONFromFile(file, &repoMetadataList)
    return repoMetadataList
}

/*
WriteMetadata appends a new deployment metadata to the JSON file.
*/
func WriteMetadata (file string, repoMetadata RepoMetadata) {
    deployMetadata := LoadMetadata(file);
    deployMetadata = append(deployMetadata, repoMetadata);
    WriteJSONToFile(file, deployMetadata);
}
