package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainerLifecycle struct {
	PostStart *WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

