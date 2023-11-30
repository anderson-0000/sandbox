package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInitContainers struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	Env *[]*WorkflowSpecTemplateDefaultsInitContainersEnv `field:"optional" json:"env" yaml:"env"`
	EnvFrom *[]*WorkflowSpecTemplateDefaultsInitContainersEnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	Image *string `field:"optional" json:"image" yaml:"image"`
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	Lifecycle *WorkflowSpecTemplateDefaultsInitContainersLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	LivenessProbe *WorkflowSpecTemplateDefaultsInitContainersLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	MirrorVolumeMounts *bool `field:"optional" json:"mirrorVolumeMounts" yaml:"mirrorVolumeMounts"`
	Ports *[]*WorkflowSpecTemplateDefaultsInitContainersPorts `field:"optional" json:"ports" yaml:"ports"`
	ReadinessProbe *WorkflowSpecTemplateDefaultsInitContainersReadinessProbe `field:"optional" json:"readinessProbe" yaml:"readinessProbe"`
	Resources *WorkflowSpecTemplateDefaultsInitContainersResources `field:"optional" json:"resources" yaml:"resources"`
	SecurityContext *WorkflowSpecTemplateDefaultsInitContainersSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	StartupProbe *WorkflowSpecTemplateDefaultsInitContainersStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	Stdin *bool `field:"optional" json:"stdin" yaml:"stdin"`
	StdinOnce *bool `field:"optional" json:"stdinOnce" yaml:"stdinOnce"`
	TerminationMessagePath *string `field:"optional" json:"terminationMessagePath" yaml:"terminationMessagePath"`
	TerminationMessagePolicy *string `field:"optional" json:"terminationMessagePolicy" yaml:"terminationMessagePolicy"`
	Tty *bool `field:"optional" json:"tty" yaml:"tty"`
	VolumeDevices *[]*WorkflowSpecTemplateDefaultsInitContainersVolumeDevices `field:"optional" json:"volumeDevices" yaml:"volumeDevices"`
	VolumeMounts *[]*WorkflowSpecTemplateDefaultsInitContainersVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

