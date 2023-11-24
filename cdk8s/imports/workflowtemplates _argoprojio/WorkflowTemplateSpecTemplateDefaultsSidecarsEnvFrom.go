package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsSidecarsEnvFrom struct {
	ConfigMapRef *WorkflowTemplateSpecTemplateDefaultsSidecarsEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowTemplateSpecTemplateDefaultsSidecarsEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

