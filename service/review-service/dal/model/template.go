package model

import "encoding/json"

type CommonModel struct {
	ID      int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	ExtJSON string `gorm:"column:ext_json;not null;comment:信息扩展" json:"ext_json"`        // 信息扩展
}

func (c *CommonModel) ExtJSONMap() map[string]interface{} {
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(c.ExtJSON), &data)
	if err != nil {
		return nil
	}
	return data
}
