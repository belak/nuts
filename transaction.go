package nuts

import "github.com/boltdb/bolt"

type Tx struct {
	raw *bolt.Tx
}

func (tx *Tx) Raw() *bolt.Tx {
	return tx.raw
}

func (tx *Tx) Bucket(name string) *Bucket {
	rawBucket := tx.raw.Bucket([]byte(name))
	if rawBucket == nil {
		return nil
	}

	return &Bucket{
		raw: rawBucket,
	}
}

func (tx *Tx) CreateBucketIfNotExists(name string) (*Bucket, error) {
	rawBucket, err := tx.raw.CreateBucketIfNotExists([]byte(name))
	if err != nil {
		return nil, err
	}

	return &Bucket{
		raw: rawBucket,
	}, nil
}

// func (tx *Tx) Bucket(name []byte) *Bucket
// func (tx *Tx) Check() <-chan error
// func (tx *Tx) Commit() error
// func (tx *Tx) Copy(w io.Writer) error
// func (tx *Tx) CopyFile(path string, mode os.FileMode) error
// func (tx *Tx) CreateBucket(name []byte) (*Bucket, error)
// func (tx *Tx) Cursor() *Cursor
// func (tx *Tx) DB() *DB
// func (tx *Tx) DeleteBucket(name []byte) error
// func (tx *Tx) ForEach(fn func(name []byte, b *Bucket) error) error
// func (tx *Tx) ID() int
// func (tx *Tx) OnCommit(fn func())
// func (tx *Tx) Page(id int) (*PageInfo, error)
// func (tx *Tx) Rollback() error
// func (tx *Tx) Size() int64
// func (tx *Tx) Stats() TxStats
// func (tx *Tx) Writable() bool
// func (tx *Tx) WriteTo(w io.Writer) (n int64, err error)
