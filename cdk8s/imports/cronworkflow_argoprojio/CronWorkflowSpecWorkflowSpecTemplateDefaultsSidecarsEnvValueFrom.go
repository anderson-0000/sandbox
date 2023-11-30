package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFrom struct {
	ConfigMapKeyRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

