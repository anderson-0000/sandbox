package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFrom struct {
	ConfigMapKeyRef *CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

