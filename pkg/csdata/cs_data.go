/*
Package csdata (stands for Character Sheet Data) provides:

a. the data structures involved in the creation and use of a
character sheet

b. functions that manipulate the data stored in a character sheet
and help to fill in the data.
*/
package csdata

// CharacterSheet includes the character information
// including the derived fields such as modifiers.
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

// SixStats structure is used in both
// describing the value of a base character stat
// (such as Intelligence) and the corresponding
// stat modifiers (both for stats modifiers and
// saving throws modifiers).
type SixStats struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
}

// SixStatsCheck is used to describe proficiency
// in saving throws.
type SixStatsCheck struct {
	Strength     bool
	Dexterity    bool
	Constitution bool
	Intelligence bool
	Wisdom       bool
	Charisma     bool
}

// Skills is used to store the modifiers
// for each skill.
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

// Skills check is used to describe proficiency
// in skills.
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
