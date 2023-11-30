package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

