package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvFrom struct {
	ConfigMapRef *CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *CronWorkflowSpecWorkflowSpecTemplatesSidecarsEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

