package workflows _argoprojio


type WorkflowSpecSynchronization struct {
	Mutex *WorkflowSpecSynchronizationMutex `field:"optional" json:"mutex" yaml:"mutex"`
	Semaphore *WorkflowSpecSynchronizationSemaphore `field:"optional" json:"semaphore" yaml:"semaphore"`
}

