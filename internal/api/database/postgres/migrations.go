package postgres

type table struct {
	Name    string
	Columns string
	Indexes []string
}

func createTable(table table) {
	schema := `
		CREATE TABLE IF NOT EXISTS ` + table.Name + ` (
			` + table.Columns + `
		)
	`
	Instance.MustExec(schema)

	for _, index := range table.Indexes {
		Instance.MustExec(index)
	}
}

func runMigrations() {
	tables := []table{
		{
			Name: "proxy",
			Columns: `
				id UUID PRIMARY KEY,
				address VARCHAR(64) UNIQUE,
				username VARCHAR(64) NOT NULL,
				password VARCHAR(64) NOT NULL,
				scheme VARCHAR(64) NOT NULL
			`,
		},
		{
			Name: "api_key",
			Columns: `
				id UUID PRIMARY KEY,
				key VARCHAR(64) UNIQUE,
				ip_address VARCHAR(64) NOT NULL,
				num_proxies INTEGER NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
				expiration_date TIMESTAMP NOT NULL
			`,
		},
		{
			Name: "client",
			Columns: `
				id UUID PRIMARY KEY,
				name VARCHAR(64),
				api_key_id UUID,
				FOREIGN KEY (api_key_id) REFERENCES api_key(id)
			`,
			Indexes: []string{
				"CREATE INDEX IF NOT EXISTS client_api_key_id_idx ON client(api_key_id)",
			},
		},
		{
			Name: "api_key_proxy",
			Columns: `
				api_key_id UUID,
				proxy_id UUID,
				PRIMARY KEY(api_key_id, proxy_id),
				FOREIGN KEY (api_key_id) REFERENCES api_key(id),
				FOREIGN KEY (proxy_id) REFERENCES proxy(id)
			`,
			Indexes: []string{
				"CREATE INDEX IF NOT EXISTS api_key_proxy_api_key_id_idx ON api_key_proxy(api_key_id)",
				"CREATE INDEX IF NOT EXISTS api_key_proxy_proxy_id_idx ON api_key_proxy(proxy_id)",
			},
		},
	}

	for _, table := range tables {
		createTable(table)
	}
}
