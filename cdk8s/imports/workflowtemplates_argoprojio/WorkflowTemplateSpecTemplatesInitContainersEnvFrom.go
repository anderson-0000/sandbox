package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesInitContainersEnvFrom struct {
	ConfigMapRef *WorkflowTemplateSpecTemplatesInitContainersEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowTemplateSpecTemplatesInitContainersEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

