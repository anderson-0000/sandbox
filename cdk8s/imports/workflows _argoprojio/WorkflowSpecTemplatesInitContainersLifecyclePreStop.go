package workflows _argoprojio


type WorkflowSpecTemplatesInitContainersLifecyclePreStop struct {
	Exec *WorkflowSpecTemplatesInitContainersLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowSpecTemplatesInitContainersLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowSpecTemplatesInitContainersLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

