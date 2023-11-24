package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsScript struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Source *string `field:"required" json:"source" yaml:"source"`
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	Env *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnv `field:"optional" json:"env" yaml:"env"`
	EnvFrom *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	Image *string `field:"optional" json:"image" yaml:"image"`
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	Lifecycle *CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	LivenessProbe *CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	Ports *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptPorts `field:"optional" json:"ports" yaml:"ports"`
	ReadinessProbe *CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptReadinessProbe `field:"optional" json:"readinessProbe" yaml:"readinessProbe"`
	Resources *CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptResources `field:"optional" json:"resources" yaml:"resources"`
	SecurityContext *CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	StartupProbe *CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	Stdin *bool `field:"optional" json:"stdin" yaml:"stdin"`
	StdinOnce *bool `field:"optional" json:"stdinOnce" yaml:"stdinOnce"`
	TerminationMessagePath *string `field:"optional" json:"terminationMessagePath" yaml:"terminationMessagePath"`
	TerminationMessagePolicy *string `field:"optional" json:"terminationMessagePolicy" yaml:"terminationMessagePolicy"`
	Tty *bool `field:"optional" json:"tty" yaml:"tty"`
	VolumeDevices *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptVolumeDevices `field:"optional" json:"volumeDevices" yaml:"volumeDevices"`
	VolumeMounts *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

