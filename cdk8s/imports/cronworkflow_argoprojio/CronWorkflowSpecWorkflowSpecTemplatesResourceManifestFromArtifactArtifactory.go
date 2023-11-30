package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactory struct {
	Url *string `field:"required" json:"url" yaml:"url"`
	PasswordSecret *CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactoryPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *CronWorkflowSpecWorkflowSpecTemplatesResourceManifestFromArtifactArtifactoryUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

