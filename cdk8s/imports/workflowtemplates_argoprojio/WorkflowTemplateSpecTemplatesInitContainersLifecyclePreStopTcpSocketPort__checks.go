//go:build !no_runtime_type_checking

package workflowtemplates_argoprojio

import (
	"fmt"
)

func validateWorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateWorkflowTemplateSpecTemplatesInitContainersLifecyclePreStopTcpSocketPort_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

