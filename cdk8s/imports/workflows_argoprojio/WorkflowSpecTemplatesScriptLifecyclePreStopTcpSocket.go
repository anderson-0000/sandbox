package workflows_argoprojio


type WorkflowSpecTemplatesScriptLifecyclePreStopTcpSocket struct {
	Port WorkflowSpecTemplatesScriptLifecyclePreStopTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

