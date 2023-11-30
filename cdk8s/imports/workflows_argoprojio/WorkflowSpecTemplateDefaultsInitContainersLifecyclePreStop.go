package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStop struct {
	Exec *WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

