package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvFrom struct {
	ConfigMapRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

