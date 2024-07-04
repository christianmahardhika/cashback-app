package migrations

import "gofr.dev/pkg/gofr/migration"

const createTenantsTable = `
CREATE TABLE IF NOT EXISTS tenants (
	id UUID PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	api_key VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);
`

func createTableTenants() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createTenantsTable)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
