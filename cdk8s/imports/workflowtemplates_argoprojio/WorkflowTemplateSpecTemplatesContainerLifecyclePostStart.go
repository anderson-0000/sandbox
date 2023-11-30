package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesContainerLifecyclePostStart struct {
	Exec *WorkflowTemplateSpecTemplatesContainerLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplatesContainerLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplatesContainerLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

