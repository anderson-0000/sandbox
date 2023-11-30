package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesSidecars struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	Env *[]*CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnv `field:"optional" json:"env" yaml:"env"`
	EnvFrom *[]*CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	Image *string `field:"optional" json:"image" yaml:"image"`
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	Lifecycle *CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	LivenessProbe *CronWorkflowSpecWorkflowSpecTemplatesSidecarsLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	MirrorVolumeMounts *bool `field:"optional" json:"mirrorVolumeMounts" yaml:"mirrorVolumeMounts"`
	Ports *[]*CronWorkflowSpecWorkflowSpecTemplatesSidecarsPorts `field:"optional" json:"ports" yaml:"ports"`
	ReadinessProbe *CronWorkflowSpecWorkflowSpecTemplatesSidecarsReadinessProbe `field:"optional" json:"readinessProbe" yaml:"readinessProbe"`
	Resources *CronWorkflowSpecWorkflowSpecTemplatesSidecarsResources `field:"optional" json:"resources" yaml:"resources"`
	SecurityContext *CronWorkflowSpecWorkflowSpecTemplatesSidecarsSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	StartupProbe *CronWorkflowSpecWorkflowSpecTemplatesSidecarsStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	Stdin *bool `field:"optional" json:"stdin" yaml:"stdin"`
	StdinOnce *bool `field:"optional" json:"stdinOnce" yaml:"stdinOnce"`
	TerminationMessagePath *string `field:"optional" json:"terminationMessagePath" yaml:"terminationMessagePath"`
	TerminationMessagePolicy *string `field:"optional" json:"terminationMessagePolicy" yaml:"terminationMessagePolicy"`
	Tty *bool `field:"optional" json:"tty" yaml:"tty"`
	VolumeDevices *[]*CronWorkflowSpecWorkflowSpecTemplatesSidecarsVolumeDevices `field:"optional" json:"volumeDevices" yaml:"volumeDevices"`
	VolumeMounts *[]*CronWorkflowSpecWorkflowSpecTemplatesSidecarsVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

