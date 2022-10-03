package connection

import "github.com/chenningg/hermitboard-api/pulid"

type DeleteConnectionInput struct {
	ID pulid.PULID `json:"id"`
}
