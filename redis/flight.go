package redis

import (
	_ "embed"
	"fmt"

	"github.com/davidmdm/yoke/pkg/helm"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

//go:embed redis-19.0.1.tgz
var archive []byte

// RenderChart renders the chart downloaded from oci://registry-1.docker.io/bitnamicharts/redis
// Producing version: 19.0.1
func RenderChart(release, namespace string, values *Values) ([]*unstructured.Unstructured, error) {
	chart, err := helm.LoadChartFromZippedArchive(archive)
	if err != nil {
		return nil, fmt.Errorf("failed to load chart from zipped archive: %w", err)
	}

	return chart.Render(release, namespace, values)
}
