package workflows_argoprojio


type WorkflowSpecTemplatesContainer struct {
	Image *string `field:"required" json:"image" yaml:"image"`
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	Env *[]*WorkflowSpecTemplatesContainerEnv `field:"optional" json:"env" yaml:"env"`
	EnvFrom *[]*WorkflowSpecTemplatesContainerEnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	Lifecycle *WorkflowSpecTemplatesContainerLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	LivenessProbe *WorkflowSpecTemplatesContainerLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Ports *[]*WorkflowSpecTemplatesContainerPorts `field:"optional" json:"ports" yaml:"ports"`
	ReadinessProbe *WorkflowSpecTemplatesContainerReadinessProbe `field:"optional" json:"readinessProbe" yaml:"readinessProbe"`
	Resources *WorkflowSpecTemplatesContainerResources `field:"optional" json:"resources" yaml:"resources"`
	SecurityContext *WorkflowSpecTemplatesContainerSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	StartupProbe *WorkflowSpecTemplatesContainerStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	Stdin *bool `field:"optional" json:"stdin" yaml:"stdin"`
	StdinOnce *bool `field:"optional" json:"stdinOnce" yaml:"stdinOnce"`
	TerminationMessagePath *string `field:"optional" json:"terminationMessagePath" yaml:"terminationMessagePath"`
	TerminationMessagePolicy *string `field:"optional" json:"terminationMessagePolicy" yaml:"terminationMessagePolicy"`
	Tty *bool `field:"optional" json:"tty" yaml:"tty"`
	VolumeDevices *[]*WorkflowSpecTemplatesContainerVolumeDevices `field:"optional" json:"volumeDevices" yaml:"volumeDevices"`
	VolumeMounts *[]*WorkflowSpecTemplatesContainerVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

