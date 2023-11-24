package workflows _argoprojio


type WorkflowSpecTemplateDefaultsScriptLifecycle struct {
	PostStart *WorkflowSpecTemplateDefaultsScriptLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *WorkflowSpecTemplateDefaultsScriptLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

