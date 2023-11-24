package workflows _argoprojio


type WorkflowSpecTemplateDefaultsContainerLifecycle struct {
	PostStart *WorkflowSpecTemplateDefaultsContainerLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *WorkflowSpecTemplateDefaultsContainerLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

