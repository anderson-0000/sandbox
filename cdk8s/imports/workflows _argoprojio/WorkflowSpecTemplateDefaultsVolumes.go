package workflows _argoprojio


type WorkflowSpecTemplateDefaultsVolumes struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	AwsElasticBlockStore *WorkflowSpecTemplateDefaultsVolumesAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	AzureDisk *WorkflowSpecTemplateDefaultsVolumesAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	AzureFile *WorkflowSpecTemplateDefaultsVolumesAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	Cephfs *WorkflowSpecTemplateDefaultsVolumesCephfs `field:"optional" json:"cephfs" yaml:"cephfs"`
	Cinder *WorkflowSpecTemplateDefaultsVolumesCinder `field:"optional" json:"cinder" yaml:"cinder"`
	ConfigMap *WorkflowSpecTemplateDefaultsVolumesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	Csi *WorkflowSpecTemplateDefaultsVolumesCsi `field:"optional" json:"csi" yaml:"csi"`
	DownwardApi *WorkflowSpecTemplateDefaultsVolumesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	EmptyDir *WorkflowSpecTemplateDefaultsVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	Ephemeral *WorkflowSpecTemplateDefaultsVolumesEphemeral `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	Fc *WorkflowSpecTemplateDefaultsVolumesFc `field:"optional" json:"fc" yaml:"fc"`
	FlexVolume *WorkflowSpecTemplateDefaultsVolumesFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	Flocker *WorkflowSpecTemplateDefaultsVolumesFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	GcePersistentDisk *WorkflowSpecTemplateDefaultsVolumesGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	GitRepo *WorkflowSpecTemplateDefaultsVolumesGitRepo `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	Glusterfs *WorkflowSpecTemplateDefaultsVolumesGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	HostPath *WorkflowSpecTemplateDefaultsVolumesHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	Iscsi *WorkflowSpecTemplateDefaultsVolumesIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	Nfs *WorkflowSpecTemplateDefaultsVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	PersistentVolumeClaim *WorkflowSpecTemplateDefaultsVolumesPersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	PhotonPersistentDisk *WorkflowSpecTemplateDefaultsVolumesPhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	PortworxVolume *WorkflowSpecTemplateDefaultsVolumesPortworxVolume `field:"optional" json:"portworxVolume" yaml:"portworxVolume"`
	Projected *WorkflowSpecTemplateDefaultsVolumesProjected `field:"optional" json:"projected" yaml:"projected"`
	Quobyte *WorkflowSpecTemplateDefaultsVolumesQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	Rbd *WorkflowSpecTemplateDefaultsVolumesRbd `field:"optional" json:"rbd" yaml:"rbd"`
	ScaleIo *WorkflowSpecTemplateDefaultsVolumesScaleIo `field:"optional" json:"scaleIo" yaml:"scaleIo"`
	Secret *WorkflowSpecTemplateDefaultsVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
	Storageos *WorkflowSpecTemplateDefaultsVolumesStorageos `field:"optional" json:"storageos" yaml:"storageos"`
	VsphereVolume *WorkflowSpecTemplateDefaultsVolumesVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

