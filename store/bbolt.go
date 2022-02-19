package store

import (
	"github.com/amirkhaki/cmuxbtr/config"
	"github.com/amirkhaki/cmuxbtr/protocol"
	"fmt"
	"context"
	"time"
	bolt "go.etcd.io/bbolt"
)


var bucketName = []byte("wp-amazon")
var Storage *boltStore
var _ protocol.Store = Storage


type boltStore struct {
	db *bolt.DB
}

func ( s *boltStore ) Set(ctx context.Context, key, value []byte ) error {
	err := s.db.Update(func(txn *bolt.Tx) error {
		b, err := txn.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return fmt.Errorf("Could not retrieve or create bucket: %w", err)
		}
		if err := b.Put(key, value); err != nil {
			return fmt.Errorf("Could not set key %x: %w", key, err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("Could not update db: %w", err)
	}
	return nil
}

func ( s *boltStore ) Get(ctx context.Context, key []byte ) (value []byte, _ error) {
	err := s.db.View(func(txn *bolt.Tx) error {
		b := txn.Bucket(bucketName)
		if b == nil {
			return fmt.Errorf("Could not retrieve bucket")
		}
		val := b.Get(key)
		if val == nil {
			return fmt.Errorf("Could not get key %x", key)
		}
		value = append(value, val...)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Could not get key %x: %w", key, err)
	}
	return value, nil
}

func ( s *boltStore ) Update( ctx context.Context, key, value []byte ) error {
	if err := s.Set(ctx, key, value); err != nil {
		return fmt.Errorf("Could not upfate key %x: %w", key, err)
	}
	return nil
}

func ( s *boltStore ) Delete(ctx context.Context, key []byte ) error {
	err := s.db.Update(func(txn *bolt.Tx) error {
		b := txn.Bucket(bucketName)
		if b == nil {
			return fmt.Errorf("Could not retrieve bucket")
		}
		if err := b.Delete(key); err != nil {
			return fmt.Errorf("Could not delete key %x: %w", key, err)		
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("Could not update db: %w", err)
	}
	return nil
}

func ( s *boltStore ) Keys(ctx context.Context) <-chan []byte {
	c := make(chan []byte)
	go func(){
		defer close(c)
		err := s.db.View(func(txn *bolt.Tx) error {
			b := txn.Bucket(bucketName)
			if b == nil {
				return fmt.Errorf("Could not retrieve bucket")
			}
			b.ForEach(func(k, v []byte) error {
				key := make([]byte, len(k))
				copy(key, k)
				select {
				case c <- k:
					return nil
				case <- ctx.Done():
					return ctx.Err()
				}
			})
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	}()
	return c

}

// Connect to store
func Connect( cfg *config.Config ) ( error ) {
	fmt.Println("connect called")
	db, err := bolt.Open(cfg.DBPath, 0644, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return fmt.Errorf("Could not open database: %w", err)
	}
	Storage = &boltStore{db: db}
	return nil

}

func Close() {
	Storage.db.Close()
}
