package workflows_argoprojio


type WorkflowSpecTemplatesSynchronization struct {
	Mutex *WorkflowSpecTemplatesSynchronizationMutex `field:"optional" json:"mutex" yaml:"mutex"`
	Semaphore *WorkflowSpecTemplatesSynchronizationSemaphore `field:"optional" json:"semaphore" yaml:"semaphore"`
}

