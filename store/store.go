package store

import (
	"fmt"
	"context"
	badger "github.com/dgraph-io/badger/v3"
)




type badgerStore struct {
	db *badger.DB
}

func ( s *badgerStore ) Set(ctx context.Context, key, value []byte ) error {
	err := s.db.Update(func(txn *badger.Txn) error {
		if err := txn.Set(key, value); err != nil {
			return fmt.Errorf("Could not set key %x: %w", key, err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("Could not update db: %w", err)
	}
	return nil
}

func ( s *badgerStore ) Get(ctx context.Context, key []byte ) (value []byte, _ error) {
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return fmt.Errorf("Could not get key %x: %w", key, err)
		}
		err = item.Value(func(val []byte) error {
			value = append(value, val...)
			return nil
		})
		if err != nil {
			return fmt.Errorf("Could not get item value: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Could not get key %x: %w", key, err)
	}
	return value, nil
}

func ( s *badgerStore ) Update( ctx context.Context, key, value []byte ) error {
	if err := s.Set(ctx, key, value); err != nil {
		return fmt.Errorf("Could not upfate key %x: %w", key, err)
	}
	return nil
}

func ( s *badgerStore ) Delete(ctx context.Context, key []byte ) error {
	err := s.db.Update(func(txn *badger.Txn) error {
		if err := txn.Delete(key); err != nil {
			return fmt.Errorf("Could not delete key %x: %w", key, err)		
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("Could not update db: %w", err)
	}
	return nil
}

func ( s *badgerStore ) Keys(ctx context.Context) <-chan []byte {
	c := make(chan []byte)
	go func(){
		defer close(c)
		err := s.db.View(func(txn *badger.Txn) error {
			opts := badger.DefaultIteratorOptions
			opts.PrefetchSize = 10
			it := txn.NewIterator(opts)
			defer it.Close()
			for it.Rewind(); it.Valid(); it.Next() {
				item := it.Item()
				select {
				case c <- item.Key():
				case <- ctx.Done():
					return ctx.Err()
				}
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	}()
	return c

}

// Connect to store
/*
func Connect( cfg *config.Config ) ( error ) {
	fmt.Println("connect called")
	db, err := badger.Open(badger.DefaultOptions(cfg.DBPath))
	if err != nil {
		return fmt.Errorf("Could not open database: %w", err)
	}
	Storage = &badgerStore{db: db}
	return nil

}
*/
