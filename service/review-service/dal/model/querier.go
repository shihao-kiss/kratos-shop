// dal/model/querier.go

package model

import "gorm.io/gen"

// 通过添加注释生成自定义方法

type Querier interface {
	// SELECT * FROM @@table WHERE version=@version
	GetByVersion(version int) ([]*gen.T, error) // 返回结构体和error
}
