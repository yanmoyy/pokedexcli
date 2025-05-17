package pokeapi

// RespShallowLocations -
type RespShallowLocations struct {
	Results  []Location `json:"results"`
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
}

// RespLocationArea
type RespLocationArea struct {
	Location             Location              `json:"location"`
	Name                 string                `json:"name"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
	GameIndex            int                   `json:"game_index"`
}

type EncounterMethodRate struct {
	EncounterMethod struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"encounter_method"`
	VersionDetails []struct {
		Rate    int     `json:"rate"`
		Version Version `json:"version"`
	} `json:"version_details"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Name struct {
	Language struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"language"`
	Name string `json:"name"`
}

type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
	VersionDetails []struct {
		EncounterDetails []struct {
			Chance          int      `json:"chance"`
			ConditionValues []string `json:"condition_values"`
			MaxLevel        int      `json:"max_level"`
			Method          struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"method"`
			MinLevel int `json:"min_level"`
		} `json:"encounter_details"`
		MaxChance int     `json:"max_chance"`
		Version   Version `json:"version"`
	} `json:"version_details"`
}
