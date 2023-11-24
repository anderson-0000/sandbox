package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbe struct {
	Exec *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeExec `field:"optional" json:"exec" yaml:"exec"`
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	Grpc *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeGrpc `field:"optional" json:"grpc" yaml:"grpc"`
	HttpGet *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	InitialDelaySeconds *float64 `field:"optional" json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	PeriodSeconds *float64 `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	TcpSocket *CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
	TerminationGracePeriodSeconds *float64 `field:"optional" json:"terminationGracePeriodSeconds" yaml:"terminationGracePeriodSeconds"`
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

