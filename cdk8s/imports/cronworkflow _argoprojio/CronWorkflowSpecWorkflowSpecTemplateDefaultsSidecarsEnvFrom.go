package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvFrom struct {
	ConfigMapRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	SecretRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

