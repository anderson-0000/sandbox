package workflows_argoprojio


type WorkflowSpecTemplatesResourceManifestFromArtifactArchive struct {
	None interface{} `field:"optional" json:"none" yaml:"none"`
	Tar *WorkflowSpecTemplatesResourceManifestFromArtifactArchiveTar `field:"optional" json:"tar" yaml:"tar"`
	Zip interface{} `field:"optional" json:"zip" yaml:"zip"`
}

