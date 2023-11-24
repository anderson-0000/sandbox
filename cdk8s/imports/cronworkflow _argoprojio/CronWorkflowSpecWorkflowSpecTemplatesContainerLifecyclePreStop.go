package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStop struct {
	Exec *CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

