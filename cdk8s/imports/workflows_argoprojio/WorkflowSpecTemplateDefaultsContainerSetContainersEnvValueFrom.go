package workflows_argoprojio


type WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFrom struct {
	ConfigMapKeyRef *WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *WorkflowSpecTemplateDefaultsContainerSetContainersEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

