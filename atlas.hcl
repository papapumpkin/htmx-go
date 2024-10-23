data "external_schema" "gorm" {
    program = [
        "go",
        "run",
        "-mod=mod",
        "./cmd/atlas"
    ]
}

env "dev" {
    src = data.external_schema.gorm.url
    dev = "postgres://user:password@localhost:5432/database?sslmode=disable"
    migration = {
        dir = "file://internals/db/migrations?format=goose"
    }
    format {
        migrate {
            diff = "{{ sql . \" \" }}"
        }
    }
}
