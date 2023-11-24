package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumes struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	AwsElasticBlockStore *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	AzureDisk *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	AzureFile *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	Cephfs *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCephfs `field:"optional" json:"cephfs" yaml:"cephfs"`
	Cinder *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCinder `field:"optional" json:"cinder" yaml:"cinder"`
	ConfigMap *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	Csi *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesCsi `field:"optional" json:"csi" yaml:"csi"`
	DownwardApi *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	EmptyDir *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	Ephemeral *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesEphemeral `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	Fc *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesFc `field:"optional" json:"fc" yaml:"fc"`
	FlexVolume *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	Flocker *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	GcePersistentDisk *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	GitRepo *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesGitRepo `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	Glusterfs *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	HostPath *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	Iscsi *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	Nfs *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	PersistentVolumeClaim *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesPersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	PhotonPersistentDisk *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesPhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	PortworxVolume *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesPortworxVolume `field:"optional" json:"portworxVolume" yaml:"portworxVolume"`
	Projected *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesProjected `field:"optional" json:"projected" yaml:"projected"`
	Quobyte *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	Rbd *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesRbd `field:"optional" json:"rbd" yaml:"rbd"`
	ScaleIo *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesScaleIo `field:"optional" json:"scaleIo" yaml:"scaleIo"`
	Secret *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
	Storageos *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesStorageos `field:"optional" json:"storageos" yaml:"storageos"`
	VsphereVolume *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

