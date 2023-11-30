package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*CronWorkflowSpecWorkflowSpecTemplatesDataSourceArtifactPathsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

