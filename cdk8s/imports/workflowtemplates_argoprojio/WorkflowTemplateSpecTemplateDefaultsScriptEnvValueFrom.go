package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFrom struct {
	ConfigMapKeyRef *WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

