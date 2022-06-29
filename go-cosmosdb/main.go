package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	endpoint := os.Getenv("AZURE_COSMOS_ENDPOINT")
	if endpoint == "" {
		return errors.New("AZURE_COSMOS_ENDPOINT not set")
	}

	key := os.Getenv("AZURE_COSMOS_KEY")
	if key == "" {
		return errors.New("AZURE_COSMOS_KEY not set")
	}

	database := os.Getenv("AZURE_COSMOS_DATABASE")
	if database == "" {
		database = "my-db"
	}

	container := os.Getenv("AZURE_COSMOS_CONTAINER")
	if container == "" {
		container = "my-container"
	}

	partitionKey := os.Getenv("AZURE_COSMOS_PARTITIONKEY")
	if partitionKey == "" {
		partitionKey = "/pk"
	}

	cred, err := azcosmos.NewKeyCredential(key)
	if err != nil {
		return err
	}

	client, err := azcosmos.NewClientWithKey(endpoint, cred, nil)
	if err != nil {
		return err
	}

	ignoreConflict := func(err error) error {
		var responseErr *azcore.ResponseError
		switch {
		case err != nil && errors.As(err, &responseErr) && responseErr.ErrorCode == "Conflict":
			log.Printf("Conflict")
			return nil
		default:
			return err
		}
	}

	databaseClient, err := client.NewDatabase(database)
	if ignoreConflict(err) != nil {
		return err
	}

	databaseProperties := azcosmos.DatabaseProperties{ID: database}

	ctx := context.Background()

	databaseResponse, err := client.CreateDatabase(ctx, databaseProperties, nil)
	if ignoreConflict(err) != nil {
		return err
	}
	_ = databaseResponse

	properties := azcosmos.ContainerProperties{
		ID: container,
		PartitionKeyDefinition: azcosmos.PartitionKeyDefinition{
			Paths: []string{partitionKey},
		},
	}

	containerResponse, err := databaseClient.CreateContainer(ctx, properties, nil)
	if ignoreConflict(err) != nil {
		return err
	}
	log.Printf("CreateContainer | RequestCharge: %v", containerResponse.RequestCharge)

	containerClient, err := client.NewContainer(database, container)
	if ignoreConflict(err) != nil {
		return err
	}

	item := struct {
		ID          string `json:"id"`
		PK          string `json:"pk"`
		Category    string
		Name        string
		Description string
		IsComplete  bool
	}{
		ID:          "1",
		PK:          "pk1",
		Category:    "personal",
		Name:        "groceries",
		Description: "Pick up apples and strawberries",
		IsComplete:  false,
	}

	b, err := json.Marshal(item)
	if err != nil {
		return err
	}

	pk := azcosmos.NewPartitionKeyString(item.PK)
	createResponse, err := containerClient.UpsertItem(ctx, pk, b, nil)
	if err != nil {
		return err
	}
	log.Printf("UpsertItem | RequestCharge: %v", createResponse.RequestCharge)

	readResponse, err := containerClient.ReadItem(ctx, pk, item.ID, nil)
	if err != nil {
		return err
	}
	log.Printf("ReadItem | RequestCharge: %v", readResponse.RequestCharge)

	result := map[string]interface{}{}
	err = json.Unmarshal(readResponse.Value, &result)
	if err != nil {
		return err
	}

	b, err = json.MarshalIndent(result, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", b)

	// TODO: remove to enable deletion
	return nil

	deleteResponse, err := containerClient.DeleteItem(ctx, pk, item.ID, nil)
	if err != nil {
		return err
	}
	log.Printf("DeleteItem | RequestCharge: %v", deleteResponse.RequestCharge)

	return nil

}
