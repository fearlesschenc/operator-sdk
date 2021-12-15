package reconcile

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
)

type Func func(ctx context.Context, object runtime.Object) (Result, error)

type Funcs []Func

func (fns Funcs) GetReconcileSteps() []Func {
	return fns
}

var _ Reconciler = Funcs{}

type Reconciler interface {
	GetReconcileSteps() []Func
}

type StatusUpdater interface {
	UpdateStatus(ctx context.Context, object runtime.Object) error
}

type Validator interface {
	Validate(ctx context.Context, object runtime.Object) (Result, error)
}

type Initializer interface {
	Initialize(ctx context.Context, object runtime.Object) (Result, error)
}

type Finalizer interface {
	Finalize(ctx context.Context, object runtime.Object) (Result, error)
}
