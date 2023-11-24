package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesScriptEnvFrom struct {
	ConfigMapRef *CronWorkflowSpecWorkflowSpecTemplatesScriptEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *CronWorkflowSpecWorkflowSpecTemplatesScriptEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

