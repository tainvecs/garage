package newssvc

import (
	"context"
	"errors"
	"strings"
	"time"
)

// Index ESNewsDoc in elasticsearch
func (dao *esDAO) Index(ctx context.Context, doc *ESNewsDoc) error {

	// check if there is missing field: uuid
	if strings.TrimSpace(doc.UUID) == "" {
		return errors.New("bad es update request: missing doc UUID")
	}

	// set created_at to now
	if doc.CreatedAt == nil {
		t := time.Now()
		doc.CreatedAt = &t
	}

	// dao.Client.Index
	return dao.DataAccessObject.Index(ctx, doc.UUID, doc)
}
