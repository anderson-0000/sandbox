package workflows_argoprojio


type WorkflowSpecHooksArgumentsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecHooksArgumentsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecHooksArgumentsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

