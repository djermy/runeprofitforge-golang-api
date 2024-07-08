package psqlstore

type Migration struct {
	Up   string
	Down string
}

var migrations = []Migration{
	{
		Up: `CREATE TABLE IF NOT EXISTS item (
			id SERIAL PRIMARY KEY,
			name VARCHAR NOT NULL,
			type VARCHAR NOT NULL,
			members VARCHAR NOT NULL,
			description VARCHAR NOT NULL,
			icon VARCHAR NOT NULL,
			icon_large VARCHAR NOT NULL,
			type_icon VARCHAR NOT NULL,
			current JSON,
			today JSON
		);`,
		Down: `DROP TABLE IF EXISTS item;`,
	},
}
