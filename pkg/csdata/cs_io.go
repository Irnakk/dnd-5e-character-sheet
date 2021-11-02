package csdata

import (
	"encoding/json"
	"fmt"
	"os"
)

func (sheet *CharacterSheet) WriteFile(name string) error {
	marshaledResult, err := json.MarshalIndent(sheet, "", "	")
	if err != nil {
		fmt.Printf("Error in marshalling response_object:\t%v\n", err)
		return err
	}

	file, err := os.Create("data/" + name + ".json")
	if err != nil {
		fmt.Printf("Error in creating file:\t%v\n", err)
		return err
	}

	_, err = file.Write(marshaledResult)
	if err != nil {
		fmt.Printf("Error in writing to file:\t%v\n", err)
		return err
	}

	err = file.Close()
	if err != nil {
		fmt.Printf("Error in closing file:\t%v\n", err)
		return err
	}

	return nil
}
