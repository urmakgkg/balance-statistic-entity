package balance_statistic_entity

type AccrualDetail struct {
	Ls      string
	Key     string
	Period  string
	Type    int
	TypeOtu int
	NameOtu string
	Tariff  float64
	Kwh     float64
	Sum     float64
	Nds     float64
	Nsp     float64
}
