package csdata

type CharacterSheet struct {
	ProficiencyBonus int

	StatsBase      SixStats
	StatsBonuses   SixStats
	StatsSum       SixStats
	StatsModifiers SixStats

	STModifiers   SixStats
	STProficiency SixStatsCheck
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
