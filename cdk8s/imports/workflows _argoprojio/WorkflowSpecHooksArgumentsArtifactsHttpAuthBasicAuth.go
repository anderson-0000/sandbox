package workflows _argoprojio


type WorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuth struct {
	PasswordSecret *WorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowSpecHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

