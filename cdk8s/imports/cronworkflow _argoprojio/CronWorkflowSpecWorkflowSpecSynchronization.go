package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecSynchronization struct {
	Mutex *CronWorkflowSpecWorkflowSpecSynchronizationMutex `field:"optional" json:"mutex" yaml:"mutex"`
	Semaphore *CronWorkflowSpecWorkflowSpecSynchronizationSemaphore `field:"optional" json:"semaphore" yaml:"semaphore"`
}

