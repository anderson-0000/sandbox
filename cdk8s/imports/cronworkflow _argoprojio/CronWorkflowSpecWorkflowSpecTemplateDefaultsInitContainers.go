package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainers struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	Env *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnv `field:"optional" json:"env" yaml:"env"`
	EnvFrom *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	Image *string `field:"optional" json:"image" yaml:"image"`
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	Lifecycle *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	LivenessProbe *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	MirrorVolumeMounts *bool `field:"optional" json:"mirrorVolumeMounts" yaml:"mirrorVolumeMounts"`
	Ports *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersPorts `field:"optional" json:"ports" yaml:"ports"`
	ReadinessProbe *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersReadinessProbe `field:"optional" json:"readinessProbe" yaml:"readinessProbe"`
	Resources *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersResources `field:"optional" json:"resources" yaml:"resources"`
	SecurityContext *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	StartupProbe *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	Stdin *bool `field:"optional" json:"stdin" yaml:"stdin"`
	StdinOnce *bool `field:"optional" json:"stdinOnce" yaml:"stdinOnce"`
	TerminationMessagePath *string `field:"optional" json:"terminationMessagePath" yaml:"terminationMessagePath"`
	TerminationMessagePolicy *string `field:"optional" json:"terminationMessagePolicy" yaml:"terminationMessagePolicy"`
	Tty *bool `field:"optional" json:"tty" yaml:"tty"`
	VolumeDevices *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersVolumeDevices `field:"optional" json:"volumeDevices" yaml:"volumeDevices"`
	VolumeMounts *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

