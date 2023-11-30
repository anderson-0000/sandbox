package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesContainerLifecycle struct {
	PostStart *CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStart `field:"optional" json:"postStart" yaml:"postStart"`
	PreStop *CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStop `field:"optional" json:"preStop" yaml:"preStop"`
}

