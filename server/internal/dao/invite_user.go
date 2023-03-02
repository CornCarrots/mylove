// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"server/internal/dao/internal"
)

// internalInviteUserDao is internal type for wrapping internal DAO implements.
type internalInviteUserDao = *internal.InviteUserDao

// inviteUserDao is the data access object for table invite_user.
// You can define custom methods on it to extend its functionality as you wish.
type inviteUserDao struct {
	internalInviteUserDao
}

var (
	// InviteUser is globally public accessible object for table invite_user operations.
	InviteUser = inviteUserDao{
		internal.NewInviteUserDao(),
	}
)

// Fill with you ideas below.
