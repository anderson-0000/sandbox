package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesContainerEnvFrom struct {
	ConfigMapRef *CronWorkflowSpecWorkflowSpecTemplatesContainerEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *CronWorkflowSpecWorkflowSpecTemplatesContainerEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

