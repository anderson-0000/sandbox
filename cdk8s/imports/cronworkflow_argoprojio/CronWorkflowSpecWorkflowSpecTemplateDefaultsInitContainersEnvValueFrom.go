package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFrom struct {
	ConfigMapKeyRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

