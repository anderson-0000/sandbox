package workflows_argoprojio


type WorkflowSpecTemplatesScriptLifecycle struct {
	PostStart *WorkflowSpecTemplatesScriptLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *WorkflowSpecTemplatesScriptLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

