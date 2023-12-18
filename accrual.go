package balance_statistic_entity

import "time"

type Accrual struct {
	ID                         int        `json:"id"`
	Ls                         string     `json:"ls"`
	Key                        string     `json:"key"`
	Period                     string     `json:"period"`
	DateBegin                  *time.Time `json:"date_begin"`
	DateEnd                    *time.Time `json:"date_end"`
	CounterDisplayPrevious     float64    `json:"counter_display_previous"`
	CounterDisplayNext         float64    `json:"counter_display_next"`
	CounterDisplayPreviousType string     `json:"counter_display_previous_type"`
	CounterDisplayNextType     string     `json:"counter_display_next_type"`
	Days                       int        `json:"days"`
	Kwh                        float64    `json:"kwh"`
	Sum                        float64    `json:"sum"`
	SumNsp                     float64    `json:"sum_nsp"`
	SumNds                     float64    `json:"sum_nds"`
	SumPenalty                 float64    `json:"sum_penalty"`
	KwhOtu                     float64    `json:"kwh_otu"`
	SumOtu                     float64    `json:"sum_otu"`
	SumNspOtu                  float64    `json:"sum_nsp_otu"`
	SumNdsOtu                  float64    `json:"sum_nds_otu"`
	GeneralKwh                 float64    `json:"general_kwh"`
	GeneralSum                 float64    `json:"general_sum"`
	SysDate                    time.Time  `json:"sys_date"`
	Header
}
