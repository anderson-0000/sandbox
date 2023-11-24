package workflows _argoprojio


type WorkflowSpecTemplatesResourceManifestFromArtifactHttp struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	Auth *WorkflowSpecTemplatesResourceManifestFromArtifactHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	Headers *[]*WorkflowSpecTemplatesResourceManifestFromArtifactHttpHeaders `field:"optional" json:"headers" yaml:"headers"`
}

