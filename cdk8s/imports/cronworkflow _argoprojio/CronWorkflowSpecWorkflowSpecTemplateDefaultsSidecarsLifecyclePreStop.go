package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStop struct {
	Exec *CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *CronWorkflowSpecWorkflowSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

