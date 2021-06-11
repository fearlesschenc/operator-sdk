package reconcile

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
)

type Func func(object runtime.Object) (Result, error)

type Reconciliation struct {
	object runtime.Object
}

func Reconcile(object runtime.Object) *Reconciliation {
	return &Reconciliation{object: object}
}

func (r *Reconciliation) WithSubProcedures(reconcileFunc ...Func) (Result, error) {
	for _, f := range reconcileFunc {
		result, err := f(r.object)
		if err != nil || result.RequeueRequest || result.CancelReconciliation {
			return result, err
		}
	}
	return Continue()
}

func (r *Reconciliation) WithSteps(reconcileFunc ...Func) (reconcile.Result, error) {
	for _, f := range reconcileFunc {
		result, err := f(r.object)

		if result.RequeueRequest {
			return RequeueRequestAfter(result.RequeueDelay, err)
		}

		if result.CancelReconciliation {
			return DoNotRequeueRequest(err)
		}
	}

	return DoNotRequeueRequest(nil)
}

func DoNotRequeueRequest(err error) (reconcile.Result, error) {
	return reconcile.Result{}, err
}

func RequeueRequestOnErr(err error) (reconcile.Result, error) {
	// note: reconcile will auto requeue failed request
	return reconcile.Result{}, err
}

func RequeueRequestAfter(duration time.Duration, err error) (reconcile.Result, error) {
	return reconcile.Result{RequeueAfter: duration}, err
}
