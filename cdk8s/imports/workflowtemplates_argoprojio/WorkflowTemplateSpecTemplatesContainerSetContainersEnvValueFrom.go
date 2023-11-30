package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFrom struct {
	ConfigMapKeyRef *WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

