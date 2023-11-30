package workflowtemplates_argoprojio


type WorkflowTemplateSpecDnsConfig struct {
	Nameservers *[]*string `field:"optional" json:"nameservers" yaml:"nameservers"`
	Options *[]*WorkflowTemplateSpecDnsConfigOptions `field:"optional" json:"options" yaml:"options"`
	Searches *[]*string `field:"optional" json:"searches" yaml:"searches"`
}

