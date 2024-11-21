package main

import (
	"fmt"
	"os"

	"github.com/Tiimie1/defactor-data-script/dummy"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	URL := os.Getenv("GRAPHQL_URL")
	ADMIN_SK := os.Getenv("HASURA_ADMIN_SECRET")
	USER_WALLET := os.Getenv("USER_WALLET")

	// Execute Mutation: Create Contacts
	for i := range dummy.DummyContacts {
		dummy.DummyContacts[i]["creator"] = USER_WALLET
	}
	for _, contact := range dummy.DummyContacts {
		createUserData, err := ExecuteGraphQL(URL, MutationCreateContact, contact, ADMIN_SK)
		if err != nil {
			fmt.Printf("Error creating contact %s: %v\n", contact["name"], err)
			continue
		}

		fmt.Printf("Successfully created contact: %s\nResponse: %s\n", contact["name"], string(createUserData))
	}

	// Execute Mutation: Create Templates
	for i := range dummy.DummyTemplates {
		dummy.DummyTemplates[i]["creator"] = USER_WALLET
	}
	for _, template := range dummy.DummyTemplates {
		createUserData, err := ExecuteGraphQL(URL, MutationCreateTemplate, template, ADMIN_SK)
		if err != nil {
			fmt.Printf("Error creating template %s: %v\n", template["template_name"], err)
			continue
		}

		fmt.Printf("Successfully created template: %s\nResponse: %s\n", template["template_name"], string(createUserData))
	}

	// Execute Mutation: Create Template Fields
	for _, field := range dummy.DummyTemplateFields {
		createUserData, err := ExecuteGraphQL(URL, MutationCreateTemplateField, field, ADMIN_SK)
		if err != nil {
			fmt.Printf("Error creating template field %s: %v\n", field["field_name"], err)
			continue
		}

		fmt.Printf("Successfully created template field: %s\nResponse: %s\n", field["field_name"], string(createUserData))
	}

	// Execute Mutation: Create Assets
	for i := range dummy.DummyAssets {
		dummy.DummyAssets[i]["creator"] = USER_WALLET
	}
	for _, asset := range dummy.DummyAssets {
		createAsset, err := ExecuteGraphQL(URL, MutationCreateAsset, asset, ADMIN_SK)
		if err != nil {
			fmt.Printf("Error creating asset %s: %v\n", asset["asset_name"], err)
			continue
		}
		fmt.Printf("Successfully created asset: %s\nResponse: %s\n", asset["asset_name"], string(createAsset))
	}

	// //Execute Mutation: Create Transactions
	for i := range dummy.DummyTransactions {
		dummy.DummyTransactions[i]["sender"] = USER_WALLET
		dummy.DummyTransactions[i]["receiver"] = USER_WALLET
	}
	for _, tx := range dummy.DummyTransactions {
		createTx, err := ExecuteGraphQL(URL, MutationCreateTransaction, tx, ADMIN_SK)
		if err != nil {
			fmt.Printf("Error creating transaction %s: %v\n", tx["type"], err)
			continue
		}

		fmt.Printf("Successfully created transaction: %s\nResponse: %s\n", tx["type"], string(createTx))
	}

	//Execute Mutation: Create MetadataEntries
	for i := range dummy.DummyMetadataEntries {
		dummy.DummyMetadataEntries[i]["creator"] = USER_WALLET
	}
	for _, metadataEntry := range dummy.DummyMetadataEntries {
		createMEntry, err := ExecuteGraphQL(URL, MutationCreateMetadataEntry, metadataEntry, ADMIN_SK)
		if err != nil {
			fmt.Printf("Error creating metadata entry %s: %v\n", metadataEntry["status"], err)
			continue
		}

		fmt.Printf("Successfully created metadata entry: %s\nResponse: %s\n", metadataEntry["status"], string(createMEntry))
	}

	// Execute Mutation: Create Audit Requests
	for _, auditRequest := range dummy.DummyAuditRequests {
		createAuditRequest, err := ExecuteGraphQL(URL, MutationCreateAuditRequest, auditRequest, ADMIN_SK)
		if err != nil {
			fmt.Printf("Error creating audit request %s: %v\n", auditRequest["email"], err)
			continue
		}

		fmt.Printf("Successfully created audit request : %s\nResponse: %s\n", auditRequest["email"], string(createAuditRequest))
	}
}
