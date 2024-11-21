package main

const (
	QueryGetUsers = `
	query GetAssets {
		global_asset {
			id
		}
	}`

	MutationCreateContact = `
	mutation CreateContact($name: String!, $wallet: String!, $creator: String!) {
		insert_global_contact(objects: {name: $name, wallet: $wallet, creator: $creator}) {
			returning {
				uuid
			}
		}
	}`

	MutationCreateTemplate = `
	mutation CreateTemplate($id: uuid, $template_name: String!, $description: String!, $creator: String!, $isDerived: Boolean!, $isCommunity: Boolean!) {
		insert_global_metadata_template(objects: {id: $id, template_name: $template_name, description: $description, creator: $creator, isDerived: $isDerived, isCommunity: $isCommunity}) {
			returning {
				id
			}
		}
	}`

	MutationCreateTemplateField = `
	mutation CreateTemplateField($id: uuid, $field_name: String!, $is_dynamic: Boolean!, $field_type: global_field_type_enum, $field_length: global_field_size_enum, $is_required: Boolean!, $is_tooltip: Boolean!, $metadata_id: uuid, $placeholder: String, $position_in_metadata: Int, $suffix: global_suffix_enum, $tooltip_text: String) {
		insert_global_metadata_template_field(objects: {id: $id, field_length: $field_length, field_name: $field_name, field_type: $field_type, metadata_id: $metadata_id, is_dynamic: $is_dynamic, is_required: $is_required, is_tooltip: $is_tooltip, placeholder:$placeholder, position_in_metadata: $position_in_metadata, suffix: $suffix, tooltip_text:$tooltip_text}) {
			returning {
				id
			}
		}
	}`

	MutationCreateAsset = `
	mutation CreateAsset($id: uuid, $asset_type: global_asset_type_enum, $asset_icon: String, $supply: String, $status: String, $spf: String, $price: String, $description: String, $chain_id: Int, $asset_name: String, $asset_category: global_asset_category_enum, $creator: String, $asset_symbol: String, $contract_address: String) {
		insert_global_asset(objects: {id: $id, asset_type: $asset_type, asset_icon: $asset_icon, supply: $supply, status: $status, spf: $spf, price: $price, description: $description, chain_id: $chain_id, asset_name: $asset_name, asset_category: $asset_category, creator: $creator, asset_symbol: $asset_symbol, contract_address: $contract_address}) {
			returning {
				id
			}
		}
	}`

	MutationCreateTransaction = `
	mutation CreateTransaction($chain: Int, $sender: String, $state: String, $type: String, $asset_id: uuid, $value: String, $receiver: String, $created_at: timestamptz, $mined_at: timestamptz, $hash: String, $signed_data: String) {
		insert_global_transaction(objects: {chain: $chain, sender: $sender, state: $state, type: $type, asset_id: $asset_id, value: $value, receiver: $receiver, created_at: $created_at, mined_at: $mined_at, hash: $hash, signed_data: $signed_data }) {
			returning {
				id
			}
		}
	}`

	MutationCreateMetadataEntry = `
	mutation CreateMetadataEntry($id: uuid $asset_id: uuid, $status: global_metadata_status_enum, $metadata_id: uuid, $creator: String, $created_at: timestamptz, $updated_at: timestamptz) {
		insert_global_metadata_entry(objects: {id: $id, asset_id: $asset_id, status: $status, metadata_id: $metadata_id, creator:$creator, created_at: $created_at, updated_at: $updated_at }) {
			returning {
				id
			}
		}
	}`
	MutationCreateAuditRequest = `
	mutation CreateAuditRequest($email: String, $metadata_entry_id: uuid) {
		insert_global_audit_request(objects: {email: $email, metadata_entry_id: $metadata_entry_id}) {
			returning {
				id
			}
		}
	}`
)
