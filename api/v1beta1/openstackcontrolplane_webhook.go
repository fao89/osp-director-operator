/*


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

// Generated by:
//
// operator-sdk create webhook --group osp-director --version v1beta1 --kind OpenStackControlPlane --programmatic-validation
//

package v1beta1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var controlplanelog = logf.Log.WithName("controlplane-resource")

// SetupWebhookWithManager - register this webhook with the controller manager
func (r *OpenStackControlPlane) SetupWebhookWithManager(mgr ctrl.Manager) error {
	if webhookClient == nil {
		webhookClient = mgr.GetClient()
	}

	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-osp-director-openstack-org-v1beta1-openstackcontrolplane,mutating=false,failurePolicy=fail,groups=osp-director.openstack.org,resources=openstackcontrolplanes,versions=v1beta1,name=vopenstackcontrolplane.kb.io

var _ webhook.Validator = &OpenStackControlPlane{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *OpenStackControlPlane) ValidateCreate() error {
	controlplanelog.Info("validate create", "name", r.Name)

	return r.checkBaseImageReqs()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *OpenStackControlPlane) ValidateUpdate(old runtime.Object) error {
	controlplanelog.Info("validate update", "name", r.Name)

	return r.checkBaseImageReqs()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *OpenStackControlPlane) ValidateDelete() error {
	controlplanelog.Info("validate delete", "name", r.Name)

	return nil
}

func (r *OpenStackControlPlane) checkBaseImageReqs() error {
	for _, role := range r.Spec.VirtualMachineRoles {
		if role.BaseImageURL == "" && role.BaseImageVolumeName == "" {
			return fmt.Errorf(fmt.Sprintf("Either \"baseImageURL\" or \"baseImageVolumeName\" must be provided for role %s", role.RoleName))
		}

	}
	return nil
}
