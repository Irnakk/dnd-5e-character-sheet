package csdata

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type CharacterSheet struct {
	ProficiencyBonus int

	StatsBase      SixStats
	StatsBonuses   SixStats
	StatsSum       SixStats
	StatsModifiers SixStats

	STModifiers   SixStats
	STProficiency SixStatsCheck

	SkillsModifiers   Skills
	SkillsProficiency SkillsCheck

	PassiveWisdom int
}

type SixStats struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
}

type SixStatsCheck struct {
	Strength     bool
	Dexterity    bool
	Constitution bool
	Intelligence bool
	Wisdom       bool
	Charisma     bool
}

type Skills struct {
	Acrobatics     int
	AnimalHandling int
	Arcana         int
	Athletics      int
	Deception      int
	History        int
	Insight        int
	Intimidation   int
	Investigation  int
	Medicine       int
	Nature         int
	Perception     int
	Performance    int
	Persuasion     int
	Religion       int
	SleightOfHand  int
	Stealth        int
	Survival       int
}

type SkillsCheck struct {
	Acrobatics     bool
	AnimalHandling bool
	Arcana         bool
	Athletics      bool
	Deception      bool
	History        bool
	Insight        bool
	Intimidation   bool
	Investigation  bool
	Medicine       bool
	Nature         bool
	Perception     bool
	Performance    bool
	Persuasion     bool
	Religion       bool
	SleightOfHand  bool
	Stealth        bool
	Survival       bool
}

func (sheet *CharacterSheet) Update() {
	// Add ProficiencyBonus count from level

	sheet.countStatsSum()

	sheet.countStatsModifiers()

	sheet.countSTModifiers()

	sheet.countSkillsModifiers()

	sheet.PassiveWisdom = 10 + sheet.SkillsModifiers.Perception
}

func (sheet *CharacterSheet) countStatsSum() {
	// Wouldn't it be better to pass the resulting SixStats{} in return?
	sheet.StatsSum = SixStats{
		Strength:     sheet.StatsBase.Strength + sheet.StatsBonuses.Strength,
		Dexterity:    sheet.StatsBase.Dexterity + sheet.StatsBonuses.Dexterity,
		Constitution: sheet.StatsBase.Constitution + sheet.StatsBonuses.Constitution,
		Intelligence: sheet.StatsBase.Intelligence + sheet.StatsBonuses.Intelligence,
		Wisdom:       sheet.StatsBase.Wisdom + sheet.StatsBonuses.Wisdom,
		Charisma:     sheet.StatsBase.Charisma + sheet.StatsBonuses.Charisma,
	}
}

func (sheet *CharacterSheet) modifier(value int) int {
	return int(math.Floor((float64(value) - 10.) / 2.))
}

func (sheet *CharacterSheet) countStatsModifiers() {
	sheet.StatsModifiers = SixStats{
		Strength:     sheet.modifier(sheet.StatsSum.Strength),
		Dexterity:    sheet.modifier(sheet.StatsSum.Dexterity),
		Constitution: sheet.modifier(sheet.StatsSum.Constitution),
		Intelligence: sheet.modifier(sheet.StatsSum.Intelligence),
		Wisdom:       sheet.modifier(sheet.StatsSum.Wisdom),
		Charisma:     sheet.modifier(sheet.StatsSum.Charisma),
	}
}

