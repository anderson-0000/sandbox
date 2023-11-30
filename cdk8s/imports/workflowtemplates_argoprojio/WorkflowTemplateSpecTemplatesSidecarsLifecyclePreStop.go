package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStop struct {
	Exec *WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

