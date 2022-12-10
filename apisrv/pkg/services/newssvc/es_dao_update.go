package newssvc

import (
	"context"
	"errors"
	"strings"
	"time"
)

// Update ESNewsDoc in elasticsearch
func (dao *esDAO) Update(ctx context.Context, doc *ESNewsDoc) error {

	// check if there is missing field: id
	if strings.TrimSpace(doc.UUID) == "" {
		return errors.New("bad es update request: missing doc UUID")
	}

	// check if createdAt, updatedAt is set
	if doc.CreatedAt != nil {
		return errors.New("bad es update request: created_at should not be set")
	}
	if doc.UpdatedAt == nil {
		t := time.Now()
		doc.UpdatedAt = &t
	}

	return dao.DataAccessObject.Update(ctx, doc.UUID, doc)
}
