package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFrom struct {
	ConfigMapKeyRef *WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

