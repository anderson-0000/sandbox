package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsInitContainers struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	Env *[]*WorkflowTemplateSpecTemplateDefaultsInitContainersEnv `field:"optional" json:"env" yaml:"env"`
	EnvFrom *[]*WorkflowTemplateSpecTemplateDefaultsInitContainersEnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	Image *string `field:"optional" json:"image" yaml:"image"`
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	Lifecycle *WorkflowTemplateSpecTemplateDefaultsInitContainersLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	LivenessProbe *WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	MirrorVolumeMounts *bool `field:"optional" json:"mirrorVolumeMounts" yaml:"mirrorVolumeMounts"`
	Ports *[]*WorkflowTemplateSpecTemplateDefaultsInitContainersPorts `field:"optional" json:"ports" yaml:"ports"`
	ReadinessProbe *WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbe `field:"optional" json:"readinessProbe" yaml:"readinessProbe"`
	Resources *WorkflowTemplateSpecTemplateDefaultsInitContainersResources `field:"optional" json:"resources" yaml:"resources"`
	SecurityContext *WorkflowTemplateSpecTemplateDefaultsInitContainersSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	StartupProbe *WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	Stdin *bool `field:"optional" json:"stdin" yaml:"stdin"`
	StdinOnce *bool `field:"optional" json:"stdinOnce" yaml:"stdinOnce"`
	TerminationMessagePath *string `field:"optional" json:"terminationMessagePath" yaml:"terminationMessagePath"`
	TerminationMessagePolicy *string `field:"optional" json:"terminationMessagePolicy" yaml:"terminationMessagePolicy"`
	Tty *bool `field:"optional" json:"tty" yaml:"tty"`
	VolumeDevices *[]*WorkflowTemplateSpecTemplateDefaultsInitContainersVolumeDevices `field:"optional" json:"volumeDevices" yaml:"volumeDevices"`
	VolumeMounts *[]*WorkflowTemplateSpecTemplateDefaultsInitContainersVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

