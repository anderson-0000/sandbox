package workflows _argoprojio


type WorkflowSpecTemplateDefaultsInputsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecTemplateDefaultsInputsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecTemplateDefaultsInputsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

