package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInputs struct {
	Artifacts *[]*WorkflowSpecTemplateDefaultsInputsArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	Parameters *[]*WorkflowSpecTemplateDefaultsInputsParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

