package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStop struct {
	Exec *WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

