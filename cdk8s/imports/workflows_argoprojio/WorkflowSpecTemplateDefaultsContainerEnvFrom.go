package workflows_argoprojio


type WorkflowSpecTemplateDefaultsContainerEnvFrom struct {
	ConfigMapRef *WorkflowSpecTemplateDefaultsContainerEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowSpecTemplateDefaultsContainerEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

