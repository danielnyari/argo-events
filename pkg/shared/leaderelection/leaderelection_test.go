package leaderelection

import (
	"context"
	"os"
	"testing"

	aev1 "github.com/argoproj/argo-events/pkg/apis/events/v1alpha1"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
)

var (
	configs = []aev1.BusConfig{
		{NATS: &aev1.NATSConfig{}},
		{JetStream: &aev1.JetStreamConfig{}},
		{JetStream: &aev1.JetStreamConfig{AccessSecret: &v1.SecretKeySelector{}}},
	}
)

func TestLeaderElectionWithInvalidEventBus(t *testing.T) {
	elector, err := NewElector(context.TODO(), aev1.BusConfig{}, "", 0, "", "", "")

	assert.Nil(t, elector)
	assert.EqualError(t, err, "invalid event bus")
}

func TestLeaderElectionWithEventBusElector(t *testing.T) {
	eventBusAuthFileMountPath = "test"

	for _, config := range configs {
		elector, err := NewElector(context.TODO(), config, "", 0, "", "", "")
		assert.Nil(t, err)

		_, ok := elector.(*natsEventBusElector)
		assert.True(t, ok)
	}
}

func TestLeaderElectionWithKubernetesElector(t *testing.T) {
	eventBusAuthFileMountPath = "test"

	os.Setenv(aev1.EnvVarLeaderElection, "k8s")

	for _, config := range configs {
		elector, err := NewElector(context.TODO(), config, "", 0, "", "", "")
		assert.Nil(t, err)

		_, ok := elector.(*kubernetesElector)
		assert.True(t, ok)
	}
}
