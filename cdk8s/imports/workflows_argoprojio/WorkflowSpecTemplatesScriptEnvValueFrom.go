package workflows_argoprojio


type WorkflowSpecTemplatesScriptEnvValueFrom struct {
	ConfigMapKeyRef *WorkflowSpecTemplatesScriptEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	FieldRef *WorkflowSpecTemplatesScriptEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	ResourceFieldRef *WorkflowSpecTemplatesScriptEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	SecretKeyRef *WorkflowSpecTemplatesScriptEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

