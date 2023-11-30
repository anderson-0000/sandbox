package workflows_argoprojio


type WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStop struct {
	Exec *WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

