package workflows _argoprojio


type WorkflowSpecHostAliases struct {
	Hostnames *[]*string `field:"optional" json:"hostnames" yaml:"hostnames"`
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
}

