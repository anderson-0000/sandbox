package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStop struct {
	Exec *WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

