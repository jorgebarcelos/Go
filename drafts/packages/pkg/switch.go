package assistant

func DayWeek(number int) string{
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Tursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid!!"
	}
}