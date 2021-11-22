package csdata

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"database/sql"
	"io/ioutil"

	_ "github.com/lib/pq"
)

// WriteFile marshals the sheet and saves the data into
// a JSON file with a given name into the data/ folder.
func (sheet *CharacterSheet) WriteFile(name string) error {
	marshaledResult, err := json.MarshalIndent(sheet, "", "	")
	if err != nil {
		log.Printf("Error in marshalling response_object:\t%v\n", err)
		return err
	}

	file, err := os.Create("data/" + name + ".json")
	if err != nil {
		log.Printf("Error in creating file:\t%v\n", err)
		return err
	}

	_, err = file.Write(marshaledResult)
	if err != nil {
		log.Printf("Error in writing to file:\t%v\n", err)
		return err
	}

	err = file.Close()
	if err != nil {
		log.Printf("Error in closing file:\t%v\n", err)
		return err
	}

	return nil
}

// ReadFromFile checks if a JSON file with a given name exists
// at data/ folder, reads the sheet object from the file
// and saves the data into the current sheet.
func (sheet *CharacterSheet) ReadFromFile(name string) error {
	file_path := "data/" + name + ".json"

	log.Printf("Checking if %s exists...\n", file_path)

	if _, err := os.Stat(file_path); err == nil {
		log.Printf("File %s; reading data from file...\n", file_path)

		jsonFile, err := os.Open(file_path)
		if err != nil {
			log.Printf("Error in opening file:\t%v\n", err)
			return err
		}
		defer jsonFile.Close()

		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			log.Printf("Error in reading file:\t%v\n", err)
			return err
		}

		if err = json.Unmarshal(byteValue, &sheet); err != nil {
			log.Printf("Error in unmarshaling json:\t%v\n", err)
			return err
		}
	} else {
		log.Printf("File does not exist:\t%v\n", err)
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
		log.Printf("Error in opening login file:\t%v\n", err)
		return err
	}

	psqlconn := string(file_content)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Printf("Error in opening DB:\t%v\n", err)
		return err
	}
	defer db.Close()

	query := fmt.Sprintf(`
	SELECT
	"current_level",
	"str_base", "dex_base", "con_base",
	"int_base", "wis_base", "cha_base",

	"str_bonus", "dex_bonus", "con_bonus",
	"int_bonus", "wis_bonus", "cha_bonus",

	"st_str_prof", "st_dex_prof", "st_con_prof",
	"st_int_prof", "st_wis_prof", "st_cha_prof",

	"acrobatics",      "animal_handling", "arcana",
	"athletics",       "deception",       "history",
    "insight",         "intimidation",    "investigation",
    "medicine",        "nature",          "perception",
    "performance",     "persuasion",      "religion",
    "sleight_of_hand", "stealth",         "survival"

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

		&sheet.SkillsProficiency.Acrobatics, &sheet.SkillsProficiency.AnimalHandling, &sheet.SkillsProficiency.Arcana,
		&sheet.SkillsProficiency.Athletics, &sheet.SkillsProficiency.Deception, &sheet.SkillsProficiency.History,
		&sheet.SkillsProficiency.Insight, &sheet.SkillsProficiency.Intimidation, &sheet.SkillsProficiency.Investigation,
		&sheet.SkillsProficiency.Medicine, &sheet.SkillsProficiency.Nature, &sheet.SkillsProficiency.Perception,
		&sheet.SkillsProficiency.Performance, &sheet.SkillsProficiency.Persuasion, &sheet.SkillsProficiency.Religion,
		&sheet.SkillsProficiency.SleightOfHand, &sheet.SkillsProficiency.Stealth, &sheet.SkillsProficiency.Survival,
	)
	if err != nil {
		log.Printf("Error in scanning:\t%v\n", err)
		return err
	}

	log.Printf("Query result:\n%v\n", sheet)

	sheet.Update()
	log.Printf("Sheet after Update():\n%v\n", sheet)

	return nil
}

// WriteToDB performs an SQL query and
// updates the row with a given id.
// The following fields are copied from the sheet:
// level,
// six stats - base,
// six stats - bonuses,
// saving throws proficiency.
// Note: if the query does not find a row with
// a given id, a new row will not be added.
func (sheet *CharacterSheet) WriteToDB(id int) error {
	file_content, err := ioutil.ReadFile("C:/d5cs/db_login")
	if err != nil {
		log.Printf("Error in opening login file:\t%v\n", err)
		return err
	}

	psqlconn := string(file_content)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Printf("Error in opening DB:\t%v\n", err)
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
	"st_int_prof" = $17, "st_wis_prof" = $18, "st_cha_prof" = $19,

	"acrobatics" = $20,      "animal_handling" = $21, "arcana" = $22,
	"athletics" = $23,       "deception" = $24,       "history" = $25,
    "insight" = $26,         "intimidation" = $27,    "investigation" = $28,
    "medicine" = $29,        "nature" = $30,          "perception" = $31,
    "performance" = $32,     "persuasion" = $33,      "religion" = $34,
    "sleight_of_hand" = $35, "stealth" = $36,         "survival" = $37

	WHERE "sheet_id" = $38;
	`

	// looks like it is a better way to perform a query with arguments
	result, err := db.Exec(query,
		sheet.Level,
		sheet.StatsBase.Strength, sheet.StatsBase.Dexterity, sheet.StatsBase.Constitution,
		sheet.StatsBase.Intelligence, sheet.StatsBase.Wisdom, sheet.StatsBase.Charisma,

		sheet.StatsBonuses.Strength, sheet.StatsBonuses.Dexterity, sheet.StatsBonuses.Constitution,
		sheet.StatsBonuses.Intelligence, sheet.StatsBonuses.Wisdom, sheet.StatsBonuses.Charisma,

		sheet.STProficiency.Strength, sheet.STProficiency.Dexterity, sheet.STProficiency.Constitution,
		sheet.STProficiency.Intelligence, sheet.STProficiency.Wisdom, sheet.STProficiency.Charisma,

		sheet.SkillsProficiency.Acrobatics, sheet.SkillsProficiency.AnimalHandling, sheet.SkillsProficiency.Arcana,
		sheet.SkillsProficiency.Athletics, sheet.SkillsProficiency.Deception, sheet.SkillsProficiency.History,
		sheet.SkillsProficiency.Insight, sheet.SkillsProficiency.Intimidation, sheet.SkillsProficiency.Investigation,
		sheet.SkillsProficiency.Medicine, sheet.SkillsProficiency.Nature, sheet.SkillsProficiency.Perception,
		sheet.SkillsProficiency.Performance, sheet.SkillsProficiency.Persuasion, sheet.SkillsProficiency.Religion,
		sheet.SkillsProficiency.SleightOfHand, sheet.SkillsProficiency.Stealth, sheet.SkillsProficiency.Survival,

		id)
	if err != nil {
		log.Printf("Error in updating DB:\t%v\n", err)
		return err
	}

	n_rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error in RowsAffected():\t%v\n", err)
		return err
	}
	log.Printf("Number of affected rows:\t%v\n", n_rows)
	if n_rows == 0 {
		log.Printf("No rows affected; could not find row with id=%d\n", id)
	}

	return nil
}
