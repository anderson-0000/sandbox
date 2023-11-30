package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvFrom struct {
	ConfigMapRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

