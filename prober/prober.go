package prober

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/Tiendeo/blackbox_exporter/config"
)

type ProbeFn func(ctx context.Context, target string, config config.Module, registry *prometheus.Registry, logger log.Logger, extraParams map[string]string) bool
