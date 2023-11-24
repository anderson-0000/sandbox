package workflows _argoprojio


type WorkflowSpecTemplatesSidecarsEnvFrom struct {
	ConfigMapRef *WorkflowSpecTemplatesSidecarsEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowSpecTemplatesSidecarsEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

