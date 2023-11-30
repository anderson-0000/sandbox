package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesSecurityContext struct {
	FsGroup *float64 `field:"optional" json:"fsGroup" yaml:"fsGroup"`
	FsGroupChangePolicy *string `field:"optional" json:"fsGroupChangePolicy" yaml:"fsGroupChangePolicy"`
	RunAsGroup *float64 `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	RunAsNonRoot *bool `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
	SeccompProfile *WorkflowTemplateSpecTemplatesSecurityContextSeccompProfile `field:"optional" json:"seccompProfile" yaml:"seccompProfile"`
	SeLinuxOptions *WorkflowTemplateSpecTemplatesSecurityContextSeLinuxOptions `field:"optional" json:"seLinuxOptions" yaml:"seLinuxOptions"`
	SupplementalGroups *[]*float64 `field:"optional" json:"supplementalGroups" yaml:"supplementalGroups"`
	Sysctls *[]*WorkflowTemplateSpecTemplatesSecurityContextSysctls `field:"optional" json:"sysctls" yaml:"sysctls"`
	WindowsOptions *WorkflowTemplateSpecTemplatesSecurityContextWindowsOptions `field:"optional" json:"windowsOptions" yaml:"windowsOptions"`
}

