package workflows _argoprojio


type WorkflowSpecTemplateDefaultsOutputs struct {
	Artifacts *[]*WorkflowSpecTemplateDefaultsOutputsArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	ExitCode *string `field:"optional" json:"exitCode" yaml:"exitCode"`
	Parameters *[]*WorkflowSpecTemplateDefaultsOutputsParameters `field:"optional" json:"parameters" yaml:"parameters"`
	Result *string `field:"optional" json:"result" yaml:"result"`
}

