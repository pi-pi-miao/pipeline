/*
Copyright 2019 The Tekton Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package builder

import (
	v1beta1 "github.com/tektoncd/pipeline/internal/builder/v1beta1"
)

// ContainerOp is an operation which modifies a Container struct.
// Deprecated: moved to internal/builder/v1alpha1
type ContainerOp = v1beta1.ContainerOp

// VolumeMountOp is an operation which modifies a VolumeMount struct.
// Deprecated: moved to internal/builder/v1alpha1
type VolumeMountOp = v1beta1.VolumeMountOp

// ResourceRequirementsOp is an operation which modifies a ResourceRequirements struct.
// Deprecated: moved to internal/builder/v1alpha1
type ResourceRequirementsOp = v1beta1.ResourceRequirementsOp

// ResourceListOp is an operation which modifies a ResourceList struct.
// Deprecated: moved to internal/builder/v1alpha1
type ResourceListOp = v1beta1.ResourceListOp

var (
	// Command sets the command to the Container (step in this case).
	// Deprecated: moved to internal/builder/v1alpha1
	Command = v1beta1.Command

	// Args sets the command arguments to the Container (step in this case).
	// Deprecated: moved to internal/builder/v1alpha1
	Args = v1beta1.Args

	// EnvVar add an environment variable, with specified name and value, to the Container (step).
	// Deprecated: moved to internal/builder/v1alpha1
	EnvVar = v1beta1.EnvVar

	// WorkingDir sets the WorkingDir on the Container.
	// Deprecated: moved to internal/builder/v1alpha1
	WorkingDir = v1beta1.WorkingDir

	// VolumeMount add a VolumeMount to the Container (step).
	// Deprecated: moved to internal/builder/v1alpha1
	VolumeMount = v1beta1.VolumeMount

	// Resources adds ResourceRequirements to the Container (step).
	// Deprecated: moved to internal/builder/v1alpha1
	Resources = v1beta1.Resources

	// Limits adds Limits to the ResourceRequirements.
	// Deprecated: moved to internal/builder/v1alpha1
	Limits = v1beta1.Limits

	// Requests adds Requests to the ResourceRequirements.
	// Deprecated: moved to internal/builder/v1alpha1
	Requests = v1beta1.Requests

	// CPU sets the CPU resource on the ResourceList.
	// Deprecated: moved to internal/builder/v1alpha1
	CPU = v1beta1.CPU

	// Memory sets the memory resource on the ResourceList.
	// Deprecated: moved to internal/builder/v1alpha1
	Memory = v1beta1.Memory

	// EphemeralStorage sets the ephemeral storage resource on the ResourceList.
	// Deprecated: moved to internal/builder/v1alpha1
	EphemeralStorage = v1beta1.EphemeralStorage

	// TerminationMessagePath sets the termination message path.
	// Deprecated: moved to internal/builder/v1alpha1
	TerminationMessagePath = v1beta1.TerminationMessagePath
)
