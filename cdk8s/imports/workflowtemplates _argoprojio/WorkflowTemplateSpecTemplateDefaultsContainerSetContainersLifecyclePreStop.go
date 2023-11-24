package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStop struct {
	Exec *WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

