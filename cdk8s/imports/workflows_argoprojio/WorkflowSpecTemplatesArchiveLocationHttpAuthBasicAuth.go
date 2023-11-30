package workflows_argoprojio


type WorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuth struct {
	PasswordSecret *WorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuthPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowSpecTemplatesArchiveLocationHttpAuthBasicAuthUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

