package workflows _argoprojio


type WorkflowSpecTemplatesSidecarsSecurityContextWindowsOptions struct {
	GmsaCredentialSpec *string `field:"optional" json:"gmsaCredentialSpec" yaml:"gmsaCredentialSpec"`
	GmsaCredentialSpecName *string `field:"optional" json:"gmsaCredentialSpecName" yaml:"gmsaCredentialSpecName"`
	HostProcess *bool `field:"optional" json:"hostProcess" yaml:"hostProcess"`
	RunAsUserName *string `field:"optional" json:"runAsUserName" yaml:"runAsUserName"`
}

