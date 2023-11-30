package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainerEnvFrom struct {
	ConfigMapRef *WorkflowTemplateSpecTemplateDefaultsContainerEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowTemplateSpecTemplateDefaultsContainerEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

