package workflows _argoprojio


type WorkflowSpecTemplateDefaultsSynchronization struct {
	Mutex *WorkflowSpecTemplateDefaultsSynchronizationMutex `field:"optional" json:"mutex" yaml:"mutex"`
	Semaphore *WorkflowSpecTemplateDefaultsSynchronizationSemaphore `field:"optional" json:"semaphore" yaml:"semaphore"`
}

