package storage

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	jww "github.com/spf13/jwalterweatherman"
	"sync"
	"time"
)

var UserDiscoveryDB Storage

// Interface declaration for storage methods
type Storage interface {
}

// Struct implementing the Database Interface with an underlying DB
type DatabaseImpl struct {
	db *gorm.DB // Stored database connection
}

// Struct implementing the Database Interface with an underlying Map
type MapImpl struct {
	sync.RWMutex
}

type User struct {
	Id        []byte `gorm:"primary_key"`
	RsaPub    []byte
	DhPub     []byte
	Salt      []byte
	Signature []byte
	Facts     []Fact `gorm:"foreignKey:UserId"`
}

type Fact struct {
	Id                 uint64 `gorm:"primary_key;AUTO_INCREMENT:true"`
	UserId             []byte `gorm:"NOT NULL"`
	Fact               string `gorm:"NOT NULL"`
	FactType           uint64 `gorm:"NOT NULL"`
	FactHash           []byte `gorm:"NOT NULL"`
	Signature          []byte `gorm:"NOT NULL"`
	ConfirmationId     []byte `gorm:"UNIQUE;NOT NULL"`
	VerificationStatus uint64 `gorm:"NOT NULL"`
	Manual             bool   `gorm:"NOT NULL"`
	Code               uint64
	Timestamp          time.Time `gorm:"NOT NULL"`
}

// Initialize the Database interface with database backend
// Returns a Storage interface, Close function, and error
func NewDatabase(username, password, database, address,
	port string) (Storage, func() error, error) {
	var err error
	var db *gorm.DB
	//connect to the database if the correct information is provided
	if address != "" && port != "" {
		// Create the database connection
		connectString := fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable",
			address, port, username, database)
		// Handle empty database password
		if len(password) > 0 {
			connectString += fmt.Sprintf(" password=%s", password)
		}
		db, err = gorm.Open("postgres", connectString)
	}

	// Return the map-backend interface
	// in the event there is a database error or information is not provided
	if (address == "" || port == "") || err != nil {

		if err != nil {
			jww.WARN.Printf("Unable to initialize database backend: %+v", err)
		} else {
			jww.WARN.Printf("Database backend connection information not provided")
		}

		defer jww.INFO.Println("Map backend initialized successfully!")

		mapImpl := &MapImpl{}

		return Storage(mapImpl), func() error { return nil }, nil
	}

	// Initialize the database logger
	db.SetLogger(jww.TRACE)
	db.LogMode(true)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(24 * time.Hour)

	// Initialize the database schema
	// WARNING: Order is important. Do not change without database testing
	models := []interface{}{User{}, Fact{}}
	for _, model := range models {
		err = db.AutoMigrate(model).Error
		if err != nil {
			return Storage(&DatabaseImpl{}), func() error { return nil }, err
		}
	}

	// Build the interface
	di := &DatabaseImpl{
		db: db,
	}

	jww.INFO.Println("Database backend initialized successfully!")
	return Storage(di), db.Close, nil
}
