package database

// Why all the mutex stuff? Seb thought he'd be smart
// and try make use of the Singleton design pattern
// in Go whilst he's at it.

// Time spent trying: 2 hours

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

var (
	instance *DB
	lock     sync.Mutex
)

func GetDB(user, password, host, port, dbname string) *DB {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			log.Println(time.Now(), "Creating initial DB connection")

			connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",user, password, host, port, dbname)

			conn, err := sql.Open("postgres", connStr)
			if err != nil {
				log.Fatalf("DB connection failed: %v", err)
			}

			if err = conn.Ping(); err != nil {
				log.Fatalf("DB ping failed: %v", err)
			}

			instance = &DB{conn: conn}
		} else {
			log.Println(time.Now(), "Redundant DB instance request")
		}
	} else {
		log.Println(time.Now(), "Redundant DB instance request")
	}

	return instance
}

func (db *DB) Conn() *sql.DB {
	return db.conn
}
