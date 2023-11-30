package workflowtemplates_argoprojio


type WorkflowTemplateSpecArtifactGcStrategy string

const (
	// OnWorkflowCompletion.
	WorkflowTemplateSpecArtifactGcStrategy_ON_WORKFLOW_COMPLETION WorkflowTemplateSpecArtifactGcStrategy = "ON_WORKFLOW_COMPLETION"
	// OnWorkflowDeletion.
	WorkflowTemplateSpecArtifactGcStrategy_ON_WORKFLOW_DELETION WorkflowTemplateSpecArtifactGcStrategy = "ON_WORKFLOW_DELETION"
	// Never.
	WorkflowTemplateSpecArtifactGcStrategy_NEVER WorkflowTemplateSpecArtifactGcStrategy = "NEVER"
)

