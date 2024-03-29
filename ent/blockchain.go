// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chenningg/hermitboard-api/ent/blockchain"
	"github.com/chenningg/hermitboard-api/pulid"
)

// Blockchain is the model entity for the Blockchain schema.
type Blockchain struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.PULID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Symbol holds the value of the "symbol" field.
	Symbol string `json:"symbol,omitempty"`
	// Icon holds the value of the "icon" field.
	Icon *string `json:"icon,omitempty"`
	// ChainID holds the value of the "chain_id" field.
	ChainID *int64 `json:"chainID,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BlockchainQuery when eager-loading is set.
	Edges BlockchainEdges `json:"edges"`
}

// BlockchainEdges holds the relations/edges for other nodes in the graph.
type BlockchainEdges struct {
	// Cryptocurrencies holds the value of the cryptocurrencies edge.
	Cryptocurrencies []*Cryptocurrency `json:"cryptocurrencies,omitempty"`
	// Transactions holds the value of the transactions edge.
	Transactions []*Transaction `json:"transactions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedCryptocurrencies map[string][]*Cryptocurrency
	namedTransactions     map[string][]*Transaction
}

// CryptocurrenciesOrErr returns the Cryptocurrencies value or an error if the edge
// was not loaded in eager-loading.
func (e BlockchainEdges) CryptocurrenciesOrErr() ([]*Cryptocurrency, error) {
	if e.loadedTypes[0] {
		return e.Cryptocurrencies, nil
	}
	return nil, &NotLoadedError{edge: "cryptocurrencies"}
}

// TransactionsOrErr returns the Transactions value or an error if the edge
// was not loaded in eager-loading.
func (e BlockchainEdges) TransactionsOrErr() ([]*Transaction, error) {
	if e.loadedTypes[1] {
		return e.Transactions, nil
	}
	return nil, &NotLoadedError{edge: "transactions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Blockchain) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blockchain.FieldID:
			values[i] = new(pulid.PULID)
		case blockchain.FieldChainID:
			values[i] = new(sql.NullInt64)
		case blockchain.FieldName, blockchain.FieldSymbol, blockchain.FieldIcon:
			values[i] = new(sql.NullString)
		case blockchain.FieldCreatedAt, blockchain.FieldUpdatedAt, blockchain.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Blockchain", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Blockchain fields.
func (b *Blockchain) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blockchain.FieldID:
			if value, ok := values[i].(*pulid.PULID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case blockchain.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case blockchain.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case blockchain.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				b.DeletedAt = new(time.Time)
				*b.DeletedAt = value.Time
			}
		case blockchain.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		case blockchain.FieldSymbol:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field symbol", values[i])
			} else if value.Valid {
				b.Symbol = value.String
			}
		case blockchain.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				b.Icon = new(string)
				*b.Icon = value.String
			}
		case blockchain.FieldChainID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field chain_id", values[i])
			} else if value.Valid {
				b.ChainID = new(int64)
				*b.ChainID = value.Int64
			}
		}
	}
	return nil
}

// QueryCryptocurrencies queries the "cryptocurrencies" edge of the Blockchain entity.
func (b *Blockchain) QueryCryptocurrencies() *CryptocurrencyQuery {
	return (&BlockchainClient{config: b.config}).QueryCryptocurrencies(b)
}

// QueryTransactions queries the "transactions" edge of the Blockchain entity.
func (b *Blockchain) QueryTransactions() *TransactionQuery {
	return (&BlockchainClient{config: b.config}).QueryTransactions(b)
}

// Update returns a builder for updating this Blockchain.
// Note that you need to call Blockchain.Unwrap() before calling this method if this Blockchain
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Blockchain) Update() *BlockchainUpdateOne {
	return (&BlockchainClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Blockchain entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Blockchain) Unwrap() *Blockchain {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Blockchain is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Blockchain) String() string {
	var builder strings.Builder
	builder.WriteString("Blockchain(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := b.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(b.Name)
	builder.WriteString(", ")
	builder.WriteString("symbol=")
	builder.WriteString(b.Symbol)
	builder.WriteString(", ")
	if v := b.Icon; v != nil {
		builder.WriteString("icon=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := b.ChainID; v != nil {
		builder.WriteString("chain_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// NamedCryptocurrencies returns the Cryptocurrencies named value or an error if the edge was not
// loaded in eager-loading with this name.
func (b *Blockchain) NamedCryptocurrencies(name string) ([]*Cryptocurrency, error) {
	if b.Edges.namedCryptocurrencies == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := b.Edges.namedCryptocurrencies[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (b *Blockchain) appendNamedCryptocurrencies(name string, edges ...*Cryptocurrency) {
	if b.Edges.namedCryptocurrencies == nil {
		b.Edges.namedCryptocurrencies = make(map[string][]*Cryptocurrency)
	}
	if len(edges) == 0 {
		b.Edges.namedCryptocurrencies[name] = []*Cryptocurrency{}
	} else {
		b.Edges.namedCryptocurrencies[name] = append(b.Edges.namedCryptocurrencies[name], edges...)
	}
}

// NamedTransactions returns the Transactions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (b *Blockchain) NamedTransactions(name string) ([]*Transaction, error) {
	if b.Edges.namedTransactions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := b.Edges.namedTransactions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (b *Blockchain) appendNamedTransactions(name string, edges ...*Transaction) {
	if b.Edges.namedTransactions == nil {
		b.Edges.namedTransactions = make(map[string][]*Transaction)
	}
	if len(edges) == 0 {
		b.Edges.namedTransactions[name] = []*Transaction{}
	} else {
		b.Edges.namedTransactions[name] = append(b.Edges.namedTransactions[name], edges...)
	}
}

// Blockchains is a parsable slice of Blockchain.
type Blockchains []*Blockchain

func (b Blockchains) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
