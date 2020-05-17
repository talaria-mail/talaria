package kv

import (
	"context"

	bolt "go.etcd.io/bbolt"
)

type boltKV struct {
	*bolt.DB
	bucket []byte
}

// NewBoltStore is a simple boltdb based key-value store
func NewBoltStore(path, bucket string) (Store, error) {
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})

	return &boltKV{db, []byte(bucket)}, nil
}

func (b *boltKV) Get(ctx context.Context, key string) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		var output []byte
		err := b.View(func(tx *bolt.Tx) error {
			bkt := tx.Bucket(b.bucket)
			value := bkt.Get([]byte(key))
			output = make([]byte, len(value))
			copy(output, value)
			return nil
		})
		if err != nil {
			return nil, ErrorNotFound
		}
		return output, nil
	}
}

func (b *boltKV) Put(ctx context.Context, key string, value []byte) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		err := b.Update(func(tx *bolt.Tx) error {
			bkt := tx.Bucket(b.bucket)
			err := bkt.Put([]byte(key), value)
			return err
		})
		return err
	}
}

func (b *boltKV) Delete(ctx context.Context, key string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		err := b.Update(func(tx *bolt.Tx) error {
			bkt := tx.Bucket(b.bucket)
			return bkt.Delete([]byte(key))
		})
		return err
	}
}
