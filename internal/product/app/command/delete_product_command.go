package command

import sharedvo "deligo/internal/shared/valueobject"

type DeleteProductCommand struct {
	ID sharedvo.ID
}
