package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*CronWorkflowSpecWorkflowSpecArgumentsArtifactsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

