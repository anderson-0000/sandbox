package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStart struct {
	Exec *WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

