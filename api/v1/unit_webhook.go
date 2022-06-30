/*
Copyright 2022.

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

package v1

import (
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var unitlog = logf.Log.WithName("unit-resource")

var CategoryDeployment = "Deployment"
var CategoryStatefulSet = "StatefulSet"

func (r *Unit) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-custom-hmlss-ml-v1-unit,mutating=true,failurePolicy=fail,sideEffects=None,groups=custom.hmlss.ml,resources=units,verbs=create;update,versions=v1,name=munit.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Unit{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Unit) Default() {
	unitlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	unitlog.Info("default", "name", r.Name)

	//default replicas set 1
	if r.Spec.Replicas == nil {
		defaultReplicas := int32(1)
		r.Spec.Replicas = &defaultReplicas
	}

	//default selector
	labelMap := make(map[string]string, 1)
	labelMap["app"] = r.Name
	r.Spec.Selector = &metav1.LabelSelector{
		MatchLabels: labelMap,
	}

	r.Spec.Template.Labels = labelMap

	r.Status.LastUpdateTime = metav1.Now()
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-custom-hmlss-ml-v1-unit,mutating=false,failurePolicy=fail,sideEffects=None,groups=custom.hmlss.ml,resources=units,verbs=create;update,versions=v1,name=vunit.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Unit{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Unit) ValidateCreate() error {
	unitlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	//检查unit.Spec.Category
	//对于这种校验操作可以放到reconcile里面，但是不适合，还是在validationWebhook合适
	switch r.Spec.Category {
	case CategoryDeployment:
		return nil
	case CategoryStatefulSet:
		return nil
	default:
		err := errors.New("spec.category only support Deployment or StatefulSet")
		unitlog.Error(err, "creating validate failed", "name", r.Name)
		return err
	}
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Unit) ValidateUpdate(old runtime.Object) error {
	unitlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Unit) ValidateDelete() error {
	unitlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
