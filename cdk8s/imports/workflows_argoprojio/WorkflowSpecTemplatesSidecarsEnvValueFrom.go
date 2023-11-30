package workflows_argoprojio


type WorkflowSpecTemplatesSidecarsEnvValueFrom struct {
	ConfigMapKeyRef *WorkflowSpecTemplatesSidecarsEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *WorkflowSpecTemplatesSidecarsEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *WorkflowSpecTemplatesSidecarsEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *WorkflowSpecTemplatesSidecarsEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

