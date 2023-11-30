package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

