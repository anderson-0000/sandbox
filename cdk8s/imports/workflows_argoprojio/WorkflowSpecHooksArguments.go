package workflows_argoprojio


type WorkflowSpecHooksArguments struct {
	Artifacts *[]*WorkflowSpecHooksArgumentsArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	Parameters *[]*WorkflowSpecHooksArgumentsParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

