package workflows _argoprojio


type WorkflowSpecTemplateDefaultsSidecarsEnvFrom struct {
	ConfigMapRef *WorkflowSpecTemplateDefaultsSidecarsEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowSpecTemplateDefaultsSidecarsEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

