package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFrom struct {
	ConfigMapKeyRef *CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *CronWorkflowSpecWorkflowSpecTemplatesInitContainersEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

