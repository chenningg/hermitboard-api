package seed

import (
	"context"
	"fmt"
	"github.com/chenningg/hermitboard-api/ent"
	"github.com/chenningg/hermitboard-api/ent/authrole"
)

func seedAuthRoles(ctx context.Context, client *ent.Client) error {
	seedAuthRoles := []authrole.Value{
		authrole.ValueDemo, authrole.ValueFree, authrole.ValuePlus, authrole.ValuePro,
		authrole.ValueEnterprise, authrole.ValueSupport, authrole.ValueAdmin, authrole.ValueSuperAdmin,
	}
	seedDescriptions := []string{
		"A demo role with restricted rights.", "A free role with limited feature access.", "A paying Plus role.",
		"A paying Pro role.", "An Enterprise role, generally with a contract.",
		"A customer support role with limited permissions.",
		"An administrative role with elevated permissions.", "A super administrative role with full rights.",
	}

	bulk := make([]*ent.AuthRoleCreate, len(seedAuthRoles))
	for i := 0; i < len(seedAuthRoles); i++ {
		bulk[i] = client.AuthRole.Create().SetValue(seedAuthRoles[i]).SetDescription(seedDescriptions[i])
	}

	// Create in bulk
	_, err := client.AuthRole.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return fmt.Errorf("seedAuthRoles(): %v", err)
	}

	return nil
}
