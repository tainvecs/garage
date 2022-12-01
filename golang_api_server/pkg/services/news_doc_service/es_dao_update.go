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

	// check if createdAt, updatedAt is set
	if doc.CreatedAt != nil {
		return errors.New("bad es update request: created_at should not be set")
	}
	if doc.UpdatedAt == nil {
		t := time.Now()
		doc.UpdatedAt = &t
	}

	return dao.Client.Update(ctx, dao.IndexIndex, doc.UUID, doc)
}
