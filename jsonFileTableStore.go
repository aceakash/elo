package elo

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

type JsonFileTableStore struct {
	Filepath string
}

func (jfts *JsonFileTableStore) Load() (Table, error) {
	var table Table
	tableBytes, err := ioutil.ReadFile(jfts.Filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file")
		return Table{}, err
	}
	if err = json.Unmarshal(tableBytes, &table); err != nil {
		fmt.Fprintln(os.Stderr, "Error decoding JSON")
		return Table{}, err
	}
	return table, nil
}

func (jfts *JsonFileTableStore) Save(table Table) error {
	b, err := json.Marshal(table)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(jfts.Filepath, b, 0600)
}



