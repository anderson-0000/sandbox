package workflows _argoprojio


type WorkflowSpecTemplatesSynchronizationSemaphore struct {
	ConfigMapKeyRef *WorkflowSpecTemplatesSynchronizationSemaphoreConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

