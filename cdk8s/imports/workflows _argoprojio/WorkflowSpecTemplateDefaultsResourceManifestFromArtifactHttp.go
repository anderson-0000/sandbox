package workflows _argoprojio


type WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecTemplateDefaultsResourceManifestFromArtifactHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

