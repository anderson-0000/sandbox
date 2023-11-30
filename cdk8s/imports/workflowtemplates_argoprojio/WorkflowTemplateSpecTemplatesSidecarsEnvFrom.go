package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesSidecarsEnvFrom struct {
	ConfigMapRef *WorkflowTemplateSpecTemplatesSidecarsEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowTemplateSpecTemplatesSidecarsEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

