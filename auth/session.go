package auth

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/pulid"
	"github.com/google/uuid"
)

// SessionRedisKey represents the key prefix (e.g. <prefix>:<actual key>) of a Session.
// It is set as a var in case it needs to be changed.
var SessionRedisKey = "session"

// Session ID key type for context
type sessionContextKey int

const SessionContextKey sessionContextKey = 0

// AuthorizationHeader is the name of the HTTP Header which contains the session id.
// Exported so that it can be changed by developers.
var AuthorizationHeader = "Authorization"

var NonStaffAuthRoles = []authrole.Value{
	authrole.ValueDemo, authrole.ValueFree, authrole.ValuePlus, authrole.ValuePro, authrole.ValueEnterprise,
}
var StaffAuthRoles = []authrole.Value{authrole.ValueSupport, authrole.ValueAdmin, authrole.ValueSuperAdmin}

// SessionID represents a session ID as a string.
type SessionID string

func (sessionID SessionID) String() string {
	return string(sessionID)
}

// Session represents an authentication session on the server side, containing the user ID and their authorization roles.
type Session struct {
	SessionID SessionID                   // Session ID of the logged-in user.
	UserID    pulid.PULID                 // ID of the user.
	AuthRoles map[authrole.Value]struct{} // Auth roles of the user.
}

// NewSessionID returns a session ID implemented as a UUIDv4.
func NewSessionID() (SessionID, error) {
	randomID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("auth.NewSessionID(): could not generate a new session ID: %v", err)
	}

	return SessionID(randomID.String()), nil
}

// ParseSessionID checks if the string provided is a valid session ID.
func ParseSessionID(s string) (SessionID, error) {
	if s == "" {
		return "", fmt.Errorf("auth.ParseSessionID(): string to be parsed is empty")
	}

	sessionID, err := uuid.Parse(s)
	if err != nil {
		return "", fmt.Errorf("auth.ParseSessionID(): invalid session ID format: %v", err)
	}
	return SessionID(sessionID.String()), nil
}

// GetSessionFromContext returns the Session from a provided context, or nil if it is not found.
func GetSessionFromContext(ctx context.Context) *Session {
	if ctx == nil {
		return nil
	}
	if session, ok := ctx.Value(SessionContextKey).(Session); ok {
		return &session
	}
	return nil
}

// HasAuthRoles checks if any auth roles specified exist on the given Session. It returns true as long as one matching auth role is found.
func HasAuthRoles(session *Session, authRoles []authrole.Value) bool {
	if session == nil {
		return false
	}

	for _, authRole := range authRoles {
		if _, ok := session.AuthRoles[authRole]; ok {
			return true
		}
	}
	return false
}

// IsLoggedIn checks if the user is logged in (has a session).
func IsLoggedIn(session *Session) bool {
	if session == nil {
		return false
	} else {
		_, err := ParseSessionID(session.SessionID.String())
		if err != nil {
			return false
		}
	}
	return true
}

// GetAuthRolesFromSession returns a list of auth roles from a session.
func GetAuthRolesFromSession(session *Session) []authrole.Value {
	if session == nil {
		return []authrole.Value{}
	}

	keys := make([]authrole.Value, len(session.AuthRoles))

	i := 0
	for key := range session.AuthRoles {
		keys[i] = key
		i++
	}

	return keys
}

// MarshalGQL implements graphql.Marshaler interface.
func (sessionID SessionID) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(sessionID.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (sessionID *SessionID) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("%w: session ID %T must be a string", ErrBadInput, val)
	}
	parsedSessionID, err := ParseSessionID(str)
	if err != nil {
		return fmt.Errorf("%w: %s is not a valid session ID", ErrBadInput, str)
	}

	*sessionID = parsedSessionID
	return nil
}
