package controllers

import (
	"context"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/fearlesschenc/operator-utils/pkg/reconcile"
)

type Reconciler struct {
}

func (r *Reconciler) UpdateStatus(ctx context.Context, cm *corev1.ConfigMap) error {
	return nil
}

func (r *Reconciler) GetReconcileSteps() []reconcile.Func {
	return []reconcile.Func{
		r.reconcileFoo,
		r.reconcileBar,
	}
}

func (r *Reconciler) reconcileFoo(ctx context.Context, obj client.Object) (reconcile.Result, error) {
	// logic
	return reconcile.Reconcile(ctx, obj).
		WithReconciler(reconcile.Funcs{
			r.reconcileFoo1,
			r.reconcileFoo2,
		})
}

func (r *Reconciler) reconcileFoo1(ctx context.Context, obj client.Object) (reconcile.Result, error) {
	return reconcile.Continue()
}

func (r *Reconciler) reconcileFoo2(ctx context.Context, obj client.Object) (reconcile.Result, error) {
	return reconcile.Continue()
}

func (r *Reconciler) reconcileBar(ctx context.Context, obj client.Object) (reconcile.Result, error) {
	return reconcile.Continue()
}

func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	obj := &corev1.ConfigMap{}

	result, err := reconcile.Reconcile(ctx, obj).WithReconciler(r)
	if result.RequeueRequest {
		return reconcile.RequeueRequestAfter(result.RequeueDelay, err)
	}

	if result.CancelReconciliation {
		return reconcile.DoNotRequeueRequest(err)
	}

	return reconcile.DoNotRequeueRequest(nil)
}
