package workflows _argoprojio


type WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecTemplateDefaultsDataSourceArtifactPathsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

