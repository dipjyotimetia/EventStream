// Code generated by avrogen. DO NOT EDIT.

package expense

import (
	"github.com/heetch/avro/avrotypegen"
)

type Expense struct {
	Expense_id  string  `json:"expense_id" avro:"expense_id"`
	User_id     string  `json:"user_id" avro:"user_id"`
	Category    string  `json:"category" avro:"category"`
	Amount      float64 `json:"amount" avro:"amount"`
	Currency    string  `json:"currency" avro:"currency"`
	Timestamp   int64   `json:"timestamp" avro:"timestamp"`
	Description *string `json:"description" avro:"description"`
	Receipt     *string `json:"receipt" avro:"receipt"`
}

// AvroRecord implements the avro.AvroRecord interface.
func (Expense) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"expense_id","type":"string"},{"name":"user_id","type":"string"},{"name":"category","type":"string"},{"name":"amount","type":"double"},{"name":"currency","type":"string"},{"logicalType":"timestamp-millis","name":"timestamp","type":"long"},{"default":null,"name":"description","type":["null","string"]},{"default":null,"name":"receipt","type":["null","string"]}],"name":"com.expense.Expense","type":"record"}`,
		Required: []bool{
			0: true,
			1: true,
			2: true,
			3: true,
			4: true,
			5: true,
		},
	}
}