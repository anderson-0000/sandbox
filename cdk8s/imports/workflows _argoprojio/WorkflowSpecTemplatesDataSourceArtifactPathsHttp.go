package workflows _argoprojio


type WorkflowSpecTemplatesDataSourceArtifactPathsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecTemplatesDataSourceArtifactPathsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecTemplatesDataSourceArtifactPathsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

