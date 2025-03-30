package aggrigates

import "delivery/internal/user/domain/entities"

type UserAggrigate struct {
	User    entities.UserEntity
	Profile entities.ProfileEntity
}
