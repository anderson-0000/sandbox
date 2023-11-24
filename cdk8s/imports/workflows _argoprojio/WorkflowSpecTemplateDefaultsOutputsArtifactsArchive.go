package workflows _argoprojio


type WorkflowSpecTemplateDefaultsOutputsArtifactsArchive struct {
	None interface{} `field:"optional" json:"none" yaml:"none"`
	Tar *WorkflowSpecTemplateDefaultsOutputsArtifactsArchiveTar `field:"optional" json:"tar" yaml:"tar"`
	Zip interface{} `field:"optional" json:"zip" yaml:"zip"`
}

