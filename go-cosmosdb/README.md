# go-cosmosdb

## Create Azure Cosmos DB Account

```bash
az group create \
    --name my-cosmosdb \
    --location eastus

az cosmosdb create \
    --resource-group my-cosmosdb \
    --name db220700 \
    --capabilities EnableServerless
```

## Run Sample

```bash
RESOURCE_GROUP='my-cosmosdb'
COSMOS_ACCOUNT_NAME='db220700'

export AZURE_COSMOS_ENDPOINT="https://${COSMOS_ACCOUNT_NAME}.documents.azure.com:443/"

export AZURE_COSMOS_KEY="$(az cosmosdb keys list \
    --resource-group $RESOURCE_GROUP \
    --name $COSMOS_ACCOUNT_NAME \
    --out tsv \
    --query primaryMasterKey)"

go run .
```
