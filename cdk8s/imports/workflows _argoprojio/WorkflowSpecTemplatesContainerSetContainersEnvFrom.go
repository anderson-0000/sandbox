package workflows _argoprojio


type WorkflowSpecTemplatesContainerSetContainersEnvFrom struct {
	ConfigMapRef *WorkflowSpecTemplatesContainerSetContainersEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowSpecTemplatesContainerSetContainersEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

