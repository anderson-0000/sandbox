package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFrom struct {
	ConfigMapKeyRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