func (sheet *CharacterSheet) countSTModifiers() {
	sheet.STModifiers.Strength = sheet.StatsModifiers.Strength
	if sheet.STProficiency.Strength {
		sheet.STModifiers.Strength += sheet.ProficiencyBonus
	}

	sheet.STModifiers.Dexterity = sheet.StatsModifiers.Dexterity
	if sheet.STProficiency.Dexterity {
		sheet.STModifiers.Dexterity += sheet.ProficiencyBonus
	}

	sheet.STModifiers.Constitution = sheet.StatsModifiers.Constitution
	if sheet.STProficiency.Constitution {
		sheet.STModifiers.Constitution += sheet.ProficiencyBonus
	}

	sheet.STModifiers.Intelligence = sheet.StatsModifiers.Intelligence
	if sheet.STProficiency.Intelligence {
		sheet.STModifiers.Intelligence += sheet.ProficiencyBonus
	}

	sheet.STModifiers.Wisdom = sheet.StatsModifiers.Wisdom
	if sheet.STProficiency.Wisdom {
		sheet.STModifiers.Wisdom += sheet.ProficiencyBonus
	}

	sheet.STModifiers.Charisma = sheet.StatsModifiers.Charisma
	if sheet.STProficiency.Charisma {
		sheet.STModifiers.Charisma += sheet.ProficiencyBonus
	}
}

func (sheet *CharacterSheet) countSkillsModifiers() {
	sheet.SkillsModifiers.Acrobatics = sheet.StatsModifiers.Dexterity
	if sheet.SkillsProficiency.Acrobatics {
		sheet.SkillsModifiers.Acrobatics += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.AnimalHandling = sheet.StatsModifiers.Wisdom
	if sheet.SkillsProficiency.AnimalHandling {
		sheet.SkillsModifiers.AnimalHandling += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.Arcana = sheet.StatsModifiers.Intelligence
	if sheet.SkillsProficiency.Arcana {
		sheet.SkillsModifiers.Arcana += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.Athletics = sheet.StatsModifiers.Strength
	if sheet.SkillsProficiency.Athletics {
		sheet.SkillsModifiers.Athletics += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.Deception = sheet.StatsModifiers.Charisma
	if sheet.SkillsProficiency.Deception {
		sheet.SkillsModifiers.Deception += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.History = sheet.StatsModifiers.Intelligence
	if sheet.SkillsProficiency.History {
		sheet.SkillsModifiers.History += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.Insight = sheet.StatsModifiers.Wisdom
	if sheet.SkillsProficiency.Insight {
		sheet.SkillsModifiers.Insight += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.Intimidation = sheet.StatsModifiers.Charisma
	if sheet.SkillsProficiency.Intimidation {
		sheet.SkillsModifiers.Intimidation += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.Investigation = sheet.StatsModifiers.Intelligence
	if sheet.SkillsProficiency.Investigation {
		sheet.SkillsModifiers.Investigation += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.Medicine = sheet.StatsModifiers.Wisdom
	if sheet.SkillsProficiency.Medicine {
		sheet.SkillsModifiers.Medicine += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.Nature = sheet.StatsModifiers.Intelligence
	if sheet.SkillsProficiency.Nature {
		sheet.SkillsModifiers.Nature += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.Perception = sheet.StatsModifiers.Wisdom
	if sheet.SkillsProficiency.Perception {
		sheet.SkillsModifiers.Perception += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.Performance = sheet.StatsModifiers.Charisma
	if sheet.SkillsProficiency.Performance {
		sheet.SkillsModifiers.Performance += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.Persuasion = sheet.StatsModifiers.Charisma
	if sheet.SkillsProficiency.Persuasion {
		sheet.SkillsModifiers.Persuasion += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.Religion = sheet.StatsModifiers.Intelligence
	if sheet.SkillsProficiency.Religion {
		sheet.SkillsModifiers.Religion += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.SleightOfHand = sheet.StatsModifiers.Dexterity
	if sheet.SkillsProficiency.SleightOfHand {
		sheet.SkillsModifiers.SleightOfHand += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.Stealth = sheet.StatsModifiers.Dexterity
	if sheet.SkillsProficiency.Stealth {
		sheet.SkillsModifiers.Stealth += sheet.ProficiencyBonus
	}

	sheet.SkillsModifiers.Survival = sheet.StatsModifiers.Wisdom
	if sheet.SkillsProficiency.Survival {
		sheet.SkillsModifiers.Survival += sheet.ProficiencyBonus
	}
}

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
