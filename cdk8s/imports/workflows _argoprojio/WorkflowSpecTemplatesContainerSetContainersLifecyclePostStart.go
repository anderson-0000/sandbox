package workflows _argoprojio


type WorkflowSpecTemplatesContainerSetContainersLifecyclePostStart struct {
	Exec *WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

