package repositories

import (
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/errors"
)

type CommonDataRepository interface {
	GetCharacteristics() ([]model.KeyValuePair, error)
	GetNatures() ([]model.KeyValuePair, error)
	GetGames() ([]model.KeyValuePair, error)
	GetMethods() ([]model.KeyValuePair, error)
	GetGenders() ([]model.KeyValuePair, error)
	GetCharacteristic(key string) (*model.KeyValuePair, error)
	GetNature(key string) (*model.KeyValuePair, error)
	GetGame(key string) (*model.KeyValuePair, error)
	GetMethod(key string) (*model.KeyValuePair, error)
	GetGender(key string) (*model.KeyValuePair, error)
}

type StaticCommonDataRepository struct {
	characteristics []model.KeyValuePair
	natures         []model.KeyValuePair
	games           []model.KeyValuePair
	methods         []model.KeyValuePair
	genders         []model.KeyValuePair
}

func NewCommonDataService() CommonDataRepository {
	cds := &StaticCommonDataRepository{}
	cds.loadCharacteristics()
	cds.loadNatures()
	cds.loadGames()
	cds.loadMethods()
	cds.loadGenders()
	return cds
}

func (cds *StaticCommonDataRepository) GetCharacteristics() ([]model.KeyValuePair, error) {
	return cds.characteristics, nil
}

func (cds *StaticCommonDataRepository) GetNatures() ([]model.KeyValuePair, error) {
	return cds.natures, nil
}

func (cds *StaticCommonDataRepository) GetGames() ([]model.KeyValuePair, error) {
	return cds.games, nil
}

func (cds *StaticCommonDataRepository) GetMethods() ([]model.KeyValuePair, error) {
	return cds.methods, nil
}

func (cds *StaticCommonDataRepository) GetGenders() ([]model.KeyValuePair, error) {
	return cds.genders, nil
}

func (cds *StaticCommonDataRepository) GetGame(key string) (*model.KeyValuePair, error) {
	for i, g := range cds.games {
		if g.Key == key {
			return &cds.games[i], nil
		}
	}
	return nil, &errors.InvalidGame
}

func (cds *StaticCommonDataRepository) GetMethod(key string) (*model.KeyValuePair, error) {
	for i, m := range cds.methods {
		if m.Key == key {
			return &cds.methods[i], nil
		}
	}
	return nil, &errors.InvalidMethod
}

func (cds *StaticCommonDataRepository) GetCharacteristic(key string) (*model.KeyValuePair, error) {
	for i, c := range cds.characteristics {
		if c.Key == key {
			return &cds.characteristics[i], nil
		}
	}
	return nil, &errors.InvalidCharacteristic
}

func (cds *StaticCommonDataRepository) GetNature(key string) (*model.KeyValuePair, error) {
	for i, n := range cds.natures {
		if n.Key == key {
			return &cds.natures[i], nil
		}
	}
	return nil, &errors.InvalidNature
}

func (cds *StaticCommonDataRepository) GetGender(key string) (*model.KeyValuePair, error) {
	for i, n := range cds.genders {
		if n.Key == key {
			return &cds.genders[i], nil
		}
	}
	return nil, &errors.InvalidGender
}

func (cds *StaticCommonDataRepository) loadNatures() {
	cds.natures = []model.KeyValuePair{
		{Key: "ADAMANT", Value: "Adamant"},
		{Key: "BASHFUL", Value: "Bashful"},
		{Key: "BOLD", Value: "Bold"},
		{Key: "BRAVE", Value: "Brave"},
		{Key: "CALM", Value: "Calm"},
		{Key: "CAREFUL", Value: "Careful"},
		{Key: "DOCILE", Value: "Docile"},
		{Key: "GENTLE", Value: "Gentle"},
		{Key: "HARDY", Value: "Hardy"},
		{Key: "HASTY", Value: "Hasty"},
		{Key: "IMPISH", Value: "Impish"},
		{Key: "JOLLY", Value: "Jolly"},
		{Key: "LAX", Value: "Lax"},
		{Key: "LONELY", Value: "Lonely"},
		{Key: "MILD", Value: "Mild"},
		{Key: "MODEST", Value: "Modest"},
		{Key: "NAIVE", Value: "Naive"},
		{Key: "NAUGHTY", Value: "Naughty"},
		{Key: "QUIET", Value: "Quiet"},
		{Key: "QUIRKY", Value: "Quirky"},
		{Key: "RASH", Value: "Rash"},
		{Key: "RELAXED", Value: "Relaxed"},
		{Key: "SASSY", Value: "Sassy"},
		{Key: "SERIOUS", Value: "Serious"},
		{Key: "TIMID", Value: "Timid"},
	}
}

