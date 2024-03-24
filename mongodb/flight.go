package mongodb

import (
	_ "embed"
	"fmt"

	"github.com/davidmdm/yoke/pkg/helm"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

//go:embed mongodb-15.1.0.tgz
var archive []byte

// RenderChart renders the chart downloaded from oci://registry-1.docker.io/bitnamicharts/mongodb
// Producing version: 15.1.0
func RenderChart(release, namespace string, values *Values) ([]*unstructured.Unstructured, error) {
	chart, err := helm.LoadChartFromZippedArchive(archive)
	if err != nil {
		return nil, fmt.Errorf("failed to load chart from zipped archive: %w", err)
	}

	return chart.Render(release, namespace, values)
}
