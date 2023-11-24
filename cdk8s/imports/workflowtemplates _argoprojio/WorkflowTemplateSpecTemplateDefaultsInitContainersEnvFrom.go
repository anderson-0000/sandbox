package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsInitContainersEnvFrom struct {
	ConfigMapRef *WorkflowTemplateSpecTemplateDefaultsInitContainersEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowTemplateSpecTemplateDefaultsInitContainersEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