func (cds *StaticCommonDataRepository) loadCharacteristics() {
	cds.characteristics = []model.KeyValuePair{
		{Key: "EAT", Value: "It loves to eat!"},
		{Key: "SIESTAS", Value: "It takes plenty of siestas!"},
		{Key: "NODS_OFF", Value: "It nods off a lot!"},
		{Key: "SCATTERS", Value: "It scatters things often!"},
		{Key: "RELAX", Value: "It likes to relax!"},
		{Key: "POWER", Value: "It's proud of its power!"},
		{Key: "THRASHES", Value: "It likes to thrash about!"},
		{Key: "LITTLE_TEMPER", Value: "It's a little quick tempered!"},
		{Key: "FIGHTS", Value: "It likes to fight!"},
		{Key: "QUICK_TEMPERED", Value: "It's quick tempered!"},
		{Key: "STURDY", Value: "It has a sturdy body!"},
		{Key: "TAKES_HITS", Value: "It's capable of taking hits!"},
		{Key: "PERSISTENT", Value: "It's highly persistent!"},
		{Key: "ENDURANCE", Value: "It has good endurance!"},
		{Key: "PERSEVERANCE", Value: "It has good perseverance!"},
		{Key: "RUNS", Value: "It likes to run!"},
		{Key: "ALERT", Value: "It's alert to sounds!"},
		{Key: "IMPETUOUS", Value: "It's impetuous and silly!"},
		{Key: "CLOWN", Value: "It's somewhat of a clown!"},
		{Key: "FLEES", Value: "It's quick to flee!"},
		{Key: "CURIOUS", Value: "It's highly curious!"},
		{Key: "MISCHIEVOUS", Value: "It's mischievous!"},
		{Key: "CUNNING", Value: "It's thoroughly cunning!"},
		{Key: "LOST", Value: "It's often lost in thought!"},
		{Key: "FINICKY", Value: "It's very finicky!"},
		{Key: "WILLED", Value: "It's strong willed!"},
		{Key: "VAIN", Value: "It's somewhat vain!"},
		{Key: "DEFIANT", Value: "It's strongly defiant!"},
		{Key: "NO_LOSE", Value: "It hates to lose!"},
		{Key: "STUBBORN", Value: "It's somewhat stubborn!"},
	}
}

func (cds *StaticCommonDataRepository) loadGames() {
	cds.games = []model.KeyValuePair{
		{Key: "X", Value: "X"},
		{Key: "Y", Value: "Y"},
		{Key: "OMEGA_RUBY", Value: "Omega Ruby"},
		{Key: "ALPHA_SAPPHIRE", Value: "Alpha Sapphire"},
		{Key: "SUN", Value: "Sun"},
		{Key: "MOON", Value: "Moon"},
		{Key: "ULTRA_SUN", Value: "Ultra Sun"},
		{Key: "ULTRA_MOON", Value: "Ultra Moon"},
		{Key: "BRILLIANT_DIAMOND", Value: "Brilliant Diamond"},
		{Key: "SHINING_PEARL", Value: "Shining Pearl"},
		{Key: "LEGENDS_ARCEUS", Value: "Legends: Arceus"},
		{Key: "SWORD", Value: "Sword"},
		{Key: "SHIELD", Value: "Shield"},
		{Key: "SCARLET", Value: "Scarlet"},
		{Key: "VIOLET", Value: "Violet"},
		{Key: "LZA", Value: "Legends: Z-A"},
	}
}

func (cds *StaticCommonDataRepository) loadMethods() {
	cds.methods = []model.KeyValuePair{
		{Key: "RANDOM", Value: "Random Encounter"},
		{Key: "FISHING", Value: "Chain Fishing"},
		{Key: "FRIEND_SAFARI", Value: "Friend Safari"},
		{Key: "DEXNAV", Value: "DexNav"},
		{Key: "MASUDA", Value: "Masuda Method"},
		{Key: "HORDE", Value: "Horde Hunting"},
		{Key: "POKE_RADAR", Value: "PokeRadar"},
		{Key: "SOS", Value: "S.O.S."},
		{Key: "ULTRA_WORMHOLE", Value: "Ultra Wormhole"},
		{Key: "CATCH_COMBO", Value: "Catch Combo"},
		{Key: "DYNAMAX", Value: "Dynamax Adventure"},
		{Key: "MASS_OUTBREAK", Value: "Mass Outbreaks"},
		{Key: "MMO", Value: "Massive Mass Outbreaks"},
		{Key: "STATIC", Value: "Static Encounter"},
		{Key: "SANDWICH", Value: "Sandwich Hunting"},
		{Key: "SV_MASS_OUTBREAK", Value: "Scarlet/Violet Mass Outbreaks"},
		{Key: "SPARKLING_DONUT", Value: "Sparkling Donut"},
	}
}

func (cds *StaticCommonDataRepository) loadGenders() {
	cds.genders = []model.KeyValuePair{
		{Key: "M", Value: "Male"},
		{Key: "F", Value: "Female"},
		{Key: "U", Value: "Unknown"},
	}
}
