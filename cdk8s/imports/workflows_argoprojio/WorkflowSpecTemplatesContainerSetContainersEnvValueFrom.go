package workflows_argoprojio


type WorkflowSpecTemplatesContainerSetContainersEnvValueFrom struct {
	ConfigMapKeyRef *WorkflowSpecTemplatesContainerSetContainersEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *WorkflowSpecTemplatesContainerSetContainersEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *WorkflowSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *WorkflowSpecTemplatesContainerSetContainersEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

