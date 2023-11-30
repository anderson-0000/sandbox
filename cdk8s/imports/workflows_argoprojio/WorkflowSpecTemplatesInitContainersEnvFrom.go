package workflows_argoprojio


type WorkflowSpecTemplatesInitContainersEnvFrom struct {
	ConfigMapRef *WorkflowSpecTemplatesInitContainersEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowSpecTemplatesInitContainersEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

