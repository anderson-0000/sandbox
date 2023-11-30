package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesScript struct {
	Image *string `field:"required" json:"image" yaml:"image"`
	Source *string `field:"required" json:"source" yaml:"source"`
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	Env *[]*WorkflowTemplateSpecTemplatesScriptEnv `field:"optional" json:"env" yaml:"env"`
	EnvFrom *[]*WorkflowTemplateSpecTemplatesScriptEnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	Lifecycle *WorkflowTemplateSpecTemplatesScriptLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	LivenessProbe *WorkflowTemplateSpecTemplatesScriptLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Ports *[]*WorkflowTemplateSpecTemplatesScriptPorts `field:"optional" json:"ports" yaml:"ports"`
	ReadinessProbe *WorkflowTemplateSpecTemplatesScriptReadinessProbe `field:"optional" json:"readinessProbe" yaml:"readinessProbe"`
	Resources *WorkflowTemplateSpecTemplatesScriptResources `field:"optional" json:"resources" yaml:"resources"`
	SecurityContext *WorkflowTemplateSpecTemplatesScriptSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	StartupProbe *WorkflowTemplateSpecTemplatesScriptStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	Stdin *bool `field:"optional" json:"stdin" yaml:"stdin"`
	StdinOnce *bool `field:"optional" json:"stdinOnce" yaml:"stdinOnce"`
	TerminationMessagePath *string `field:"optional" json:"terminationMessagePath" yaml:"terminationMessagePath"`
	TerminationMessagePolicy *string `field:"optional" json:"terminationMessagePolicy" yaml:"terminationMessagePolicy"`
	Tty *bool `field:"optional" json:"tty" yaml:"tty"`
	VolumeDevices *[]*WorkflowTemplateSpecTemplatesScriptVolumeDevices `field:"optional" json:"volumeDevices" yaml:"volumeDevices"`
	VolumeMounts *[]*WorkflowTemplateSpecTemplatesScriptVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

