package csdata

import (
	"math"
)

func (sheet *CharacterSheet) Update() {
	sheet.countProficiencyBonus()

	sheet.countStatsSum()

	sheet.countStatsModifiers()

	sheet.countSTModifiers()

	sheet.countSkillsModifiers()

	sheet.PassiveWisdom = 10 + sheet.SkillsModifiers.Perception
}

func (sheet *CharacterSheet) countProficiencyBonus() {
	if sheet.Level < 5 {
		sheet.ProficiencyBonus = 2
	} else if sheet.Level < 9 {
		sheet.ProficiencyBonus = 3
	} else if sheet.Level < 13 {
		sheet.ProficiencyBonus = 4
	} else if sheet.Level < 17 {
		sheet.ProficiencyBonus = 5
	} else {
		sheet.ProficiencyBonus = 6
	}
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
