package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInitContainersEnvFrom struct {
	ConfigMapRef *WorkflowSpecTemplateDefaultsInitContainersEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowSpecTemplateDefaultsInitContainersEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

