package workflows _argoprojio


type WorkflowSpecTemplatesOutputsArtifactsArchive struct {
	None interface{} `field:"optional" json:"none" yaml:"none"`
	Tar *WorkflowSpecTemplatesOutputsArtifactsArchiveTar `field:"optional" json:"tar" yaml:"tar"`
	Zip interface{} `field:"optional" json:"zip" yaml:"zip"`
}

