package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFrom struct {
	ConfigMapKeyRef *WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

