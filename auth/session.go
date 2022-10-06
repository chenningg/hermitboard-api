package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/chenningg/hermitboard-api/ent/authrole"
	"github.com/chenningg/hermitboard-api/pulid"
	"github.com/go-redis/redis/v9"
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

type SessionToken string

// Session represents an authentication session on the server side, containing the user ID and their authorization roles.
type Session struct {
	Token     SessionToken     // Session ID of the logged-in user.
	UserID    pulid.PULID      // ID of the user.
	AuthRoles []authrole.Value // Auth roles of the user.
}

// NewSessionToken returns a new SessionToken.
func NewSessionToken() (SessionToken, error) {
	b := make([]byte, 20)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("NewSessionToken(): could not generate a new session ID")
	}
	return SessionToken(hex.EncodeToString(b)), nil
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
func HasAuthRoles(session *Session, authRoles ...authrole.Value) bool {
	if session == nil {
		return false
	}

	// Make session auth roles into a map for easier searching.
	var m map[authrole.Value]bool
	for _, sessionAuthRole := range session.AuthRoles {
		m[sessionAuthRole] = true
	}

	// Search for existence of the authRole value.
	for _, authRole := range authRoles {
		if m[authRole] {
			return true
		}
	}
	return false
}

// IsLoggedIn checks if the user is logged in (has a session).
func IsLoggedIn(session *Session) bool {
	if session == nil || session.Token == "" || session.UserID == "" || len(session.AuthRoles) == 0 {
		return false
	}

	_, err := ParseSessionToken(session.Token.String())
	if err != nil {
		return false
	}

	return true
}

// createSession makes a new Session instance with a random session ID.
func (authService AuthService) createSession(
	ctx context.Context, userID pulid.PULID, authRoles ...authrole.Value,
) (*Session, error) {
	sessionToken, err := NewSessionToken()
	if err != nil {
		return nil, fmt.Errorf("auth.createSession(): %v", err)
	}

	newSession := Session{
		Token:     sessionToken,
		UserID:    userID,
		AuthRoles: authRoles,
	}

	// Store newly created session in Redis.
	err = authService.setSessionInStore(ctx, &newSession)
	if err != nil {
		return nil, fmt.Errorf("auth.createSession(): %v", err)
	}

	return &newSession, nil
}

// setSessionInStore stringifies a Session object and add it to Redis under the session key.
func (authService AuthService) setSessionInStore(ctx context.Context, session *Session) error {
	expiration := authService.Config().SessionTimeout

	// Marshal session into JSON.
	bytes, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf(
			"%w: auth.setSessionInStore(): unable to serialize session for Redis storage: %v", ErrInternal, err,
		)
	}

	// Store in Redis as JSON string.
	_, err = authService.redisService.Client().Set(
		ctx,
		fmt.Sprintf("%s:%s", SessionRedisKey, session.Token), bytes, expiration,
	).Result()
	if err != nil {
		return fmt.Errorf(
			"%w: auth.setSessionInStore(): could not store session in Redis: %v", ErrInternal, err,
		)
	}

	return nil
}

// GetSessionFromStore retrieves a Session from Redis given the key and session ID. It also resets the expiry time of the session.
func (authService AuthService) GetSessionFromStore(ctx context.Context, sessionToken SessionToken) (*Session, error) {
	// Retrieve session from Redis.
	sessionJSON, err := authService.redisService.Client().Get(
		ctx, fmt.Sprintf("%s:%s", SessionRedisKey, sessionToken.String()),
	).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("%w: auth.GetSessionFromStore(): session does not exist", ErrNotFound)
		}
		return nil, fmt.Errorf(
			"%w: auth.GetSessionFromStore(): could not retrieve session from Redis: %v", ErrInternal, err,
		)
	}

	// Unmarshal the JSON.
	var freshSession Session
	err = json.Unmarshal([]byte(sessionJSON), &freshSession)
	if err != nil {
		return nil, fmt.Errorf("%w: auth.GetSessionFromStore(): could not unmarshal session: %v", ErrInternal, err)
	}

	// Reset expiry on the session.
	expiration := authService.Config().SessionTimeout

	_, err = authService.redisService.Client().Expire(
		ctx, fmt.Sprintf("%s:%s", SessionRedisKey, freshSession.Token.String()), expiration,
	).Result()
	if err != nil {
		return nil, fmt.Errorf(
			"%w: auth.GetSessionFromStore(): could not reset expiry on session: %v", ErrInternal, err,
		)
	}

	return &freshSession, nil
}

// ParseSessionToken parses a session ID string and verifies that it is correct.
func ParseSessionToken(sessionTokenStr string) (SessionToken, error) {
	if sessionTokenStr == "" || len(sessionTokenStr) != 40 {
		return "", fmt.Errorf("%w: auth.ParseSessionToken(): session id string length is not correct", ErrBadInput)
	}
	return SessionToken(sessionTokenStr), nil
}

// String implements fmt.Stringer interface.
func (sessionToken SessionToken) String() string {
	return string(sessionToken)
}

// MarshalGQL implements graphql.Marshaler interface.
func (sessionToken SessionToken) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(sessionToken.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (sessionToken *SessionToken) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("%w: session ID %T must be a string", ErrBadInput, val)
	}

	*sessionToken = SessionToken(str)
	return nil
}
