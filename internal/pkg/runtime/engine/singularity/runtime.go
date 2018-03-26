package runtime

import (
	singularityConfig "github.com/singularityware/singularity/internal/pkg/runtime/engine/singularity/config"
	config "github.com/singularityware/singularity/pkg/runtime/config"
	oci "github.com/singularityware/singularity/pkg/runtime/oci/config"
)

type RuntimeEngine struct {
	singularityConfig.RuntimeEngineConfig
}

func (e *RuntimeEngine) InitConfig() *config.RuntimeConfig {
	if e.FileConfig == nil {
		e.FileConfig = &singularityConfig.Configuration{}
		if err := config.Parser("/usr/local/etc/singularity/singularity.conf", e.FileConfig); err != nil {
			return nil
		}
	}
	cfg := &e.RuntimeConfig
	oci.DefaultRuntimeOciConfig(&cfg.OciConfig)
	return cfg
}

func (e *RuntimeEngine) IsRunAsInstance() bool {
	return false
}
