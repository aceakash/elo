package main

import (
	"github.com/aceakash/elo"
	"encoding/json"
	"fmt"
	"os"
)
type inMemoryTableStore struct {
	storage string
}

func (jfts *inMemoryTableStore) Load() (elo.Table, error) {
	var table elo.Table
	if err := json.Unmarshal([]byte(jfts.storage), &table); err != nil {
		fmt.Fprintln(os.Stderr, "Error decoding JSON")
		return elo.Table{}, err
	}
	return table, nil
}

func (jfts *inMemoryTableStore) Save(table elo.Table) error {
	b, err := json.Marshal(table)
	if err != nil {
		return err
	}
	jfts.storage = string(b)
	return nil
}



