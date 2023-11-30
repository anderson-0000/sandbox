//go:build no_runtime_type_checking

package cronworkflow_argoprojio

// Building without runtime type checking enabled, so all the below just return nil

func validateCronWorkflow_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateCronWorkflow_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCronWorkflow_ManifestParameters(props *CronWorkflowProps) error {
	return nil
}

func validateCronWorkflow_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewCronWorkflowParameters(scope constructs.Construct, id *string, props *CronWorkflowProps) error {
	return nil
}

