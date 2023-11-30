package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStop struct {
	Exec *CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *CronWorkflowSpecWorkflowSpecTemplatesSidecarsLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

