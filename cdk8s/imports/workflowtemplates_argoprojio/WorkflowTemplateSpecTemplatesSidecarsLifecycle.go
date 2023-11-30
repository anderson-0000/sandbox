package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesSidecarsLifecycle struct {
	PostStart *WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

