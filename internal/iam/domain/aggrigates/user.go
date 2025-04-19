package aggrigates

import "deligo/internal/user/domain/entities"

type UserAggrigate struct {
	User    entities.UserEntity
	Profile entities.ProfileEntity
}
