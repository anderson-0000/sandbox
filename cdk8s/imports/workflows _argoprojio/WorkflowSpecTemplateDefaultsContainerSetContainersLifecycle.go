package workflows _argoprojio


type WorkflowSpecTemplateDefaultsContainerSetContainersLifecycle struct {
	PostStart *WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

