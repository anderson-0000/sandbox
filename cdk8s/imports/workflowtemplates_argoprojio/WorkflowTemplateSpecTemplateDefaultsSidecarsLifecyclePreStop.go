package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStop struct {
	Exec *WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

