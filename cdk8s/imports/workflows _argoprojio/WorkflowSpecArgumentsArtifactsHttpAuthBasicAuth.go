package workflows _argoprojio


type WorkflowSpecArgumentsArtifactsHttpAuthBasicAuth struct {
	PasswordSecret *WorkflowSpecArgumentsArtifactsHttpAuthBasicAuthPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowSpecArgumentsArtifactsHttpAuthBasicAuthUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

