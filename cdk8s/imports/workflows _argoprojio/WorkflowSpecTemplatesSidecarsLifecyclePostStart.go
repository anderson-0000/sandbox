package workflows _argoprojio


type WorkflowSpecTemplatesSidecarsLifecyclePostStart struct {
	Exec *WorkflowSpecTemplatesSidecarsLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowSpecTemplatesSidecarsLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowSpecTemplatesSidecarsLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

