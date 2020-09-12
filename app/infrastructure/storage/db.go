package storage

import (
	"context"
	"errors"
	"fmt"
	"sales/config"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

var (
	errorConnectDB    = errors.New("error connect database")
	errorRunQuery     = errors.New("error run query")
	errorRunMutuation = errors.New("error run mutuation")
	errorAlterSchema  = errors.New("error alter schema")
	errorDeleteData   = errors.New("error delete data")
)

type Storage struct {
	config *config.DBConfig
}

func NewStorage(config *config.DBConfig) *Storage {
	return &Storage{
		config: config,
	}
}

var db *dgo.Dgraph
var ctx = context.Background()

func (s *Storage) Connect() error {
	config := s.config
	conn, err := grpc.Dial(config.Host+config.Port, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return errorConnectDB
	}
	fmt.Println("data base connected...")
	db = dgo.NewDgraphClient(api.NewDgraphClient(conn))
	return nil
}

func (s *Storage) RunQuery(query string) ([]byte, error) {
	txn := db.NewTxn()
	defer txn.Discard(ctx)
	res, err := txn.Query(ctx, query)
	if err != nil {
		fmt.Println(err)
		return nil, errorRunQuery
	}
	return res.GetJson(), nil
}

func (s *Storage) RunMutation(mutuation []byte) error {
	txn := db.NewTxn()
	defer txn.Discard(ctx)
	mu := &api.Mutation{
		CommitNow: true,
	}
	mu.SetJson = mutuation
	_, err := txn.Mutate(ctx, mu)
	if err != nil {
		fmt.Println(err)
		return errorRunMutuation
	}
	return nil
}

func (s *Storage) DeleteAllPredicate() error {
	err := db.Alter(ctx, &api.Operation{DropOp: api.Operation_ALL})
	//err := db.Alter(ctx, &api.Operation{DropAll: true})
	if err != nil {
		fmt.Println(err)
		return errorDeleteData
	}
	alterSchema()
	return nil
}

func setup(schema string) error {
	err := db.Alter(ctx, &api.Operation{
		Schema: schema,
	})
	if err != nil {
		fmt.Println(err)
		return errorAlterSchema
	}
	return nil
}

func alterSchema() {
	schema := `
		id: string @index(term) .
		name: string @index(term) .
		created_at: string @index(term) .
		productId: string @index(term) .
		transactionId: string @index(term) .
		ipAddress: string @index(term) .
		buyerId: string @index(term) .
		device: string @index(term) .
	`
	setup(schema)
}
