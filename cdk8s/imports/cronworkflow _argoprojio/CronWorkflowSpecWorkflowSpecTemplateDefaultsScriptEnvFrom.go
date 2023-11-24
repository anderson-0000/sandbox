package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvFrom struct {
	ConfigMapRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

