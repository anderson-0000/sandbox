package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsSecurityContext struct {
	FsGroup *float64 `field:"optional" json:"fsGroup" yaml:"fsGroup"`
	FsGroupChangePolicy *string `field:"optional" json:"fsGroupChangePolicy" yaml:"fsGroupChangePolicy"`
	RunAsGroup *float64 `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	RunAsNonRoot *bool `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
	SeccompProfile *WorkflowTemplateSpecTemplateDefaultsSecurityContextSeccompProfile `field:"optional" json:"seccompProfile" yaml:"seccompProfile"`
	SeLinuxOptions *WorkflowTemplateSpecTemplateDefaultsSecurityContextSeLinuxOptions `field:"optional" json:"seLinuxOptions" yaml:"seLinuxOptions"`
	SupplementalGroups *[]*float64 `field:"optional" json:"supplementalGroups" yaml:"supplementalGroups"`
	Sysctls *[]*WorkflowTemplateSpecTemplateDefaultsSecurityContextSysctls `field:"optional" json:"sysctls" yaml:"sysctls"`
	WindowsOptions *WorkflowTemplateSpecTemplateDefaultsSecurityContextWindowsOptions `field:"optional" json:"windowsOptions" yaml:"windowsOptions"`
}

