package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsContainer struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	Env *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnv `field:"optional" json:"env" yaml:"env"`
	EnvFrom *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	Image *string `field:"optional" json:"image" yaml:"image"`
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	Lifecycle *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	LivenessProbe *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	Ports *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerPorts `field:"optional" json:"ports" yaml:"ports"`
	ReadinessProbe *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerReadinessProbe `field:"optional" json:"readinessProbe" yaml:"readinessProbe"`
	Resources *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerResources `field:"optional" json:"resources" yaml:"resources"`
	SecurityContext *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	StartupProbe *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	Stdin *bool `field:"optional" json:"stdin" yaml:"stdin"`
	StdinOnce *bool `field:"optional" json:"stdinOnce" yaml:"stdinOnce"`
	TerminationMessagePath *string `field:"optional" json:"terminationMessagePath" yaml:"terminationMessagePath"`
	TerminationMessagePolicy *string `field:"optional" json:"terminationMessagePolicy" yaml:"terminationMessagePolicy"`
	Tty *bool `field:"optional" json:"tty" yaml:"tty"`
	VolumeDevices *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerVolumeDevices `field:"optional" json:"volumeDevices" yaml:"volumeDevices"`
	VolumeMounts *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

