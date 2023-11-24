package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsScriptEnvFrom struct {
	ConfigMapRef *WorkflowTemplateSpecTemplateDefaultsScriptEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowTemplateSpecTemplateDefaultsScriptEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

