package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStart struct {
	Exec *CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

