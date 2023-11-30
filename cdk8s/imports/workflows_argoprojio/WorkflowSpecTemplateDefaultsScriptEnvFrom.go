package workflows_argoprojio


type WorkflowSpecTemplateDefaultsScriptEnvFrom struct {
	ConfigMapRef *WorkflowSpecTemplateDefaultsScriptEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowSpecTemplateDefaultsScriptEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

