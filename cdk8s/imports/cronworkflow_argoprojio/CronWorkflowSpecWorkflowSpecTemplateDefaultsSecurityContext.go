package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContext struct {
	FsGroup *float64 `field:"optional" json:"fsGroup" yaml:"fsGroup"`
	FsGroupChangePolicy *string `field:"optional" json:"fsGroupChangePolicy" yaml:"fsGroupChangePolicy"`
	RunAsGroup *float64 `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	RunAsNonRoot *bool `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
	SeccompProfile *CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContextSeccompProfile `field:"optional" json:"seccompProfile" yaml:"seccompProfile"`
	SeLinuxOptions *CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContextSeLinuxOptions `field:"optional" json:"seLinuxOptions" yaml:"seLinuxOptions"`
	SupplementalGroups *[]*float64 `field:"optional" json:"supplementalGroups" yaml:"supplementalGroups"`
	Sysctls *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContextSysctls `field:"optional" json:"sysctls" yaml:"sysctls"`
	WindowsOptions *CronWorkflowSpecWorkflowSpecTemplateDefaultsSecurityContextWindowsOptions `field:"optional" json:"windowsOptions" yaml:"windowsOptions"`
}

