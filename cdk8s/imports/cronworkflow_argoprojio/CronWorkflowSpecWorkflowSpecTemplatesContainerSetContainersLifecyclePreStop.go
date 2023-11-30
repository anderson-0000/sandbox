package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStop struct {
	Exec *CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

