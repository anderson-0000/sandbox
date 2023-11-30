package workflows_argoprojio


type WorkflowSpecArgumentsArtifactsArchive struct {
	None interface{} `field:"optional" json:"none" yaml:"none"`
	Tar *WorkflowSpecArgumentsArtifactsArchiveTar `field:"optional" json:"tar" yaml:"tar"`
	Zip interface{} `field:"optional" json:"zip" yaml:"zip"`
}

