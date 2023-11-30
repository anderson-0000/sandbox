package workflows_argoprojio


type WorkflowSpecTemplatesInitContainersLifecycle struct {
	PostStart *WorkflowSpecTemplatesInitContainersLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *WorkflowSpecTemplatesInitContainersLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

