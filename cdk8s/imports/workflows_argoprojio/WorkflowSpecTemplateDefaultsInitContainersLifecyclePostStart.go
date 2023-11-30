package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStart struct {
	Exec *WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

