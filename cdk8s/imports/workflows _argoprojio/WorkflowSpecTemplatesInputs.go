package workflows _argoprojio


type WorkflowSpecTemplatesInputs struct {
	Artifacts *[]*WorkflowSpecTemplatesInputsArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	Parameters *[]*WorkflowSpecTemplatesInputsParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

