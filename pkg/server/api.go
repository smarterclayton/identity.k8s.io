package server

import (
	"net/http"

	"github.com/golang/glog"

	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/identity/pkg/apis/identity"
	"k8s.io/identity/pkg/apis/identity/install"
	"k8s.io/identity/pkg/apis/identity/v1alpha1"
	"k8s.io/identity/pkg/jwt"
	identitydocumentstorage "k8s.io/identity/pkg/registry/identity/identitydocument"
)

var (
	groupFactoryRegistry = make(announced.APIGroupFactoryRegistry)
	registry             = registered.NewOrDie("")
	Scheme               = runtime.NewScheme()
	Codecs               = serializer.NewCodecFactory(Scheme)
)

func init() {
	install.Install(groupFactoryRegistry, registry, Scheme)

	// we need to add the options to empty v1
	// TODO fix the server code to avoid this
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})

	// TODO: keep the generic API server from wanting this
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)
}

type ExtraConfig struct {
	// Place you custom config here.
}

type Config struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}

// Server contains state for a Kubernetes cluster master/api server.
type Server struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}

type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (cfg *Config) Complete() CompletedConfig {
	c := completedConfig{
		cfg.GenericConfig.Complete(),
		&cfg.ExtraConfig,
	}

	c.GenericConfig.Version = &version.Info{
		Major: "1",
		Minor: "0",
	}

	return CompletedConfig{&c}
}

// New returns a new instance of erver from the given config.
func (c completedConfig) New() (*Server, error) {
	genericServer, err := c.GenericConfig.New("identity-api", genericapiserver.EmptyDelegate)
	if err != nil {
		return nil, err
	}

	s := &Server{
		GenericAPIServer: genericServer,
	}

	mux := s.GenericAPIServer.Handler.NonGoRestfulMux
	mux.Handle("/webhook/authorize", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		glog.V(2).Infof("authorize")
		w.WriteHeader(401)
	}))
	mux.Handle("/certs", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(404)
	}))

	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(identity.GroupName, registry, Scheme, metav1.ParameterCodec, Codecs)
	apiGroupInfo.GroupMeta.GroupVersion = v1alpha1.SchemeGroupVersion
	v1alpha1storage := map[string]rest.Storage{}
	v1alpha1storage["identitydocuments"] = identitydocumentstorage.NewREST(jwt.NewSigner())
	apiGroupInfo.VersionedResourcesStorageMap["v1alpha1"] = v1alpha1storage

	if err := s.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}

	return s, nil
}
