package workflows_argoprojio


type WorkflowSpecTemplateDefaultsSynchronizationSemaphore struct {
	ConfigMapKeyRef *WorkflowSpecTemplateDefaultsSynchronizationSemaphoreConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

