package mysql

import (
	"database/sql"

	"tanjunchen.io.webapp/models"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func GetCategoryList() (categoryList []*models.Category, err error) {

	sqlStr := "select category_id, category_name from category"
	err = db.Select(&categoryList, sqlStr)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	return
}

func MGetCategory(categoryIDs []int32) (categoryMap map[int32]*models.Category, err error) {

	sqlStr := "select category_id, category_name from category where category_id in(?)"

	inSqlStr, params, err := sqlx.In(sqlStr, categoryIDs)
	if err != nil {
		zap.L().Error("sqlx.In failed", zap.String("sqlStr", sqlStr), zap.Error(err))
		return
	}

	categoryMap = make(map[int32]*models.Category, len(categoryIDs))
	var categoryList []*models.Category
	err = db.Select(&categoryList, inSqlStr, params...)
	if err != nil {
		zap.L().Error("MGetCategory failed", zap.Any("categoryIDs", categoryIDs), zap.Error(err))
		return
	}

	for _, v := range categoryList {
		categoryMap[v.CategoryID] = v
	}
	return
}
