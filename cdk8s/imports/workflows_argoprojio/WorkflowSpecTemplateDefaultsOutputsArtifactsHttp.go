package workflows_argoprojio


type WorkflowSpecTemplateDefaultsOutputsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecTemplateDefaultsOutputsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecTemplateDefaultsOutputsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

