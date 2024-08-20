package guard

import "embed"

//go:embed schema/*.json
var schemaMigrations embed.FS

func loadModel() ([]byte, error) {
	return schemaMigrations.ReadFile("schema/product_catalogue.json")
}
