package workflows_argoprojio


type WorkflowSpecTemplatesContainerEnvFrom struct {
	ConfigMapRef *WorkflowSpecTemplatesContainerEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowSpecTemplatesContainerEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

