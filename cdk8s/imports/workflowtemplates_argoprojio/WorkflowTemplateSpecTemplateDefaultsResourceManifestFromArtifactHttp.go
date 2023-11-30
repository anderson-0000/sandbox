package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

