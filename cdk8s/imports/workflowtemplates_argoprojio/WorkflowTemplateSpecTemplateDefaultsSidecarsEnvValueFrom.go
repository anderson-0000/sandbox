package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFrom struct {
	ConfigMapKeyRef *WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

