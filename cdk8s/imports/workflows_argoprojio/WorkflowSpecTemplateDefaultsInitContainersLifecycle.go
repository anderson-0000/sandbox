package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInitContainersLifecycle struct {
	PostStart *WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

