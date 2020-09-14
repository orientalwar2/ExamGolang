package main

import "path"

type AppConfig struct {
	dataPath string
	db       []*Barang
}

func newConfig() *AppConfig {
	const DATA_PATH = "data"
	const DATA_FILE_NAME = "data_barang.json"

	var dataBarang = make([]*Barang, 1)
	return &AppConfig{
		dataPath: path.Join(DATA_PATH, DATA_FILE_NAME),
		db:       dataBarang,
	}

}

func (c *AppConfig) getDb() []*Barang {
	return c.db
}

func (c *AppConfig) getDbPath() string {
	return c.dataPath
}
