package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvFrom struct {
	ConfigMapRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

