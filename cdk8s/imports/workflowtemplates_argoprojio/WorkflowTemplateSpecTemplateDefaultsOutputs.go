package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsOutputs struct {
	Artifacts *[]*WorkflowTemplateSpecTemplateDefaultsOutputsArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	ExitCode *string `field:"optional" json:"exitCode" yaml:"exitCode"`
	Parameters *[]*WorkflowTemplateSpecTemplateDefaultsOutputsParameters `field:"optional" json:"parameters" yaml:"parameters"`
	Result *string `field:"optional" json:"result" yaml:"result"`
}

