package csdata

import (
	"encoding/json"
	"fmt"
	"os"

	"database/sql"
	"io/ioutil"

	_ "github.com/lib/pq"
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

/*
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

*/

func (sheet *CharacterSheet) ReadFromFile(name string) error {
	file_path := "data/" + name + ".json"

	fmt.Printf("Checking if %s exists...\n", file_path)

	if _, err := os.Stat(file_path); err == nil {
		fmt.Printf("File %s; reading data from file...\n", file_path)

		jsonFile, err := os.Open(file_path)
		if err != nil {
			fmt.Printf("Error in opening file:\t%v\n", err)
			return err
		}
		defer jsonFile.Close()

		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Printf("Error in reading file:\t%v\n", err)
			return err
		}

		if err = json.Unmarshal(byteValue, &sheet); err != nil {
			fmt.Printf("Error in unmarshaling json:\t%v\n", err)
			return err
		}
	} else {
		fmt.Printf("File does not exist:\t%v\n", err)
		return err
	}

	return nil
}

// ReadFromDB performs an SQL query that returns:
// level,
// six stats - base,
// six stats - bonuses,
// saving throws proficiency
// for the given DB ID,
// and saves the data into the Character Sheet.
func (sheet *CharacterSheet) ReadFromDB(id int) error {
	file_content, err := ioutil.ReadFile("C:/d5cs/db_login")
	if err != nil {
		fmt.Printf("Error in opening login file:\t%v\n", err)
		return err
	}

	psqlconn := string(file_content)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Printf("Error in opening DB:\t%v\n", err)
		return err
	}
	defer db.Close()

	query := fmt.Sprintf(`
	SELECT "current_level",
	"str_base", "dex_base", "con_base",
	"int_base", "wis_base", "cha_base",
	"str_bonus", "dex_bonus", "con_bonus",
	"int_bonus", "wis_bonus", "cha_bonus",
	"st_str_prof", "st_dex_prof", "st_con_prof",
	"st_int_prof", "st_wis_prof", "st_cha_prof"
	FROM "sheets"
	WHERE "sheet_id" = %d;`, id)

	result := db.QueryRow(query)

	err = result.Scan(
		&sheet.Level,
		&sheet.StatsBase.Strength, &sheet.StatsBase.Dexterity, &sheet.StatsBase.Constitution,
		&sheet.StatsBase.Intelligence, &sheet.StatsBase.Wisdom, &sheet.StatsBase.Charisma,
		&sheet.StatsBonuses.Strength, &sheet.StatsBonuses.Dexterity, &sheet.StatsBonuses.Constitution,
		&sheet.StatsBonuses.Intelligence, &sheet.StatsBonuses.Wisdom, &sheet.StatsBonuses.Charisma,
		&sheet.STProficiency.Strength, &sheet.STProficiency.Dexterity, &sheet.STProficiency.Constitution,
		&sheet.STProficiency.Intelligence, &sheet.STProficiency.Wisdom, &sheet.STProficiency.Charisma,
	)
	if err != nil {
		fmt.Printf("Error in scanning:\t%v\n", err)
		return err
	}
	fmt.Printf("Query result:\n%v\n", sheet)
	sheet.Update()
	fmt.Printf("Sheet after Update():\n%v\n", sheet)

	return nil
}

func (sheet *CharacterSheet) WriteToDB(id int) error {
	file_content, err := ioutil.ReadFile("C:/d5cs/db_login")
	if err != nil {
		fmt.Printf("Error in opening login file:\t%v\n", err)
		return err
	}

	psqlconn := string(file_content)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Printf("Error in opening DB:\t%v\n", err)
		return err
	}
	defer db.Close()

	query := `
	UPDATE "sheets"
	SET "current_level" = $1,

	"str_base" = $2, "dex_base" = $3, "con_base" = $4,
	"int_base" = $5, "wis_base" = $6, "cha_base" = $7,

	"str_bonus" = $8, "dex_bonus" = $9, "con_bonus" = $10,
	"int_bonus" = $11, "wis_bonus" = $12, "cha_bonus" = $13,

	"st_str_prof" = $14, "st_dex_prof" = $15, "st_con_prof" = $16,
	"st_int_prof" = $17, "st_wis_prof" = $18, "st_cha_prof" = $19

	WHERE "sheet_id" = $20;
	`

	result, err := db.Exec(query,
		sheet.Level,
		sheet.StatsBase.Strength, sheet.StatsBase.Dexterity, sheet.StatsBase.Constitution,
		sheet.StatsBase.Intelligence, sheet.StatsBase.Wisdom, sheet.StatsBase.Charisma,

		sheet.StatsBonuses.Strength, sheet.StatsBonuses.Dexterity, sheet.StatsBonuses.Constitution,
		sheet.StatsBonuses.Intelligence, sheet.StatsBonuses.Wisdom, sheet.StatsBonuses.Charisma,

		sheet.STProficiency.Strength, sheet.STProficiency.Dexterity, sheet.STProficiency.Constitution,
		sheet.STProficiency.Intelligence, sheet.STProficiency.Wisdom, sheet.STProficiency.Charisma,

		id)
	if err != nil {
		fmt.Printf("Error in updating DB:\t%v\n", err)
		return err
	}

	n_rows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Error in RowsAffected():\t%v\n", err)
		return err
	}
	fmt.Printf("Number of affected rows:\t%v\n", n_rows)
	if n_rows == 0 {
		fmt.Printf("No rows affected; could not find row with id=%d\n", id)
	}

	return nil
}
