package workflows_argoprojio


type WorkflowSpecTemplatesSidecarsLifecyclePreStop struct {
	Exec *WorkflowSpecTemplatesSidecarsLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

