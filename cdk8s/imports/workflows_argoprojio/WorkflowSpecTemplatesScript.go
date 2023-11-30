package workflows_argoprojio


type WorkflowSpecTemplatesScript struct {
	Image *string `field:"required" json:"image" yaml:"image"`
	Source *string `field:"required" json:"source" yaml:"source"`
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	Env *[]*WorkflowSpecTemplatesScriptEnv `field:"optional" json:"env" yaml:"env"`
	EnvFrom *[]*WorkflowSpecTemplatesScriptEnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	Lifecycle *WorkflowSpecTemplatesScriptLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	LivenessProbe *WorkflowSpecTemplatesScriptLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Ports *[]*WorkflowSpecTemplatesScriptPorts `field:"optional" json:"ports" yaml:"ports"`
	ReadinessProbe *WorkflowSpecTemplatesScriptReadinessProbe `field:"optional" json:"readinessProbe" yaml:"readinessProbe"`
	Resources *WorkflowSpecTemplatesScriptResources `field:"optional" json:"resources" yaml:"resources"`
	SecurityContext *WorkflowSpecTemplatesScriptSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	StartupProbe *WorkflowSpecTemplatesScriptStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	Stdin *bool `field:"optional" json:"stdin" yaml:"stdin"`
	StdinOnce *bool `field:"optional" json:"stdinOnce" yaml:"stdinOnce"`
	TerminationMessagePath *string `field:"optional" json:"terminationMessagePath" yaml:"terminationMessagePath"`
	TerminationMessagePolicy *string `field:"optional" json:"terminationMessagePolicy" yaml:"terminationMessagePolicy"`
	Tty *bool `field:"optional" json:"tty" yaml:"tty"`
	VolumeDevices *[]*WorkflowSpecTemplatesScriptVolumeDevices `field:"optional" json:"volumeDevices" yaml:"volumeDevices"`
	VolumeMounts *[]*WorkflowSpecTemplatesScriptVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

