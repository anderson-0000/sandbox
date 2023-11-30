package workflows_argoprojio


type WorkflowSpecTemplatesContainerEnvValueFrom struct {
	ConfigMapKeyRef *WorkflowSpecTemplatesContainerEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *WorkflowSpecTemplatesContainerEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *WorkflowSpecTemplatesContainerEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *WorkflowSpecTemplatesContainerEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

