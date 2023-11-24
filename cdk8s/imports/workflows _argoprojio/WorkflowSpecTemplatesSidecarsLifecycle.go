package workflows _argoprojio


type WorkflowSpecTemplatesSidecarsLifecycle struct {
	PostStart *WorkflowSpecTemplatesSidecarsLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *WorkflowSpecTemplatesSidecarsLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

