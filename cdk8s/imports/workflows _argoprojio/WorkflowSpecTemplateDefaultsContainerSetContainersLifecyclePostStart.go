package workflows _argoprojio


type WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStart struct {
	Exec *WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

