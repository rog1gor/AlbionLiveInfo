package albionAPI

import "database/sql"

var AlbionApiURL string = "https://west.albion-online-data.com/"

var PricesPrefix string = "api/v2/stats/prices/"

var LocationsQuery string = "locations="
var QualitiesQuery string = "qualities="
var ArgumentSeparator = ","
var QueryConcat string = "&"
var SufixQueryStart = "?"

var allMarketsQuery string = LocationsQuery + "3005,0007,1002,2004,3008,4002,3013-Auction2,5003"
var allQualitiesQuery string = QualitiesQuery + "1,2,3,4,5"
var queryAll string = SufixQueryStart + allMarketsQuery + QueryConcat + allQualitiesQuery

var HttpsQueryMaxLen int = 4096
var AllTypesAndMarketsQuerySufixLen int = len(SufixQueryStart) + len(allMarketsQuery) + len(QueryConcat) + len(allQualitiesQuery)
var AllTypesAndMarketsQuerySpaceLeft int = HttpsQueryMaxLen - AllTypesAndMarketsQuerySufixLen

var AlbionDatabase string = "Albion.db"

func InitializeDbConnection() (*sql.DB, error) {
	return sql.Open("sqlite3", AlbionDatabase)
}
