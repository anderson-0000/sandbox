package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesOutputs struct {
	Artifacts *[]*WorkflowTemplateSpecTemplatesOutputsArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	ExitCode *string `field:"optional" json:"exitCode" yaml:"exitCode"`
	Parameters *[]*WorkflowTemplateSpecTemplatesOutputsParameters `field:"optional" json:"parameters" yaml:"parameters"`
	Result *string `field:"optional" json:"result" yaml:"result"`
}

