package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecDnsConfig struct {
	Nameservers *[]*string `field:"optional" json:"nameservers" yaml:"nameservers"`
	Options *[]*CronWorkflowSpecWorkflowSpecDnsConfigOptions `field:"optional" json:"options" yaml:"options"`
	Searches *[]*string `field:"optional" json:"searches" yaml:"searches"`
}

