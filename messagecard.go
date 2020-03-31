package goteamsnotify

import (
	"errors"
	"strings"
)

// AddSection adds one or many additional MessageCardSection values to a
// MessageCard.
func (mc *MessageCard) AddSection(section ...MessageCardSection) {

	//logger.Printf("DEBUG: Existing sections: %+v\n", mc.Sections)
	//logger.Printf("DEBUG: Incoming sections: %+v\n", section)
	mc.Sections = append(mc.Sections, section...)
	//logger.Printf("Sections after append() call: %+v\n", mc.Sections)
}

// AddAction adds one or many additional MessageCardPotentialAction values to
// a MessageCard. It is also possible to add MessageCardPotentialAction values
// to specific sections as well.
func (mc *MessageCard) AddAction(action ...MessageCardPotentialAction) {

	//logger.Printf("DEBUG: Existing main card actions: %+v\n", mc.PotentialAction)
	//logger.Printf("DEBUG: Incoming main card actions: %+v\n", action)

	// FIXME: No more than four actions are currently supported according to the reference doc.
	mc.PotentialAction = append(mc.PotentialAction, action...)

	//logger.Printf("main card actions after append() call: %+v\n", mc.PotentialAction)
}

// AddFact adds one or many additional MessageCardSectionFact values to a
// MessageCardSection
func (mcs *MessageCardSection) AddFact(fact ...MessageCardSectionFact) {

	//logger.Printf("DEBUG: Existing sections: %+v\n", mcs.Facts)
	//logger.Printf("DEBUG: Incoming sections: %+v\n", fact)
	mcs.Facts = append(mcs.Facts, fact...)
	//logger.Printf("Facts after append() call: %+v\n", mcs.Facts)

}

// AddFactFromKeyValue accepts a key and slice of values and converts them to
// MessageCardSectionFact values
func (mcs *MessageCardSection) AddFactFromKeyValue(key string, values ...string) error {

	// validate arguments

	if key == "" {
		return errors.New("empty key received for new fact")
	}

	if len(values) < 1 {
		return errors.New("no values received for new fact")
	}

	fact := MessageCardSectionFact{
		Name:  key,
		Value: strings.Join(values, ", "),
	}

	mcs.Facts = append(mcs.Facts, fact)

	// if we made it this far then all should be well
	return nil
}

// AddAction adds one or many additional MessageCardPotentialAction values to
// a MessageCard section.
func (mcs *MessageCardSection) AddAction(sectionAction ...MessageCardPotentialAction) {

	//logger.Printf("DEBUG: Existing section actions: %+v\n", mcs.PotentialAction)
	//logger.Printf("DEBUG: Incoming section actions: %+v\n", sectionAction)

	// FIXME: No more than four actions are currently supported according to the reference doc.
	mcs.PotentialAction = append(mcs.PotentialAction, sectionAction...)

	//logger.Printf("Section actions after append() call: %+v\n", mcs.PotentialAction)
}
