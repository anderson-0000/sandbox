package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStart struct {
	Exec *WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

