module github.com/fearlesschenc/operator-utils

go 1.15

require (
	k8s.io/api v0.18.19
	k8s.io/apimachinery v0.18.19
	k8s.io/client-go v0.18.19
	sigs.k8s.io/controller-runtime v0.6.5
)

// Pinned to kubernetes-1.18.19
replace (
	k8s.io/api => k8s.io/api v0.18.19
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.19
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.19
	k8s.io/client-go => k8s.io/client-go v0.18.19
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.6.5
)
