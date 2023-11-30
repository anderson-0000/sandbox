package workflows_argoprojio


type WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuth struct {
	PasswordSecret *WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuthPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	UsernameSecret *WorkflowSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuthUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

