package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"encoding/csv"
)

func writeStringDataMap(path string, dataMap *map[string]string) error {
	if len(*dataMap) == 0 {
		// do nothing
		return nil
	}
	bytes, _ := json.Marshal(*dataMap)
	return ioutil.WriteFile(path, bytes, os.FileMode(0600))
}

func readStringDataMap(path string) (dataMap map[string]string, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		// do nothing
		return nil, err
	}
	err = json.Unmarshal(file, &dataMap)
	return
}

func overwriteStringDataMapCSV(path string, dataMap *map[string]string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	writer := csv.NewWriter(file)
	for key, value := range *dataMap {
		err = writer.Write([]string{key, value})
		//		fmt.Printf("key:%s -> value=%d\n", key, value)
	}
	writer.Flush()
	return
}

func addStringDataMapCSV(path string, externalID string, internalID string) (err error) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write([]string{externalID, internalID})
	writer.Flush()
	return
}

func readStringDataMapCSV(path string, dataMap *map[string]string) (err error) {
	file, err := os.Open(path)
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	allRecords, err := reader.ReadAll()

	for _, rec := range allRecords {
		//		fmt.Printf("%s:%s\n", rec[0], rec[1])
		(*dataMap)[rec[0]] = rec[1]
	}
	return
}
