package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

