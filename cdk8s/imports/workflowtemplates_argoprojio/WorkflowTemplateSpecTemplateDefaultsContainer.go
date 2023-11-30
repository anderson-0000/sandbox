package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainer struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	Env *[]*WorkflowTemplateSpecTemplateDefaultsContainerEnv `field:"optional" json:"env" yaml:"env"`
	EnvFrom *[]*WorkflowTemplateSpecTemplateDefaultsContainerEnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	Image *string `field:"optional" json:"image" yaml:"image"`
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	Lifecycle *WorkflowTemplateSpecTemplateDefaultsContainerLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	LivenessProbe *WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	Ports *[]*WorkflowTemplateSpecTemplateDefaultsContainerPorts `field:"optional" json:"ports" yaml:"ports"`
	ReadinessProbe *WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbe `field:"optional" json:"readinessProbe" yaml:"readinessProbe"`
	Resources *WorkflowTemplateSpecTemplateDefaultsContainerResources `field:"optional" json:"resources" yaml:"resources"`
	SecurityContext *WorkflowTemplateSpecTemplateDefaultsContainerSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	StartupProbe *WorkflowTemplateSpecTemplateDefaultsContainerStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	Stdin *bool `field:"optional" json:"stdin" yaml:"stdin"`
	StdinOnce *bool `field:"optional" json:"stdinOnce" yaml:"stdinOnce"`
	TerminationMessagePath *string `field:"optional" json:"terminationMessagePath" yaml:"terminationMessagePath"`
	TerminationMessagePolicy *string `field:"optional" json:"terminationMessagePolicy" yaml:"terminationMessagePolicy"`
	Tty *bool `field:"optional" json:"tty" yaml:"tty"`
	VolumeDevices *[]*WorkflowTemplateSpecTemplateDefaultsContainerVolumeDevices `field:"optional" json:"volumeDevices" yaml:"volumeDevices"`
	VolumeMounts *[]*WorkflowTemplateSpecTemplateDefaultsContainerVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

