package workflows_argoprojio


type WorkflowSpecTemplateDefaultsScriptLifecycle struct {
	PostStart *WorkflowSpecTemplateDefaultsScriptLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *WorkflowSpecTemplateDefaultsScriptLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

