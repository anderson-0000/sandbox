package workflowtemplates _argoprojio


type WorkflowTemplateSpecArguments struct {
	Artifacts *[]*WorkflowTemplateSpecArgumentsArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	Parameters *[]*WorkflowTemplateSpecArgumentsParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

