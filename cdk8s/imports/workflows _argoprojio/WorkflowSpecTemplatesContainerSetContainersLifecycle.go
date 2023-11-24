package workflows _argoprojio


type WorkflowSpecTemplatesContainerSetContainersLifecycle struct {
	PostStart *WorkflowSpecTemplatesContainerSetContainersLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *WorkflowSpecTemplatesContainerSetContainersLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

