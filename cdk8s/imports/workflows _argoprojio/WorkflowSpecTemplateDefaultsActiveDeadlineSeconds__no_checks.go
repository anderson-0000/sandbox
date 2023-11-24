//go:build no_runtime_type_checking

package workflows _argoprojio

// Building without runtime type checking enabled, so all the below just return nil

func validateWorkflowSpecTemplateDefaultsActiveDeadlineSeconds_FromNumberParameters(value *float64) error {
	return nil
}

func validateWorkflowSpecTemplateDefaultsActiveDeadlineSeconds_FromStringParameters(value *string) error {
	return nil
}

