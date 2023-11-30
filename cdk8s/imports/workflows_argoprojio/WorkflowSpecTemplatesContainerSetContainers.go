package workflows_argoprojio


type WorkflowSpecTemplatesContainerSetContainers struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	Env *[]*WorkflowSpecTemplatesContainerSetContainersEnv `field:"optional" json:"env" yaml:"env"`
	EnvFrom *[]*WorkflowSpecTemplatesContainerSetContainersEnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	Image *string `field:"optional" json:"image" yaml:"image"`
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	Lifecycle *WorkflowSpecTemplatesContainerSetContainersLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	LivenessProbe *WorkflowSpecTemplatesContainerSetContainersLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	Ports *[]*WorkflowSpecTemplatesContainerSetContainersPorts `field:"optional" json:"ports" yaml:"ports"`
	ReadinessProbe *WorkflowSpecTemplatesContainerSetContainersReadinessProbe `field:"optional" json:"readinessProbe" yaml:"readinessProbe"`
	Resources *WorkflowSpecTemplatesContainerSetContainersResources `field:"optional" json:"resources" yaml:"resources"`
	SecurityContext *WorkflowSpecTemplatesContainerSetContainersSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	StartupProbe *WorkflowSpecTemplatesContainerSetContainersStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	Stdin *bool `field:"optional" json:"stdin" yaml:"stdin"`
	StdinOnce *bool `field:"optional" json:"stdinOnce" yaml:"stdinOnce"`
	TerminationMessagePath *string `field:"optional" json:"terminationMessagePath" yaml:"terminationMessagePath"`
	TerminationMessagePolicy *string `field:"optional" json:"terminationMessagePolicy" yaml:"terminationMessagePolicy"`
	Tty *bool `field:"optional" json:"tty" yaml:"tty"`
	VolumeDevices *[]*WorkflowSpecTemplatesContainerSetContainersVolumeDevices `field:"optional" json:"volumeDevices" yaml:"volumeDevices"`
	VolumeMounts *[]*WorkflowSpecTemplatesContainerSetContainersVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

