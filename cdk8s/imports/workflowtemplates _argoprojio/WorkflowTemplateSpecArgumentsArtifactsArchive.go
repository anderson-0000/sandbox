package workflowtemplates _argoprojio


type WorkflowTemplateSpecArgumentsArtifactsArchive struct {
	None interface{} `field:"optional" json:"none" yaml:"none"`
	Tar *WorkflowTemplateSpecArgumentsArtifactsArchiveTar `field:"optional" json:"tar" yaml:"tar"`
	Zip interface{} `field:"optional" json:"zip" yaml:"zip"`
}

