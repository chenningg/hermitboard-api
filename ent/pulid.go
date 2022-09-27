package ent

import (
    "context"
    "fmt"

    "github.com/chenningg/hermitboard-api/ent/account"
    "github.com/chenningg/hermitboard-api/ent/asset"
    "github.com/chenningg/hermitboard-api/ent/assetclass"
    "github.com/chenningg/hermitboard-api/ent/authrole"
    "github.com/chenningg/hermitboard-api/ent/authtype"
    "github.com/chenningg/hermitboard-api/ent/blockchain"
    "github.com/chenningg/hermitboard-api/ent/connection"
    "github.com/chenningg/hermitboard-api/ent/cryptocurrency"
    "github.com/chenningg/hermitboard-api/ent/dailyassetprice"
    "github.com/chenningg/hermitboard-api/ent/exchange"
    "github.com/chenningg/hermitboard-api/ent/portfolio"
    "github.com/chenningg/hermitboard-api/ent/source"
    "github.com/chenningg/hermitboard-api/ent/sourcetype"
    "github.com/chenningg/hermitboard-api/ent/staffaccount"
    "github.com/chenningg/hermitboard-api/ent/transaction"
    "github.com/chenningg/hermitboard-api/ent/transactiontype"
    "github.com/chenningg/hermitboard-api/pulid"
)

// prefixMap maps pulid.PULID prefixes to table names.
var prefixMap = map[string]string{
    "ACC": account.Table,
    "AST": asset.Table,
    "ASC": assetclass.Table,
    "ATR": authrole.Table,
    "AHT": authtype.Table,
    "BKC": blockchain.Table,
    "CNN": connection.Table,
    "CRP": cryptocurrency.Table,
    "DAP": dailyassetprice.Table,
    "EXG": exchange.Table,
    "PTF": portfolio.Table,
    "SRC": source.Table,
    "SCT": sourcetype.Table,
    "SAC": staffaccount.Table,
    "TRX": transaction.Table,
    "TRT": transactiontype.Table,
}

// IDToType maps a pulid.PULID to the underlying table.
func IDToType(ctx context.Context, id pulid.PULID) (string, error) {
    err := pulid.IsAPULID(id)
    if err != nil {
        return "", fmt.Errorf("IDToType(): invalid PULID: %v", err)
    }

    prefix, _, err := pulid.ParsePULID(id)
    if err != nil {
        return "", fmt.Errorf("IDToType(): unable to parse PULID: %v", err)
    }

    typ := prefixMap[prefix]
    if typ == "" {
        return "", fmt.Errorf("IDToType(): could not map prefix '%s' to a type", prefix)
    }
    return typ, nil
}
