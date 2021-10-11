package adapters

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
)

type Storage struct {
	cli *storage.Client
	ctx context.Context
}

func MustNewGCS(cli *storage.Client, ctx context.Context) *Storage {
	return &Storage{
		cli: cli,
		ctx: ctx,
		// gcs: gcs,
	}
}

// func NewStorage() *Storage {
// 	ctx := context.Background()

// 	client, err := storage.NewClient(ctx)
// 	if err != nil {
// 		fmt.Println("Could not connect to the Google Cloud Storage")
// 	}

// 	if err := client.Close(); err != nil {
// 		fmt.Println("Could not close the GCS")
// 	}

// 	return &Storage{
// 		cli: client,
// 		ctx: ctx,
// 	}
// }

func (d *Storage) ListBucket() {
	bucket := d.cli.Bucket("tradulab-files")

	fmt.Println(d, "----------------------------")

	storage22 := bucket.Object("oi.txt")

	attrs, _ := storage22.Attrs(d.ctx)
	fmt.Println(attrs)

	// query := &storage.Query{Prefix: ""}
	// it := bucket.Objects(d.ctx, query)
	// for {
	// 	obj, err := it.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(obj)
	// }
}
