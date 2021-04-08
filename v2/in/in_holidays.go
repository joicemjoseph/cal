package in

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
)

var (
	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "New Year's Day", Type: cal.ObservanceUnknown})

	// RepublicDay of India.
	RepublicDay = &cal.Holiday{
		Name:  "Republic Day",
		Type:  cal.ObservancePublic,
		Month: time.January,
		Day:   26,
		Func:  cal.CalcDayOfMonth,
	}

	// MahaShivRathri .
	MahaShivRathri = &cal.Holiday{
		Name: "Mahashivratri",
		Type: cal.ObservancePublic,
	}

	// Holi .
	Holi = &cal.Holiday{
		Name: "Holi",
		Type: cal.ObservancePublic,
	}

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = aa.GoodFriday.Clone(&cal.Holiday{Name: "Good Friday", Type: cal.ObservancePublic})

	// BabaSahebJayanti is Dr.Baba Saheb Ambedkar's birth day.
	BabaSahebJayanti = &cal.Holiday{
		Name:  "Dr. Baba Saheb Ambedkar Jayanti",
		Type:  cal.ObservancePublic,
		Month: time.April,
		Day:   14,
		Func:  cal.CalcDayOfMonth,
	}

	// RamaNavami .
	RamaNavami = &cal.Holiday{
		Name: "Ram Navami",
		Type: cal.ObservancePublic,
	}

	// WorkersDay represents Labor Day on the first Monday in May
	WorkersDay = aa.WorkersDay.Clone(&cal.Holiday{Type: cal.ObservancePublic})

	// Ramzan or Id-Ul-Fitr.
	Ramzan = &cal.Holiday{
		Name: "Id-Ul-Fitr (Ramzan ID)",
		Type: cal.ObservancePublic,
	}

	// BakrID is Bakri ID.
	BakrID = &cal.Holiday{
		Name: "Bakri ID",
		Type: cal.ObservancePublic,
	}

	// Moharram .
	Moharram = &cal.Holiday{
		Name: "Moharram",
		Type: cal.ObservancePublic,
	}

	// IndependanceDay .
	IndependanceDay = &cal.Holiday{
		Name:  "Independance Day",
		Type:  cal.ObservancePublic,
		Month: time.August,
		Day:   15,
		Func:  cal.CalcDayOfMonth,
	}
	// GaneshChaturthi .
	GaneshChaturthi = &cal.Holiday{
		Name: "Ganesh Chaturthi",
		Type: cal.ObservancePublic,
	}

	// Dussehra  or Vijayadashami.
	// On 10th of Ashvin as per Hindu calendar.
	Dussehra = &cal.Holiday{
		Name: "Dussehra",
		Type: cal.ObservancePublic,
	}

	// DiwaliLaxmiPujan .
	DiwaliLaxmiPujan = &cal.Holiday{
		Name: "Diwali - Laxmi Pujan",
		Type: cal.ObservanceOther,
	}

	// DiwaliBalipratipada .
	DiwaliBalipratipada = &cal.Holiday{
		Name: "Diwali Balipratipada",
		Type: cal.ObservancePublic,
	}

	// GandhiJayanti .
	GandhiJayanti = &cal.Holiday{
		Name:  "Gandhi Jayanti",
		Type:  cal.ObservancePublic,
		Month: time.October,
		Day:   2,
		Func:  cal.CalcDayOfMonth,
	}

	// GurunanakJayanti .
	GurunanakJayanti = &cal.Holiday{
		Name:  "Gurunanak Jayanti",
		Type:  cal.ObservancePublic,
		Month: time.November,
		Day:   19,
		Func:  cal.CalcDayOfMonth,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Type: cal.ObservancePublic})
)
