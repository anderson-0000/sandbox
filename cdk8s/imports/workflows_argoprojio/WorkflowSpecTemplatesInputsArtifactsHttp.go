package workflows_argoprojio


type WorkflowSpecTemplatesInputsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecTemplatesInputsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecTemplatesInputsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

