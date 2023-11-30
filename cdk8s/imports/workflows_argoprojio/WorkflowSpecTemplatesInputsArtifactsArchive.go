package workflows_argoprojio


type WorkflowSpecTemplatesInputsArtifactsArchive struct {
	None interface{} `field:"optional" json:"none" yaml:"none"`
	Tar *WorkflowSpecTemplatesInputsArtifactsArchiveTar `field:"optional" json:"tar" yaml:"tar"`
	Zip interface{} `field:"optional" json:"zip" yaml:"zip"`
}

