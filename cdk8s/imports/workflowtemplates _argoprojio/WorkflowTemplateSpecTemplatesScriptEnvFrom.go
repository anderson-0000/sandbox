package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesScriptEnvFrom struct {
	ConfigMapRef *WorkflowTemplateSpecTemplatesScriptEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowTemplateSpecTemplatesScriptEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

