package types

import "time"

// PipelineConfig model.
type PipelineConfig struct {
	ID        string    `json:"id" yaml:"id"`
	RawConfig string    `json:"rawConfig" yaml:"rawConfig"`
	CreatedAt time.Time `json:"createdAt" yaml:"createdAt"`
}

// PipelineConfigHistoryParams request payload for querying the pipeline config history.
type PipelineConfigHistoryParams struct {
	Last *uint64
}
