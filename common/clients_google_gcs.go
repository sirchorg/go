package common

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

func (self *GCPClients) GetObjectAndUnmarshal(bucket *storage.BucketHandle, objectName string, dst interface{}) error {
	b, err := self.GetObject(bucket, objectName)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

func (self *GCPClients) GetObject(bucket *storage.BucketHandle, objectName string) ([]byte, error) {
	r, err := bucket.Object(objectName).NewReader(context.Background())
	if err != nil {
		return nil, fmt.Errorf("storage.GetObject: %w", err)
	}
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("storage.GetObject: %w", err)
	}
	if err := r.Close(); err != nil {
		return nil, fmt.Errorf("storage.GetObject: %w", err)
	}
	return b, nil
}
