package reconcile

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type job struct {
	ctx    context.Context
	object metav1.Object
}

func Reconcile(ctx context.Context, obj metav1.Object) *job {
	return &job{ctx: ctx, object: obj}
}

func (j *job) WithReconciler(reconciler Reconciler) (result Result, err error) {
	steps := []Func{}
	if initializer, ok := reconciler.(Initializer); ok {
		steps = append(steps, initializer.Initialize)
	}
	if validator, ok := reconciler.(Validator); ok {
		steps = append(steps, validator.Validate)
	}
	if finalizer, ok := reconciler.(Finalizer); ok {
		steps = append(steps, finalizer.Finalize)
	}
	if genericSteps := reconciler.GetReconcileSteps(); genericSteps != nil && len(genericSteps) > 0 {
		steps = append(steps, genericSteps...)
	}

	func() {
		if statusUpdater, ok := reconciler.(StatusUpdater); ok {
			defer func() {
				if updateErr := statusUpdater.UpdateStatus(j.ctx, j.object); err != nil {
					err = updateErr
				}
			}()
		}

		for _, f := range steps {
			result, err = f(j.ctx, j.object)
			if err != nil || result.RequeueRequest || result.CancelReconciliation {
				return
			}
		}
	}()

	return
}
