package albiondatabase

type City struct {
	ID   string `json:"Index"`
	NAME string `json:"UniqueName"`
}

type Localized struct {
	EN string `json:"EN-US"`
	PL string `json:"PL-PL"`
	JP string `json:"JA-JP"`
}

func getEn(el Localized) string {
	return el.EN
}

func getPl(el Localized) string {
	return el.PL
}

func getJp(el Localized) string {
	return el.JP
}

type Item struct {
	ID          string    `json:"Index"`
	NAME        Localized `json:"LocalizedNames"`
	DESCRIPTION Localized `json:"LocalizedDescriptions"`
	UNIQUE_NAME string    `json:"UniqueName"`
}

func getEnName(el Item) string {
	return getEn(el.NAME)
}

func getEnDescription(el Item) string {
	return getEn(el.DESCRIPTION)
}

func getPlName(el Item) string {
	return getPl(el.NAME)
}

func getPlDescription(el Item) string {
	return getPl(el.DESCRIPTION)
}

func getJpName(el Item) string {
	return getJp(el.NAME)
}

func getJpDescription(el Item) string {
	return getJp(el.DESCRIPTION)
}

type ItemReadable struct {
	ID          string
	NAME        string
	TIER        int
	ENCHANTMENT int
}

func ItemToReadable(item Item) ItemReadable {
	var item_readable ItemReadable

	item_readable.ID = item.ID
	item_readable.NAME = getEnName(item)
	item_readable.TIER = int(item.UNIQUE_NAME[1] - '0')

	var enchantment int
	if item.UNIQUE_NAME[len(item.UNIQUE_NAME)-2] == '@' {
		enchantment = int(item.UNIQUE_NAME[len(item.UNIQUE_NAME)-1] - '0')
	} else {
		enchantment = -1
	}
	item_readable.ENCHANTMENT = enchantment

	return item_readable
}
