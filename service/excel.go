package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
	"heritage-go/model"
	"heritage-go/repository"

	"gorm.io/gorm"
)

// ImportFromExcel 解析上传的 Excel 文件并进行批量原子化写入
func ImportFromExcel(filePath string) (int, error) {
	// 1. 打开 Excel 文件
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("读取 Excel 文件失败: %w", err)
	}
	defer f.Close()

	// 默认读取第一个工作表
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return 0, fmt.Errorf("读取工作表数据失败: %w", err)
	}

	if len(rows) <= 1 {
		return 0, fmt.Errorf("Excel 文件内未检测到有效的家族成员数据")
	}

	var members []model.Member

	// 2. 遍历数据行（跳过第一行表头）
	for i, row := range rows {
		if i == 0 {
			continue 
		}
		// 容错处理：防止空行或不规整的数据行
		if len(row) < 2 || strings.TrimSpace(row[1]) == "" {
			continue
		}

		// 解析各个字段
		id, _ := strconv.Atoi(strings.TrimSpace(row[0]))
		name := strings.TrimSpace(row[1])
		
		spouse := ""
		if len(row) > 2 {
			spouse = strings.TrimSpace(row[2])
		}

		fatherID := 0
		if len(row) > 3 && strings.TrimSpace(row[3]) != "" {
			fatherID, _ = strconv.Atoi(strings.TrimSpace(row[3]))
		}

		generation := 1
		if len(row) > 4 && strings.TrimSpace(row[4]) != "" {
			generation, _ = strconv.Atoi(strings.TrimSpace(row[4]))
		}

		members = append(members, model.Member{
			ID:         id,
			Name:       name,
			Spouse:     spouse,
			FatherID:   fatherID,
			Generation: generation,
		})
	}

	// 3. 执行数据库批量安全写入（采用 GORM 事务保护机制）
	if len(members) > 0 {
		err := repository.DB.Transaction(func(tx *gorm.DB) error {
			// 先清空旧数据（或者你可以根据业务改成 Upsert 改写）
			if err := tx.Exec("DELETE FROM members").Error; err != nil {
				return err
			}
			// 批量插入
			if err := tx.Create(&members).Error; err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return 0, fmt.Errorf("数据库批量写入失败: %w", err)
		}
	}

	return len(members), nil
}
