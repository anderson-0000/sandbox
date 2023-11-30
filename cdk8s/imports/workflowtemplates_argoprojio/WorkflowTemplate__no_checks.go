//go:build no_runtime_type_checking

package workflowtemplates_argoprojio

// Building without runtime type checking enabled, so all the below just return nil

func validateWorkflowTemplate_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateWorkflowTemplate_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWorkflowTemplate_ManifestParameters(props *WorkflowTemplateProps) error {
	return nil
}

func validateWorkflowTemplate_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewWorkflowTemplateParameters(scope constructs.Construct, id *string, props *WorkflowTemplateProps) error {
	return nil
}

