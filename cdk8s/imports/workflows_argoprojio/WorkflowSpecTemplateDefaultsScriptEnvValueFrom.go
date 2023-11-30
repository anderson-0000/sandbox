package workflows_argoprojio


type WorkflowSpecTemplateDefaultsScriptEnvValueFrom struct {
	ConfigMapKeyRef *WorkflowSpecTemplateDefaultsScriptEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *WorkflowSpecTemplateDefaultsScriptEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *WorkflowSpecTemplateDefaultsScriptEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *WorkflowSpecTemplateDefaultsScriptEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

