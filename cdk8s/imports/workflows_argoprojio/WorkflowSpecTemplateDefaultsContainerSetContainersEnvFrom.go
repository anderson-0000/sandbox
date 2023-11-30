package workflows_argoprojio


type WorkflowSpecTemplateDefaultsContainerSetContainersEnvFrom struct {
	ConfigMapRef *WorkflowSpecTemplateDefaultsContainerSetContainersEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowSpecTemplateDefaultsContainerSetContainersEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

