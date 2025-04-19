package aggrigates

import "deligo/internal/iam/domain/entities"

type UserAggrigate struct {
	User    entities.UserEntity
	Profile entities.ProfileEntity
}
