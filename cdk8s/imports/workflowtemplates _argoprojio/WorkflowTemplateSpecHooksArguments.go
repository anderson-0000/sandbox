package workflowtemplates _argoprojio


type WorkflowTemplateSpecHooksArguments struct {
	Artifacts *[]*WorkflowTemplateSpecHooksArgumentsArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	Parameters *[]*WorkflowTemplateSpecHooksArgumentsParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

