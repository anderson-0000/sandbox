package workflows _argoprojio


type WorkflowSpecTemplatesScriptLifecyclePostStartTcpSocket struct {
	Port WorkflowSpecTemplatesScriptLifecyclePostStartTcpSocketPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
}

