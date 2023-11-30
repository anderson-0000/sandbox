package workflows_argoprojio


type WorkflowSpecTemplatesDagTasksArguments struct {
	Artifacts *[]*WorkflowSpecTemplatesDagTasksArgumentsArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	Parameters *[]*WorkflowSpecTemplatesDagTasksArgumentsParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

