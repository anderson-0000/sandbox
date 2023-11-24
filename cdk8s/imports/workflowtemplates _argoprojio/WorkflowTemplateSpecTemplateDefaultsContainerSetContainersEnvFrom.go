package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvFrom struct {
	ConfigMapRef *WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

