package kudoinit

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/kudobuilder/kudo/pkg/kudoctl/kube"
)

const (
	DefaultNamespace      = "kudo-system"
	defaultGracePeriod    = 10
	defaultServiceAccount = "kudo-manager"
)

type InitStep interface {
	fmt.Stringer

	// Should return an error if the installation will not be possible
	PreInstallCheck(client *kube.Client) Result

	// Executes the actual installation
	Install(client *kube.Client) error

	// Returns the installed artifacts as yaml manifests
	AsYamlManifests() ([]string, error)
	AsArray() []runtime.Object
}

func GenerateLabels(labels map[string]string) map[string]string {
	labels["app"] = "kudo-manager"
	return labels
}
