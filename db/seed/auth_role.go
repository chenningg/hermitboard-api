package seed

import (
	"context"
	"fmt"
	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/authrole"
)

func seedAuthRoles(ctx context.Context, client *ent.Client) error {
	seedAuthRoles := []authrole.AuthRole{
		authrole.AuthRoleDemo, authrole.AuthRoleFree, authrole.AuthRolePlus, authrole.AuthRolePro,
		authrole.AuthRoleEnterprise, authrole.AuthRoleSupport, authrole.AuthRoleAdmin, authrole.AuthRoleSuperAdmin,
	}
	seedDescriptions := []string{
		"A demo role with restricted rights.", "A free role with limited feature access.", "A paying Plus role.",
		"A paying Pro role.", "A customer support role with limited permissions.",
		"An administrative role with elevated permissions.", "A super administrative role with full rights.",
	}

	bulk := make([]*ent.AuthRoleCreate, len(seedAuthRoles))
	for i := 0; i < len(seedAuthRoles); i++ {
		bulk[i] = client.AuthRole.Create().SetAuthRole(seedAuthRoles[i]).SetDescription(seedDescriptions[i])
	}

	// Create in bulk
	_, err := client.AuthRole.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return fmt.Errorf("seedAuthRoles(): %v", err)
	}

	return nil
}
