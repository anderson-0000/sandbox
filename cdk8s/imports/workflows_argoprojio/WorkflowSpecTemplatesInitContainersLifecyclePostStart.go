package workflows_argoprojio


type WorkflowSpecTemplatesInitContainersLifecyclePostStart struct {
	Exec *WorkflowSpecTemplatesInitContainersLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowSpecTemplatesInitContainersLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowSpecTemplatesInitContainersLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

