package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // required for SQL access
	migrate "github.com/rubenv/sql-migrate"
)

// Config defines the options that are used when connecting to a PostgreSQL instance
type Config struct {
	Host        string
	Port        string
	PortRead    string
	PortWrite   string
	User        string
	Pass        string
	Name        string
	SSLMode     string
	SSLCert     string
	SSLKey      string
	SSLRootCert string
}

// func Connect(cfg Config) (*sqlx.DB, error) {
// 	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s sslcert=%s sslkey=%s sslrootcert=%s", cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Pass, cfg.SSLMode, cfg.SSLCert, cfg.SSLKey, cfg.SSLRootCert)
// 	db, err := sqlx.Open("postgres", url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err := migrateDB(db.DB); err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

func ConnectRead(cfg Config) (*sqlx.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s sslcert=%s sslkey=%s sslrootcert=%s \n", cfg.Host, cfg.PortRead, cfg.User, cfg.Name, cfg.Pass, cfg.SSLMode, cfg.SSLCert, cfg.SSLKey, cfg.SSLRootCert)
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectWrite(cfg Config) (*sqlx.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s sslcert=%s sslkey=%s sslrootcert=%s \n", cfg.Host, cfg.PortWrite, cfg.User, cfg.Name, cfg.Pass, cfg.SSLMode, cfg.SSLCert, cfg.SSLKey, cfg.SSLRootCert)
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	if err := migrateDB(db.DB); err != nil {
		return nil, err
	}
	return db, nil
}

func migrateDB(db *sql.DB) error {
	migrations := &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			{
				Id: "user_table",
				Up: []string{
					`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
					`CREATE TABLE IF NOT EXISTS user (
						id             	UUID            DEFAULT uuid_generate_v4() PRIMARY KEY,
						created_at     	TIMESTAMP       DEFAULT NOW(),
						updated_at     	TIMESTAMP       DEFAULT NOW(),
						username       	VARCHAR(254)    NOT NULL,
						password       	VARCHAR(254)    NOT NULL,
						phone       	VARCHAR(20)    	NOT NULL,
						role       		VARCHAR(20)    	NOT NULL,
						status       	VARCHAR(20)    	NOT NULL
					)`,
				},
				Down: []string{
					`DROP TABLE user`,
				},
			},
			{
				Id: "game_table",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS game (
						id             	UUID            DEFAULT uuid_generate_v4() PRIMARY KEY,
						created_at     	TIMESTAMP       DEFAULT NOW(),
						updated_at     	TIMESTAMP       DEFAULT NOW(),
						name       		VARCHAR(254)    NOT NULL,
						images       	TEXT            NOT NULL,
						type       		VARCHAR(20)    	NOT NULL,
						exchange_allow  BOOLEAN         NOT NULL,
						tutorial       	TEXT         	NOT NULL,
						user_id       	UUID            NOT NULL REFERENCES user(id)
					)`,
				},
				Down: []string{
					`DROP TABLE game`,
				},
			},
		},
	}

	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	return err
}
