package albionAPI

type Localized struct {
	EN string `json:"EN-US"`
	PL string `json:"PL-PL"`
	JP string `json:"JA-JP"`
}

func (l Localized) getEn() string {
	return l.EN
}

func (l Localized) getPl() string {
	return l.PL
}

func (l Localized) getJp() string {
	return l.JP
}

type Item struct {
	ID          string    `json:"Index"`
	NAME        Localized `json:"LocalizedNames"`
	DESCRIPTION Localized `json:"LocalizedDescriptions"`
	UNIQUE_NAME string    `json:"UniqueName"`

	TIER        int
	ENCHANTMENT int
}

func (i Item) getEnName() string {
	return i.NAME.getEn()
}

func (i Item) getEnDescription() string {
	return i.DESCRIPTION.getEn()
}

func (i Item) getPlName() string {
	return i.NAME.getPl()
}

func (i Item) getPlDescription() string {
	return i.DESCRIPTION.getPl()
}

func (i Item) getJpName() string {
	return i.NAME.getJp()
}

func (i Item) getJpDescription() string {
	return i.DESCRIPTION.getJp()
}
