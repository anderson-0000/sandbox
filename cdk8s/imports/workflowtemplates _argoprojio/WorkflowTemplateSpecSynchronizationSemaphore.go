package workflowtemplates _argoprojio


type WorkflowTemplateSpecSynchronizationSemaphore struct {
	ConfigMapKeyRef *WorkflowTemplateSpecSynchronizationSemaphoreConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

