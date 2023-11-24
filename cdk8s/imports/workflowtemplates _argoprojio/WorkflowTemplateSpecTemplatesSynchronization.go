package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesSynchronization struct {
	Mutex *WorkflowTemplateSpecTemplatesSynchronizationMutex `field:"optional" json:"mutex" yaml:"mutex"`
	Semaphore *WorkflowTemplateSpecTemplatesSynchronizationSemaphore `field:"optional" json:"semaphore" yaml:"semaphore"`
}

