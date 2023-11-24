package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvFrom struct {
	ConfigMapRef *CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

