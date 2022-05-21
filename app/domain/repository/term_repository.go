package repository

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/google/uuid"
)

type ITermRepository interface {
	FindTermById(ctx context.Context, uuid uuid.UUID) (*entities.Term, error)
}
