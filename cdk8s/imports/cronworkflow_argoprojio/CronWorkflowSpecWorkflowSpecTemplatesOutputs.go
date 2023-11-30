package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesOutputs struct {
	Artifacts *[]*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	ExitCode *string `field:"optional" json:"exitCode" yaml:"exitCode"`
	Parameters *[]*CronWorkflowSpecWorkflowSpecTemplatesOutputsParameters `field:"optional" json:"parameters" yaml:"parameters"`
	Result *string `field:"optional" json:"result" yaml:"result"`
}

