package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvFrom struct {
	ConfigMapRef *CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

