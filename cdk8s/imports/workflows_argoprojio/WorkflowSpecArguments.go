package workflows_argoprojio


type WorkflowSpecArguments struct {
	Artifacts *[]*WorkflowSpecArgumentsArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	Parameters *[]*WorkflowSpecArgumentsParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

