package workflowtemplates_argoprojio


type WorkflowTemplateSpecSynchronization struct {
	Mutex *WorkflowTemplateSpecSynchronizationMutex `field:"optional" json:"mutex" yaml:"mutex"`
	Semaphore *WorkflowTemplateSpecSynchronizationSemaphore `field:"optional" json:"semaphore" yaml:"semaphore"`
}

