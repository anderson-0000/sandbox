package workflows_argoprojio


type WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGit struct {
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	Depth *float64 `field:"optional" json:"depth" yaml:"depth"`
	DisableSubmodules *bool `field:"optional" json:"disableSubmodules" yaml:"disableSubmodules"`
	Fetch *[]*string `field:"optional" json:"fetch" yaml:"fetch"`
	InsecureIgnoreHostKey *bool `field:"optional" json:"insecureIgnoreHostKey" yaml:"insecureIgnoreHostKey"`
	PasswordSecret *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitPasswordSecret `field:"optional" json:"passwordSecret" yaml:"passwordSecret"`
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
	SingleBranch *bool `field:"optional" json:"singleBranch" yaml:"singleBranch"`
	SshPrivateKeySecret *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitSshPrivateKeySecret `field:"optional" json:"sshPrivateKeySecret" yaml:"sshPrivateKeySecret"`
	UsernameSecret *WorkflowSpecTemplateDefaultsDataSourceArtifactPathsGitUsernameSecret `field:"optional" json:"usernameSecret" yaml:"usernameSecret"`
}

