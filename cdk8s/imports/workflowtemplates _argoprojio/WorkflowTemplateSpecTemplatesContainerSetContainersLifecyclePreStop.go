package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStop struct {
	Exec *WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

