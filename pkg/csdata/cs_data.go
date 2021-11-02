package csdata

type CharacterSheet struct {
	Level            int
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
