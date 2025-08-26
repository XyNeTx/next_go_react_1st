package model

import "time"

type TB_Import_Error struct {
	F_PDS_CD      string    `json:"F_PDS_CD"`
	F_Row         int       `json:"F_Row"`
	F_Field       string    `json:"F_Field"`
	F_Remark      string    `json:"F_Remark"`
	F_Update_By   string    `json:"F_Update_By"`
	F_Update_Date time.Time `json:"F_Update_Date"`
	F_Type        string    `json:"F_Type"`
}

func (TB_Import_Error) TableName() string {
	return "TB_Import_Error"
}
