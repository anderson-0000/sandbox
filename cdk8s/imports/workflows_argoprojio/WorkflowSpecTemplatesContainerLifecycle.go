package workflows_argoprojio


type WorkflowSpecTemplatesContainerLifecycle struct {
	PostStart *WorkflowSpecTemplatesContainerLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *WorkflowSpecTemplatesContainerLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

