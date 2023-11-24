package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStop struct {
	Exec *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

