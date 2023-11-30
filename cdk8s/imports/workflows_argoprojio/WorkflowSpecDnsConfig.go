package workflows_argoprojio


type WorkflowSpecDnsConfig struct {
	Nameservers *[]*string `field:"optional" json:"nameservers" yaml:"nameservers"`
	Options *[]*WorkflowSpecDnsConfigOptions `field:"optional" json:"options" yaml:"options"`
	Searches *[]*string `field:"optional" json:"searches" yaml:"searches"`
}

