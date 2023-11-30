package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesContainerEnvFrom struct {
	ConfigMapRef *WorkflowTemplateSpecTemplatesContainerEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowTemplateSpecTemplatesContainerEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

