package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInputsArtifactsGit struct {
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	Depth *float64 `field:"optional" json:"depth" yaml:"depth"`
	DisableSubmodules *bool `field:"optional" json:"disableSubmodules" yaml:"disableSubmodules"`
	Fetch *[]*string `field:"optional" json:"fetch" yaml:"fetch"`
	InsecureIgnoreHostKey *bool `field:"optional" json:"insecureIgnoreHostKey" yaml:"insecureIgnoreHostKey"`
	PasswordSecret *WorkflowSpecTemplateDefaultsInputsArtifactsGitPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
	SingleBranch *bool `field:"optional" json:"singleBranch" yaml:"singleBranch"`
	SshPrivateKeySecret *WorkflowSpecTemplateDefaultsInputsArtifactsGitSshPrivateKeySecret `field:"optional" json:"sshPrivateKeySecret" yaml:"sshPrivateKeySecret"`
	UsernameSecret *WorkflowSpecTemplateDefaultsInputsArtifactsGitUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

