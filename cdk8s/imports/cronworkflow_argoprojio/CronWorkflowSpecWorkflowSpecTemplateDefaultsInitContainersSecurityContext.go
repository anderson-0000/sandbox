package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContext struct {
	AllowPrivilegeEscalation *bool `field:"optional" json:"allowPrivilegeEscalation" yaml:"allowPrivilegeEscalation"`
	Capabilities *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContextCapabilities `field:"optional" json:"capabilities" yaml:"capabilities"`
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	ProcMount *string `field:"optional" json:"procMount" yaml:"procMount"`
	ReadOnlyRootFilesystem *bool `field:"optional" json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	RunAsGroup *float64 `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	RunAsNonRoot *bool `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
	SeccompProfile *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContextSeccompProfile `field:"optional" json:"seccompProfile" yaml:"seccompProfile"`
	SeLinuxOptions *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContextSeLinuxOptions `field:"optional" json:"seLinuxOptions" yaml:"seLinuxOptions"`
	WindowsOptions *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContextWindowsOptions `field:"optional" json:"windowsOptions" yaml:"windowsOptions"`
}

