package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStart struct {
	Exec *WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

