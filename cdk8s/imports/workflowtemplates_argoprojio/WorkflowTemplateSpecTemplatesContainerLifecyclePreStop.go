package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesContainerLifecyclePreStop struct {
	Exec *WorkflowTemplateSpecTemplatesContainerLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplatesContainerLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

