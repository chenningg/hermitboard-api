package auth

import (
	"fmt"

	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/pulid"
	"github.com/google/uuid"
)

// SessionRedisKey represents the key prefix (e.g. <prefix>:<actual key>) of a Session.
// It is set as a var in case it needs to be changed.
var SessionRedisKey = "session"

// SessionID represents a session ID as a string.
type SessionID string

// Session represents an authentication session, containing the user ID and their authorization roles.
type Session struct {
	UserID pulid.PULID      `redis:"userID"`
	Roles  []authrole.Value `redis:"roles"`
}

// NewSessionID returns a session ID implemented as a UUIDv4.
func NewSessionID() (SessionID, error) {
	randomID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("could not generate a new UUIDv4 for session ID: %v", err)
	}

	return SessionID(randomID.String()), nil
}

// ParseSessionID checks if the string provided is a valid session ID.
func ParseSessionID(s string) (SessionID, error) {
	sessionID, err := uuid.Parse(s)
	if err != nil {
		return "", fmt.Errorf("invalid session ID format: %v", err)
	}
	return SessionID(sessionID.String()), nil
}
