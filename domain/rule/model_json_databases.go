package rule

// Armor struct
type Armor struct {
	Name              string `json:"name"`
	Title             string `json:"title"`
	Kind              string `json:"kind"`
	Cost              int    `json:"cost"`
	CoinType          string `json:"coin_type"`
	ArmorClass        int    `json:"armor_class"`
	DexterityModifier int    `json:"dexterity_modifier"`
	Stealth           bool   `json:"stealth"`
	Strength          int    `json:"strength"`
	Weight            int    `json:"weight"`
	Measure           string `json:"measure"`
}

// Weapon struct
type Weapon struct {
	Name           string `json:"name"`
	Title          string `json:"title"`
	Kind           string `json:"kind"`
	Cost           int    `json:"cost"`
	CoinType       string `json:"coin_type"`
	Damage         string `json:"damage"`
	DamageTwoHands string `json:"damage_two_hands,omitempty"`
	DamageType     string `json:"damage_type"`
	Weight         int    `json:"weight"`
	Measure        string `json:"measure"`
	Properties     string `json:"properties"`
}

// Gear struct
type Gear struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Kind     string `json:"kind"`
	Cost     int    `json:"cost"`
	CoinType string `json:"coin_type"`
	Weight   int    `json:"weight"`
	Measure  string `json:"measure"`
	Number   int    `json:"number"`
}

// Packs struct
type Packs struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	CoinType    string `json:"coin_type"`
}

// Tools struct
type Tools struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Kind        string `json:"kind"`
	Cost        int    `json:"cost"`
	CoinType    string `json:"coin_type"`
	Weight      int    `json:"weight"`
	Measure     string `json:"measure"`
	Description string `json:"description"`
}

// Mounts struct
type Mounts struct {
	Name                    string `json:"name"`
	Title                   string `json:"title"`
	Cost                    int    `json:"cost"`
	CoinType                string `json:"coin_type"`
	CarryingCapacity        int    `json:"carrying_capacity"`
	CarryingCapacityMeasure string `json:"carrying_capacity_measure"`
	Speed                   int    `json:"speed"`
	SpeedMeasure            string `json:"speed_measure"`
}

// Services struct
type Services struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Cost     int    `json:"cost"`
	CoinType string `json:"coin_type"`
	Unit     string `json:"unit"`
	Source   string `json:"source"`
}

// TreasureHoard struct
type TreasureHoard struct {
	Name        string `json:"name"`
	Kind        string `json:"kind"`
	Description string `json:"description"`
	Value       int    `json:"value"`
	CoinType    string `json:"coin_type"`
}
