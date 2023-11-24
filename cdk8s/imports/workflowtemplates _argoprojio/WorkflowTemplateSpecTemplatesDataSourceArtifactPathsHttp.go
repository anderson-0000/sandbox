package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

