package workflows _argoprojio


type WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStop struct {
	Exec *WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

