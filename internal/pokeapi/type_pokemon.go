package pokeapi

type RespPokemon struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	BaseExperience int       `json:"base_experience"`
	Height         int       `json:"height"`
	IsDefault      bool      `json:"is_default"`
	Order          int       `json:"order"`
	Weight         int       `json:"weight"`
	Abilities      Abilities `json:"abilities"`
	Forms          []NameURL `json:"forms"`
	GameIndices    []struct {
		GameIndex int     `json:"game_index"`
		Version   NameURL `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item           NameURL `json:"item"`
		VersionDetails []struct {
			Rarity  int     `json:"rarity"`
			Version NameURL `json:"version"`
		}
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move                NameURL `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int     `json:"level_learned_at"`
			VersionGroup    NameURL `json:"version-group"`
			MoveLearnMethod NameURL `json:"move_learn_method"`
			Order           int     `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species NameURL `json:"species"`
	Sprites struct {
		SpriteAll
		Other struct {
			DreamWorld      SpriteFront2 `json:"dream_world"`
			Home            SpriteGenVI  `json:"home"`
			OfficialArtwork SpriteFront  `json:"official-artwork"`
			Showdown        SpriteAll    `json:"showdown"`
			Versions        struct {
				GernerationI struct {
					RedBlue SpriteGenI `json:"red-blue"`
					Yellow  SpriteGenI `json:"yellow"`
				} `json:"generation-i"`
				GenerationII struct {
					Crystal SpriteGenII `json:"crystal"`
					Gold    SpriteGenII `json:"gold"`
					Silver  SpriteGenII `json:"silver"`
				} `json:"generation-ii"`
				GenerationIII struct {
					Emerald          SpriteFront `json:"emerald"`
					FireredLeafgreen SpriteGenII `json:"firered-leafgreen"`
					RubySapphire     SpriteGenII `json:"ruby-sapphire"`
				} `json:"generation-iii"`
				GenerationIV struct {
					DiamondPearl        SpriteAll `json:"diamond-pearl"`
					HeartgoldSoulsilver SpriteAll `json:"heartgold-soulsilver"`
					Platinum            SpriteAll `json:"platinum"`
				} `json:"generation-iv"`
				GenerationV struct {
					BlackWhite struct {
						Animated SpriteAll `json:"animated"`
						SpriteAll
					} `json:"black-white"`
				} `json:"generation-v"`
				GenerationVI struct {
					OmegarubyAlphasapphire SpriteGenVI `json:"omegaruby-alphasapphire"`
					XY                     SpriteGenVI `json:"x-y"`
				} `json:"generation-vi"`
				GenerationVII struct {
					Icons             SpriteFront2 `json:"icons"`
					UltraSunUltraMoon SpriteGenVI  `json:"ultra-sun-ultra-moon"`
				} `json:"generation-vii"`
				GenerationVIII struct {
					Icons SpriteFront2 `json:"icons"`
				} `json:"generation-viii"`
			} `json:"versions"`
		} `json:"other"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int     `json:"base_stat"`
		Effort   int     `json:"effort"`
		Stat     NameURL `json:"stat"`
	} `json:"stats"`
	Types     Types `json:"types"`
	PastTypes []struct {
		Generation NameURL `json:"generation"`
		Types      Types   `json:"types"`
	} `json:"past_types"`
	PastAbilities []struct {
		Generation NameURL   `json:"generation"`
		Abilities  Abilities `json:"abilities"`
	} `json:"past_abilities"`
}

type SpriteAll struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemail      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemail string `json:"front_shiny_female"`
}

type SpriteGenI struct {
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
	FrontGray    string `json:"front_gray"`
}

type SpriteGenII struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type SpriteFront struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type SpriteFront2 struct {
	FrontDefault string `json:"front_default"`
	FrontFemail  string `json:"front_female"`
}

type SpriteGenVI struct {
	FrontDefault     string `json:"front_default"`
	FrontFemail      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemail string `json:"front_shiny_female"`
}

type Types []struct {
	Slot int     `json:"slot"`
	Type NameURL `json:"type"`
}

type Abilities []struct {
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
	Ability  NameURL `json:"ability"`
}

