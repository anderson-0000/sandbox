package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecArguments struct {
	Artifacts *[]*CronWorkflowSpecWorkflowSpecArgumentsArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	Parameters *[]*CronWorkflowSpecWorkflowSpecArgumentsParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

