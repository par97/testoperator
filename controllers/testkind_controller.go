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

package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"

	testgroupv1alpha1 "github.com/par97/testoperator/api/v1alpha1"
)

// TestkindReconciler reconciles a Testkind object
type TestkindReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=testgroup.par.xyz,resources=testkinds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=testgroup.par.xyz,resources=testkinds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=testgroup.par.xyz,resources=testkinds/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Testkind object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *TestkindReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	log := ctrllog.Log.WithValues("testkind", req.NamespacedName)
	log.Info("Begging of reconcile()")

	testkind := &testgroupv1alpha1.Testkind{}
	err := r.Get(ctx, req.NamespacedName, testkind)
	if err != nil {
		// Error reading the object - requeue the req.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("name: " + testkind.Name)
	log.Info("country: " + testkind.Spec.Country)
	log.Info("cities: " + fmt.Sprint(testkind.Spec.Cities))
	log.Info(fmt.Sprintf("city number: %d", len(testkind.Spec.Cities)))
	log.Info("params: " + fmt.Sprint(testkind.Spec.Params))
	log.Info(fmt.Sprintf("param number: %d", len(testkind.Spec.Params)))

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TestkindReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testgroupv1alpha1.Testkind{}).
		Complete(r)
}
