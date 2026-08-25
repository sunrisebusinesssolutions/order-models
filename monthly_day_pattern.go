package models

type MonthlyDayPattern string

const (
	MonthlyDayPatternFirstOccurrence     MonthlyDayPattern = "FIRST_OCCURRENCE"
	MonthlyDayPatternSecondOccurrence    MonthlyDayPattern = "SECOND_OCCURRENCE"
	MonthlyDayPatternThirdOccurrence     MonthlyDayPattern = "THIRD_OCCURRENCE"
	MonthlyDayPatternFourthOccurrence    MonthlyDayPattern = "FOURTH_OCCURRENCE"
	MonthlyDayPatternLastOccurrence      MonthlyDayPattern = "LAST_OCCURRENCE"
	MonthlyDayPatternLastBusinessDay     MonthlyDayPattern = "LAST_BUSINESS_DAY"
)

func (mdp MonthlyDayPattern) GetOccurrenceNumber() int {
	switch mdp {
	case MonthlyDayPatternFirstOccurrence:
		return 1
	case MonthlyDayPatternSecondOccurrence:
		return 2
	case MonthlyDayPatternThirdOccurrence:
		return 3
	case MonthlyDayPatternFourthOccurrence:
		return 4
	case MonthlyDayPatternLastOccurrence:
		return 5
	case MonthlyDayPatternLastBusinessDay:
		return 6
	default:
		return 0
	}
}

func (mdp MonthlyDayPattern) GetDisplayName() string {
	switch mdp {
	case MonthlyDayPatternFirstOccurrence:
		return "First"
	case MonthlyDayPatternSecondOccurrence:
		return "Second"
	case MonthlyDayPatternThirdOccurrence:
		return "Third"
	case MonthlyDayPatternFourthOccurrence:
		return "Fourth"
	case MonthlyDayPatternLastOccurrence:
		return "Last"
	case MonthlyDayPatternLastBusinessDay:
		return "Last Business Day"
	default:
		return ""
	}
}
