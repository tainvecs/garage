package news_doc_service

import (
	"context"
	"errors"
	"strings"
	"time"
)

func (dao *newsDocESDAO) Index(ctx context.Context, doc *NewsDoc) error {

	// check if there is missing field: uuid
	if strings.TrimSpace(doc.UUID) == "" {
		return errors.New("bad es update request: missing doc UUID")
	}

	// set created_at to now
	if doc.CreatedAt == nil {
		t := time.Now()
		doc.CreatedAt = &t
	}

	return dao.Client.Index(ctx, dao.IndexIndex, doc.UUID, doc)
}
