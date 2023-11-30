package workflows_argoprojio


type WorkflowSpecSynchronizationMutex struct {
	Name *string `field:"optional" json:"name" yaml:"name"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

