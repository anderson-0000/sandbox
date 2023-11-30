package workflows_argoprojio


type WorkflowSpecTemplatesContainerSetContainersLifecyclePreStop struct {
	Exec *WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

