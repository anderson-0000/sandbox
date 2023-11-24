package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesSynchronization struct {
	Mutex *CronWorkflowSpecWorkflowSpecTemplatesSynchronizationMutex `field:"optional" json:"mutex" yaml:"mutex"`
	Semaphore *CronWorkflowSpecWorkflowSpecTemplatesSynchronizationSemaphore `field:"optional" json:"semaphore" yaml:"semaphore"`
}

