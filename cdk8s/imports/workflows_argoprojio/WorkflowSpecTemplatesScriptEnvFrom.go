package workflows_argoprojio


type WorkflowSpecTemplatesScriptEnvFrom struct {
	ConfigMapRef *WorkflowSpecTemplatesScriptEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *WorkflowSpecTemplatesScriptEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

