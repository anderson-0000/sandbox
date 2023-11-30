package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContext struct {
	AllowPrivilegeEscalation *bool `field:"optional" json:"allowPrivilegeEscalation" yaml:"allowPrivilegeEscalation"`
	Capabilities *WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContextCapabilities `field:"optional" json:"capabilities" yaml:"capabilities"`
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	ProcMount *string `field:"optional" json:"procMount" yaml:"procMount"`
	ReadOnlyRootFilesystem *bool `field:"optional" json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	RunAsGroup *float64 `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	RunAsNonRoot *bool `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
	SeccompProfile *WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContextSeccompProfile `field:"optional" json:"seccompProfile" yaml:"seccompProfile"`
	SeLinuxOptions *WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContextSeLinuxOptions `field:"optional" json:"seLinuxOptions" yaml:"seLinuxOptions"`
	WindowsOptions *WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContextWindowsOptions `field:"optional" json:"windowsOptions" yaml:"windowsOptions"`
}

