package psqldao

import "gorm.io/gorm"

// QueryConfig defines the dao common query configs
type QueryConfig struct {
	Fields              []string
	Offset              int
	Limit               int
	PreloadAssociations []string
}

// Apply QueryConfig to database client
func (conf QueryConfig) Apply(client *gorm.DB) *gorm.DB {

	if len(conf.Fields) > 0 {
		client = client.Select(conf.Fields)
	}

	if conf.Offset > 0 {
		client = client.Offset(conf.Offset)
	}

	if conf.Limit > 0 {
		client = client.Limit(conf.Limit)
	}

	for _, assoc := range conf.PreloadAssociations {
		client = client.Preload(assoc)
	}

	return client
}
