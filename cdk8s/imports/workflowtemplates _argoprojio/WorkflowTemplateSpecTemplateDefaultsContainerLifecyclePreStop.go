package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStop struct {
	Exec *WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

