package ent
    
import (
    "entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeGroup   = "Group"
	TypeKey     = "Key"
	TypeMember  = "Member"
	TypeMessage = "Message"
)
