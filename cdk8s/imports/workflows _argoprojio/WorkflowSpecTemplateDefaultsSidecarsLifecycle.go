package workflows _argoprojio


type WorkflowSpecTemplateDefaultsSidecarsLifecycle struct {
	PostStart *WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

