package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

