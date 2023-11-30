package workflows_argoprojio


type WorkflowSpecTemplateDefaultsSidecarsEnvValueFrom struct {
	ConfigMapKeyRef *WorkflowSpecTemplateDefaultsSidecarsEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *WorkflowSpecTemplateDefaultsSidecarsEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *WorkflowSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *WorkflowSpecTemplateDefaultsSidecarsEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

