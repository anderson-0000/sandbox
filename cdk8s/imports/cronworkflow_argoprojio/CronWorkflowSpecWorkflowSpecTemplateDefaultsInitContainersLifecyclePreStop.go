package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStop struct {
	Exec *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

