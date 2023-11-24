package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsSynchronization struct {
	Mutex *WorkflowTemplateSpecTemplateDefaultsSynchronizationMutex `field:"optional" json:"mutex" yaml:"mutex"`
	Semaphore *WorkflowTemplateSpecTemplateDefaultsSynchronizationSemaphore `field:"optional" json:"semaphore" yaml:"semaphore"`
}

