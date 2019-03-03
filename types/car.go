package types

type Car struct {
	Manufacturer  string `json:"manufacturer" bson:"manufacturer"`
	Model         string `json:"model" bson:"model"`
	License       string `json:"license" bson:"license"`
	EngineType    string `json:"engine_type" bson:"engine_type"`
	VIN           string `json:"vin" bson:"vin"`
	EngineCC      uint32 `json:"engine_cc" bson:"engine_cc"`
	EngineKW      uint32 `json:"engine_kw" bsob:"engine_kw"`
	EngineKM      uint32 `json:"engine_km" bson:"engine_km"`
	ManufacturerY uint32 `json:"manufacturer_y" bson:"manufacturer_y"`
	ManufacturerM uint32 `json:"manufacturer_m" bson:"manufacturer_m"`
}
