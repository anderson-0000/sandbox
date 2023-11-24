package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesContainerLifecycle struct {
	PostStart *WorkflowTemplateSpecTemplatesContainerLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *WorkflowTemplateSpecTemplatesContainerLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

