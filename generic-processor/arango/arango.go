package arango

import (
	"context"
	"errors"
	"log"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type ArangoConfig struct {
	URL      string
	User     string
	Password string
	Database string
}

type ArangoConn struct {
	Db driver.Database
}

func NewArangoConnection(cfg ArangoConfig) (*ArangoConn, error) {
	// check configuration
	if cfg.URL == "" || cfg.User == "" || cfg.Password == "" || cfg.Database == "" {
		return nil, errors.New("ArangoDB config has empty field")
	}

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{cfg.URL},
	})
	if err != nil {
		return nil, errors.New("failed to create HTTP connection")
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(cfg.User, cfg.Password),
	})
	if err != nil {
		return nil, errors.New("failed to create ArangoDB client")
	}

	db, err := client.Database(context.TODO(), cfg.Database)
	if err != nil {
		return nil, errors.New("failed to open database")
	}

	return &ArangoConn{
		Db: db,
	}, nil

}

func (arangoConn *ArangoConn) GetCollection(name string) driver.Collection {
	exists, err := arangoConn.Db.CollectionExists(context.TODO(), name)
	if err != nil {
		log.Fatalf("Error checking collection %s: %v", name, err)
	}

	var col driver.Collection

	if exists {
		col, err = arangoConn.Db.Collection(context.TODO(), name)
		if err != nil {
			log.Fatalf("Error opening collection %s: %v", name, err)
		}
	} else {
		options := driver.CreateCollectionOptions{}
		col, err = arangoConn.Db.CreateCollection(context.TODO(), name, &options)
		if err != nil {
			log.Fatalf("Error creating collection %s: %v", name, err)
		}
	}

	return col
}
