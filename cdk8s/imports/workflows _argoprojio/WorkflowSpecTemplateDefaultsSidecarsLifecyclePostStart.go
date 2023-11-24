package workflows _argoprojio


type WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStart struct {
	Exec *WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

