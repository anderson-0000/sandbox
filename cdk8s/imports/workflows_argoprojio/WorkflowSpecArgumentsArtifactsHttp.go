package workflows_argoprojio


type WorkflowSpecArgumentsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecArgumentsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecArgumentsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

