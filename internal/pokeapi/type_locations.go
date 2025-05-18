package pokeapi

// RespShallowLocations -
type RespShallowLocations struct {
	Results  []NameURL `json:"results"`
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
}

// RespLocationArea
type RespLocationArea struct {
	Location             NameURL               `json:"location"`
	Name                 string                `json:"name"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
	GameIndex            int                   `json:"game_index"`
}

type EncounterMethodRate struct {
	EncounterMethod NameURL `json:"encounter_method"`
	VersionDetails  []struct {
		Rate    int     `json:"rate"`
		Version NameURL `json:"version"`
	} `json:"version_details"`
}
type Name struct {
	Language NameURL `json:"language"`
	Name     string  `json:"name"`
}

type PokemonEncounter struct {
	Pokemon        NameURL `json:"pokemon"`
	VersionDetails []struct {
		EncounterDetails []struct {
			Chance          int      `json:"chance"`
			ConditionValues []string `json:"condition_values"`
			MaxLevel        int      `json:"max_level"`
			Method          NameURL  `json:"method"`
			MinLevel        int      `json:"min_level"`
		} `json:"encounter_details"`
		MaxChance int     `json:"max_chance"`
		Version   NameURL `json:"version"`
	} `json:"version_details"`
}
