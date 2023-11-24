package workflows _argoprojio


type WorkflowSpecTemplatesOutputsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecTemplatesOutputsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecTemplatesOutputsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

