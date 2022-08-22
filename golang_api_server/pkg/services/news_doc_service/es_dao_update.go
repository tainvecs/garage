package news_doc_service

import (
	"context"
	"errors"
	"strings"
	"time"
)

func (dao *ESDAO) Update(ctx context.Context, doc *NewsDoc) error {

	// check if there is missing field: id
	if strings.TrimSpace(doc.UUID) == "" {
		return errors.New("bad es update request: missing doc UUID")
	}

	// check if createdAt, updatedAt, deletedAt is set
	if doc.CreatedAt != nil {
		return errors.New("bad es update request: created_at should not be set")
	}
	if doc.UpdatedAt != nil {
		return errors.New("bad es update request: updated_at should not be set")
	}
	if doc.DeletedAt != nil {
		return errors.New("bad es update request: deleted_at should not be set")
	}

	// set updated_at to now
	t := time.Now()
	doc.UpdatedAt = &t

	return dao.Client.Update(ctx, dao.IndexIndex, doc.UUID, doc)
}
