package workflows_argoprojio


type WorkflowSpecTemplatesOutputs struct {
	Artifacts *[]*WorkflowSpecTemplatesOutputsArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	ExitCode *string `field:"optional" json:"exitCode" yaml:"exitCode"`
	Parameters *[]*WorkflowSpecTemplatesOutputsParameters `field:"optional" json:"parameters" yaml:"parameters"`
	Result *string `field:"optional" json:"result" yaml:"result"`
}

