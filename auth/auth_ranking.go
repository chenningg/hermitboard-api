package auth

import "github.com/chenningg/hermitboard-api/ent/authrole"

var AuthRoleRanking = map[authrole.Value]int{
	authrole.ValueDemo:       1,
	authrole.ValueFree:       1,
	authrole.ValuePlus:       1,
	authrole.ValuePro:        1,
	authrole.ValueEnterprise: 1,
	authrole.ValueSupport:    50,
	authrole.ValueAdmin:      99,
	authrole.ValueSuperAdmin: 100,
}

// IsHigherAuthority checks that the authenticated user has enough authority over another set of auth roles.
func IsHigherAuthority(session *Session, against ...authrole.Value) bool {
	myHighest := -1
	againstHighest := 0
	for _, myAuthRole := range session.AuthRoles {
		if myAuthority, ok := AuthRoleRanking[myAuthRole]; ok {
			if myAuthority > myHighest {
				myHighest = myAuthority
			}
		} else {
			// Invalid auth role supplied.
			return false
		}
	}
	for _, againstAuthRole := range against {
		if againstAuthority, ok := AuthRoleRanking[againstAuthRole]; ok {
			if againstAuthority > myHighest {
				againstHighest = againstAuthority
			}
		} else {
			// Invalid auth role supplied.
			return false
		}
	}

	if myHighest > againstHighest {
		return true
	}
	return false
}
