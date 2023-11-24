package workflows _argoprojio


type WorkflowSpecSynchronizationSemaphore struct {
	ConfigMapKeyRef *WorkflowSpecSynchronizationSemaphoreConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

