package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesContainerSetContainersEnvFrom struct {
	ConfigMapRef *WorkflowTemplateSpecTemplatesContainerSetContainersEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowTemplateSpecTemplatesContainerSetContainersEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

