package elo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type JsonFileTableStore struct {
	Filepath string
}

func (jfts *JsonFileTableStore) Load() (Table, error) {
	var table Table
	tableBytes, err := ioutil.ReadFile(jfts.Filepath)
	if err != nil {
		fmt.Println("----", err.Error())
		if strings.Contains(err.Error(), "no such file or directory") {
			return NewTable(32, 2000), nil
		} else {
			return Table{}, err
		}
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
