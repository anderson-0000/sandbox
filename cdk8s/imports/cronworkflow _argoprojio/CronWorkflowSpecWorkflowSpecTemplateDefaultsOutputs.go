package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputs struct {
	Artifacts *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	ExitCode *string `field:"optional" json:"exitCode" yaml:"exitCode"`
	Parameters *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsOutputsParameters `field:"optional" json:"parameters" yaml:"parameters"`
	Result *string `field:"optional" json:"result" yaml:"result"`
}

