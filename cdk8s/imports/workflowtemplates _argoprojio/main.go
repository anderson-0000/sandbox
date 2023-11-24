// workflowtemplates _argoprojio
package workflowtemplates _argoprojio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplate",
		reflect.TypeOf((*WorkflowTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_WorkflowTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateProps",
		reflect.TypeOf((*WorkflowTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpec",
		reflect.TypeOf((*WorkflowTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityNodeAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAntiAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArguments",
		reflect.TypeOf((*WorkflowTemplateSpecArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifacts",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsArchive",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsAzure",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsGcs",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsGit",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHttp",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsOss",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsRaw",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsS3",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsParameters",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsParametersValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecArtifactRepositoryRef",
		reflect.TypeOf((*WorkflowTemplateSpecArtifactRepositoryRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecDnsConfig",
		reflect.TypeOf((*WorkflowTemplateSpecDnsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecDnsConfigOptions",
		reflect.TypeOf((*WorkflowTemplateSpecDnsConfigOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecExecutor",
		reflect.TypeOf((*WorkflowTemplateSpecExecutor)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooks",
		reflect.TypeOf((*WorkflowTemplateSpecHooks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArguments",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifacts",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsArchive",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecHooksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsAzure",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsGcs",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsGit",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHttp",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsOss",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsRaw",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsS3",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsParameters",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsParametersValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecHooksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHooksTemplateRef",
		reflect.TypeOf((*WorkflowTemplateSpecHooksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecHostAliases",
		reflect.TypeOf((*WorkflowTemplateSpecHostAliases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecImagePullSecrets",
		reflect.TypeOf((*WorkflowTemplateSpecImagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecMetrics",
		reflect.TypeOf((*WorkflowTemplateSpecMetrics)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecMetricsPrometheus",
		reflect.TypeOf((*WorkflowTemplateSpecMetricsPrometheus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecMetricsPrometheusCounter",
		reflect.TypeOf((*WorkflowTemplateSpecMetricsPrometheusCounter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecMetricsPrometheusGauge",
		reflect.TypeOf((*WorkflowTemplateSpecMetricsPrometheusGauge)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecMetricsPrometheusHistogram",
		reflect.TypeOf((*WorkflowTemplateSpecMetricsPrometheusHistogram)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecMetricsPrometheusLabels",
		reflect.TypeOf((*WorkflowTemplateSpecMetricsPrometheusLabels)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecPodDisruptionBudget",
		reflect.TypeOf((*WorkflowTemplateSpecPodDisruptionBudget)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable",
		reflect.TypeOf((*WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecPodDisruptionBudgetMaxUnavailable{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecPodDisruptionBudgetMinAvailable",
		reflect.TypeOf((*WorkflowTemplateSpecPodDisruptionBudgetMinAvailable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecPodDisruptionBudgetMinAvailable{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecPodDisruptionBudgetSelector",
		reflect.TypeOf((*WorkflowTemplateSpecPodDisruptionBudgetSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecPodDisruptionBudgetSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecPodDisruptionBudgetSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecPodGc",
		reflect.TypeOf((*WorkflowTemplateSpecPodGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecPodGcLabelSelector",
		reflect.TypeOf((*WorkflowTemplateSpecPodGcLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecPodGcLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecPodGcLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecRetryStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecRetryStrategyAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecRetryStrategyAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecRetryStrategyBackoff",
		reflect.TypeOf((*WorkflowTemplateSpecRetryStrategyBackoff)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecRetryStrategyBackoffFactor",
		reflect.TypeOf((*WorkflowTemplateSpecRetryStrategyBackoffFactor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecRetryStrategyBackoffFactor{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecRetryStrategyLimit",
		reflect.TypeOf((*WorkflowTemplateSpecRetryStrategyLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecRetryStrategyLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecSecurityContext",
		reflect.TypeOf((*WorkflowTemplateSpecSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowTemplateSpecSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowTemplateSpecSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecSecurityContextSysctls",
		reflect.TypeOf((*WorkflowTemplateSpecSecurityContextSysctls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowTemplateSpecSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecSynchronization",
		reflect.TypeOf((*WorkflowTemplateSpecSynchronization)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecSynchronizationMutex",
		reflect.TypeOf((*WorkflowTemplateSpecSynchronizationMutex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecSynchronizationSemaphore",
		reflect.TypeOf((*WorkflowTemplateSpecSynchronizationSemaphore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecSynchronizationSemaphoreConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecSynchronizationSemaphoreConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaults",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaults)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsActiveDeadlineSeconds{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocation",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationAzure",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationGcs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationGit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationOss",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationRaw",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsArchiveLocationS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainer",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerEnv",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerEnvFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerEnvFromSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecycle",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStart",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStop",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerPorts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerResources",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerResourcesRequests",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSecurityContext",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainers",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnv",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvFromSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecycle",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStart",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStop",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersPorts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResources",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesLimits",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContext",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersVolumeDevices",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetContainersVolumeMounts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategyRetries{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerSetVolumeMounts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerSetVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerStartupProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsContainerStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerVolumeDevices",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsContainerVolumeMounts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsContainerVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDag",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDag)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasks",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArguments",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifacts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArchive",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsAzure",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsGcs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsGit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsOss",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsRaw",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsS3",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsParameters",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsParametersValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksContinueOn",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksContinueOn)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooks",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArguments",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifacts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArchive",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsAzure",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGcs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOss",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsRaw",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsParameters",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsParametersValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksHooksTemplateRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksHooksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksTemplateRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequence",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequence)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceCount{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceEnd{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsDagTasksWithSequenceStart{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsData",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSource",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPaths",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPaths)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArchive",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArchiveTar",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsAzure",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsGcs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsGit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsOss",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsRaw",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsS3",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataSourceArtifactPathsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsDataTransformation",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsDataTransformation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsExecutor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsExecutor)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsHostAliases",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsHostAliases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsHttpBodyFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsHttpBodyFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsHttpHeadersValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsHttpHeadersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsHttpHeadersValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsHttpHeadersValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainers",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersEnv",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersEnvFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersEnvFromSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersEnvValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecycle",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStart",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStop",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersPorts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersResources",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesRequests",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersSecurityContext",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsInitContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersVolumeDevices",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInitContainersVolumeMounts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInitContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifacts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArchive",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsAzure",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsGcs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsGit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsOss",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsRaw",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsS3",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsParameters",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsParametersValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsInputsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsInputsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsMemoize",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsMemoize)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsMemoizeCache",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsMemoizeCache)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsMemoizeCacheConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsMemoizeCacheConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsMetrics",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsMetrics)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsMetricsPrometheus",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsMetricsPrometheus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsMetricsPrometheusCounter",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsMetricsPrometheusCounter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsMetricsPrometheusGauge",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsMetricsPrometheusGauge)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsMetricsPrometheusHistogram",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsMetricsPrometheusHistogram)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsMetricsPrometheusLabels",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsMetricsPrometheusLabels)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifacts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArchive",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsAzure",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsGcs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsGit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsOss",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsRaw",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsParameters",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsParametersValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsOutputsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsOutputsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResource",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifact",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifact)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArchive",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArchiveTar",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactAzure",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactGcs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactGit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactOss",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactRaw",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactS3",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsResourceManifestFromArtifactS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsRetryStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsRetryStrategyAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsRetryStrategyAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoff",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoff)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsRetryStrategyBackoffFactor{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsRetryStrategyLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScript",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScript)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptEnv",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptEnvFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptEnvFromSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecycle",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStart",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStop",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptPorts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptResources",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptSecurityContext",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptStartupProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsScriptStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptVolumeDevices",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsScriptVolumeMounts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsScriptVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSecurityContext",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSecurityContextSysctls",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSecurityContextSysctls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecars",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecars)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsEnv",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsEnvFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsEnvFromSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecycle",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStart",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStop",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsPorts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsResources",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesRequests",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsSecurityContext",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsSidecarsStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsVolumeDevices",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSidecarsVolumeMounts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSidecarsVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSuspend",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSuspend)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSynchronization",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSynchronization)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSynchronizationMutex",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSynchronizationMutex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSynchronizationSemaphore",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSynchronizationSemaphore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsSynchronizationSemaphoreConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsSynchronizationSemaphoreConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsTolerations",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumes",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesAwsElasticBlockStore",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesAwsElasticBlockStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesAzureDisk",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesAzureDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesAzureFile",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesAzureFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesCephfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesCephfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesCephfsSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesCephfsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesCinder",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesCinder)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesCinderSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesCinderSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesConfigMapItems",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesCsi",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesCsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesCsiNodePublishSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesCsiNodePublishSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesDownwardApi",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesDownwardApiItems",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesDownwardApiItemsFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesEmptyDir",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesEmptyDir)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesEmptyDirSizeLimit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesEmptyDirSizeLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsVolumesEmptyDirSizeLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesEphemeral",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesEphemeral)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplate",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecDataSource",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResources",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesFc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesFc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesFlexVolume",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesFlexVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesFlexVolumeSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesFlexVolumeSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesFlocker",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesFlocker)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesGcePersistentDisk",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesGcePersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesGitRepo",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesGitRepo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesGlusterfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesGlusterfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesHostPath",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesHostPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesIscsi",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesIscsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesIscsiSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesIscsiSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesNfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesNfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesPersistentVolumeClaim",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesPersistentVolumeClaim)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesPhotonPersistentDisk",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesPhotonPersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesPortworxVolume",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesPortworxVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesProjected",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesProjected)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSources",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesConfigMapItems",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApi",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItems",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesSecretItems",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesServiceAccountToken",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesQuobyte",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesQuobyte)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesRbd",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesRbd)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesRbdSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesRbdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesScaleIo",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesScaleIo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesScaleIoSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesScaleIoSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesSecretItems",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesStorageos",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesStorageos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesStorageosSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesStorageosSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplateDefaultsVolumesVsphereVolume",
		reflect.TypeOf((*WorkflowTemplateSpecTemplateDefaultsVolumesVsphereVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplates",
		reflect.TypeOf((*WorkflowTemplateSpecTemplates)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesActiveDeadlineSeconds",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesActiveDeadlineSeconds)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesActiveDeadlineSeconds{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityNodeAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAntiAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocation",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationAzure",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationGcs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationGit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationOss",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationRaw",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationS3",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesArchiveLocationS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesArchiveLocationS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainer",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerEnv",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerEnvFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerEnvFromSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerEnvValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecycle",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePostStart",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePreStop",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLivenessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLivenessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerPorts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerReadinessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerReadinessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerResources",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerResourcesLimits",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerResourcesRequests",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSecurityContext",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainers",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersEnv",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersEnvFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersEnvFromSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecycle",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStart",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStop",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersPorts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersResources",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersResourcesLimits",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersResourcesRequests",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersSecurityContext",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersVolumeDevices",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetContainersVolumeMounts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetRetryStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerSetRetryStrategyRetries{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerSetVolumeMounts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerSetVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerStartupProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerStartupProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerStartupProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesContainerStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerVolumeDevices",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesContainerVolumeMounts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesContainerVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDag",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDag)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasks",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArguments",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifacts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArchive",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsAzure",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsGcs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsGit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsOss",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsRaw",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsS3",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsParameters",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsParametersValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksContinueOn",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksContinueOn)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooks",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArguments",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArguments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifacts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArchive",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsAzure",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsGcs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsGit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsOss",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsRaw",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsParameters",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsParametersValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksHooksTemplateRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksHooksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksTemplateRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksTemplateRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksWithSequence",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksWithSequence)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesDagTasksWithSequenceCount{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesDagTasksWithSequenceEnd{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesDagTasksWithSequenceStart{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesData",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSource",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPaths",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPaths)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArchive",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArchiveTar",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsAzure",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsGcs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsGit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsOss",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsRaw",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsS3",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataSourceArtifactPathsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataSourceArtifactPathsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesDataTransformation",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesDataTransformation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesExecutor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesExecutor)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesHostAliases",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesHostAliases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesHttpBodyFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesHttpBodyFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesHttpHeadersValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesHttpHeadersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesHttpHeadersValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesHttpHeadersValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainers",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersEnv",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersEnvFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersEnvFromSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersEnvValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesInitContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecycle",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStart",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesInitContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStop",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLivenessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLivenessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesInitContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesInitContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersPorts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersReadinessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersReadinessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesInitContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesInitContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersResources",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersResourcesLimits",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesInitContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersResourcesRequests",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesInitContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersSecurityContext",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersStartupProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersStartupProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersStartupProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesInitContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesInitContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersVolumeDevices",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInitContainersVolumeMounts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInitContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifacts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsArchive",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecTemplatesInputsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecTemplatesInputsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecTemplatesInputsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsAzure",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsGcs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsGit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsOss",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsRaw",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsS3",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsParameters",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsParametersValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesInputsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesInputsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesMemoize",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesMemoize)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesMemoizeCache",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesMemoizeCache)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesMemoizeCacheConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesMemoizeCacheConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesMetrics",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesMetrics)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesMetricsPrometheus",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesMetricsPrometheus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesMetricsPrometheusCounter",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesMetricsPrometheusCounter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesMetricsPrometheusGauge",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesMetricsPrometheusGauge)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesMetricsPrometheusHistogram",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesMetricsPrometheusHistogram)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesMetricsPrometheusLabels",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesMetricsPrometheusLabels)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifacts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsArchive",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsArchiveTar",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsAzure",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsGcs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsGit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsOss",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsRaw",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsS3",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsArtifactsS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsArtifactsS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsParameters",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsParametersValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsParametersValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesOutputsParametersValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesOutputsParametersValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResource",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifact",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifact)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArchive",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArchive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArchiveTar",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArchiveTar)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactGc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactGcPodMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactGcPodMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy)(nil)).Elem(),
		map[string]interface{}{
			"ON_WORKFLOW_COMPLETION": WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy_ON_WORKFLOW_COMPLETION,
			"ON_WORKFLOW_DELETION": WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy_ON_WORKFLOW_DELETION,
			"NEVER": WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy_NEVER,
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactory",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactoryPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactoryPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactoryUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactArtifactoryUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactAzure",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactAzureAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactAzureAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactGcs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactGcs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactGcsServiceAccountKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactGcsServiceAccountKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactGit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactGit)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactGitPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactGitPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactGitSshPrivateKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactGitSshPrivateKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactGitUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactGitUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHdfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHdfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHdfsKrbCCacheSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHdfsKrbCCacheSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHdfsKrbConfigConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHdfsKrbConfigConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHdfsKrbKeytabSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHdfsKrbKeytabSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttp",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuth",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuthPasswordSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuthPasswordSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuthUsernameSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthBasicAuthUsernameSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthClientCert",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthClientCertClientCertSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthClientCertClientCertSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthClientCertClientKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthClientCertClientKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2ClientIdSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2ClientIdSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2ClientSecretSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2ClientSecretSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2EndpointParams",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2EndpointParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2TokenUrlSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpAuthOauth2TokenUrlSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactOss",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactOss)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactOssAccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactOssAccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactOssLifecycleRule",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactOssLifecycleRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactOssSecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactOssSecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactRaw",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactRaw)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactS3",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactS3AccessKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactS3AccessKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactS3CaSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactS3CaSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactS3CreateBucketIfNotPresent",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactS3CreateBucketIfNotPresent)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactS3EncryptionOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactS3EncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactS3EncryptionOptionsServerSideCustomerKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactS3EncryptionOptionsServerSideCustomerKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesResourceManifestFromArtifactS3SecretKeySecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesResourceManifestFromArtifactS3SecretKeySecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesRetryStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesRetryStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesRetryStrategyAffinity",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesRetryStrategyAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesRetryStrategyBackoff",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesRetryStrategyBackoff)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesRetryStrategyBackoffFactor{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesRetryStrategyLimit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesRetryStrategyLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesRetryStrategyLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScript",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScript)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptEnv",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptEnvFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptEnvFromSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptEnvValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesScriptEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecycle",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePostStart",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesScriptLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesScriptLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePreStop",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLivenessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLivenessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesScriptLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesScriptLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptPorts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptReadinessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptReadinessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesScriptReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesScriptReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptResources",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptResourcesLimits",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesScriptResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptResourcesRequests",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesScriptResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptSecurityContext",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptStartupProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptStartupProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptStartupProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesScriptStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesScriptStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptVolumeDevices",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesScriptVolumeMounts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesScriptVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSecurityContext",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSecurityContextSysctls",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSecurityContextSysctls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecars",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecars)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsEnv",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsEnvFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsEnvFromConfigMapRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsEnvFromSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsEnvValueFrom",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsEnvValueFromFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsEnvValueFromSecretKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecycle",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStart",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStop",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLivenessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLivenessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLivenessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLivenessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLivenessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLivenessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsLivenessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsPorts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsReadinessProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsReadinessProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsReadinessProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsReadinessProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsReadinessProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsReadinessProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsReadinessProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsResources",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsResourcesLimits",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsResourcesRequests",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsSecurityContext",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsSecurityContextCapabilities",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsSecurityContextSeLinuxOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsSecurityContextSeccompProfile",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsSecurityContextWindowsOptions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsStartupProbe",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsStartupProbeExec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsStartupProbeGrpc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGet",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsStartupProbeTcpSocket",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsStartupProbeTcpSocketPort",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesSidecarsStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsVolumeDevices",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSidecarsVolumeMounts",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSidecarsVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSuspend",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSuspend)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSynchronization",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSynchronization)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSynchronizationMutex",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSynchronizationMutex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSynchronizationSemaphore",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSynchronizationSemaphore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesSynchronizationSemaphoreConfigMapKeyRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesSynchronizationSemaphoreConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesTolerations",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumes",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesAwsElasticBlockStore",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesAwsElasticBlockStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesAzureDisk",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesAzureDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesAzureFile",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesAzureFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesCephfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesCephfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesCephfsSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesCephfsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesCinder",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesCinder)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesCinderSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesCinderSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesConfigMapItems",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesCsi",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesCsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesCsiNodePublishSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesCsiNodePublishSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesDownwardApi",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesDownwardApiItems",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesDownwardApiItemsFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesVolumesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesEmptyDir",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesEmptyDir)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesVolumesEmptyDirSizeLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesEphemeral",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesEphemeral)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplate",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpec",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSource",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResources",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelector",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesFc",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesFc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesFlexVolume",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesFlexVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesFlexVolumeSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesFlexVolumeSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesFlocker",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesFlocker)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesGcePersistentDisk",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesGcePersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesGitRepo",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesGitRepo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesGlusterfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesGlusterfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesHostPath",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesHostPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesIscsi",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesIscsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesIscsiSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesIscsiSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesNfs",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesNfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesPersistentVolumeClaim",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesPersistentVolumeClaim)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesPhotonPersistentDisk",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesPhotonPersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesPortworxVolume",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesPortworxVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesProjected",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesProjected)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesProjectedSources",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesProjectedSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesProjectedSourcesConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesProjectedSourcesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesProjectedSourcesConfigMapItems",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesProjectedSourcesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesProjectedSourcesDownwardApi",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesProjectedSourcesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesProjectedSourcesDownwardApiItems",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesProjectedSourcesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesProjectedSourcesDownwardApiItemsFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesProjectedSourcesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesProjectedSourcesSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesProjectedSourcesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesProjectedSourcesSecretItems",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesProjectedSourcesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesProjectedSourcesServiceAccountToken",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesProjectedSourcesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesQuobyte",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesQuobyte)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesRbd",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesRbd)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesRbdSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesRbdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesScaleIo",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesScaleIo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesScaleIoSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesScaleIoSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesSecret",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesSecretItems",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesStorageos",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesStorageos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesStorageosSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesStorageosSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTemplatesVolumesVsphereVolume",
		reflect.TypeOf((*WorkflowTemplateSpecTemplatesVolumesVsphereVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTolerations",
		reflect.TypeOf((*WorkflowTemplateSpecTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecTtlStrategy",
		reflect.TypeOf((*WorkflowTemplateSpecTtlStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimGc",
		reflect.TypeOf((*WorkflowTemplateSpecVolumeClaimGc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplates",
		reflect.TypeOf((*WorkflowTemplateSpecVolumeClaimTemplates)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesSpec",
		reflect.TypeOf((*WorkflowTemplateSpecVolumeClaimTemplatesSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesSpecDataSource",
		reflect.TypeOf((*WorkflowTemplateSpecVolumeClaimTemplatesSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesSpecDataSourceRef",
		reflect.TypeOf((*WorkflowTemplateSpecVolumeClaimTemplatesSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesSpecResources",
		reflect.TypeOf((*WorkflowTemplateSpecVolumeClaimTemplatesSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits",
		reflect.TypeOf((*WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests",
		reflect.TypeOf((*WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecVolumeClaimTemplatesSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesSpecSelector",
		reflect.TypeOf((*WorkflowTemplateSpecVolumeClaimTemplatesSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesSpecSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecVolumeClaimTemplatesSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesStatus",
		reflect.TypeOf((*WorkflowTemplateSpecVolumeClaimTemplatesStatus)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources",
		reflect.TypeOf((*WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecVolumeClaimTemplatesStatusAllocatedResources{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity",
		reflect.TypeOf((*WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecVolumeClaimTemplatesStatusCapacity{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumeClaimTemplatesStatusConditions",
		reflect.TypeOf((*WorkflowTemplateSpecVolumeClaimTemplatesStatusConditions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumes",
		reflect.TypeOf((*WorkflowTemplateSpecVolumes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesAwsElasticBlockStore",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesAwsElasticBlockStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesAzureDisk",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesAzureDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesAzureFile",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesAzureFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesCephfs",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesCephfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesCephfsSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesCephfsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesCinder",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesCinder)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesCinderSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesCinderSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesConfigMapItems",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesCsi",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesCsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesCsiNodePublishSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesCsiNodePublishSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesDownwardApi",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesDownwardApiItems",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesDownwardApiItemsFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesEmptyDir",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesEmptyDir)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesEmptyDirSizeLimit",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesEmptyDirSizeLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecVolumesEmptyDirSizeLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesEphemeral",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesEphemeral)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplate",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpec",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResources",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecSelector",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesFc",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesFc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesFlexVolume",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesFlexVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesFlexVolumeSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesFlexVolumeSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesFlocker",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesFlocker)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesGcePersistentDisk",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesGcePersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesGitRepo",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesGitRepo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesGlusterfs",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesGlusterfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesHostPath",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesHostPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesIscsi",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesIscsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesIscsiSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesIscsiSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesNfs",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesNfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesPersistentVolumeClaim",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesPersistentVolumeClaim)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesPhotonPersistentDisk",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesPhotonPersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesPortworxVolume",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesPortworxVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesProjected",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesProjected)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesProjectedSources",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesProjectedSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesProjectedSourcesConfigMap",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesProjectedSourcesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesProjectedSourcesConfigMapItems",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesProjectedSourcesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesProjectedSourcesDownwardApi",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesProjectedSourcesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItems",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesProjectedSourcesSecret",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesProjectedSourcesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesProjectedSourcesSecretItems",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesProjectedSourcesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesProjectedSourcesServiceAccountToken",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesProjectedSourcesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesQuobyte",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesQuobyte)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesRbd",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesRbd)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesRbdSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesRbdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesScaleIo",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesScaleIo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesScaleIoSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesScaleIoSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesSecret",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesSecretItems",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesStorageos",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesStorageos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesStorageosSecretRef",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesStorageosSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecVolumesVsphereVolume",
		reflect.TypeOf((*WorkflowTemplateSpecVolumesVsphereVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecWorkflowMetadata",
		reflect.TypeOf((*WorkflowTemplateSpecWorkflowMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecWorkflowMetadataLabelsFrom",
		reflect.TypeOf((*WorkflowTemplateSpecWorkflowMetadataLabelsFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"workflowtemplates _argoprojio.WorkflowTemplateSpecWorkflowTemplateRef",
		reflect.TypeOf((*WorkflowTemplateSpecWorkflowTemplateRef)(nil)).Elem(),
	)
}
