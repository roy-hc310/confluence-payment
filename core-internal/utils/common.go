package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func PrepareInsertQuery(tableName string, colunmsName []string, promotionsInterface [][]interface{}) string {

	promotionLen := len(promotionsInterface)
	fieldLen := len(promotionsInterface[0])

	var placeholders []string
	for i := 0; i < promotionLen; i++ {
		var recordPlaceholders []string
		for j := 0; j < fieldLen; j++ {
			recordPlaceholders = append(recordPlaceholders, fmt.Sprintf("$%d", i*fieldLen+j+1))
		}
		placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(recordPlaceholders, ", ")))
	}
	finalPlaceholder := strings.Join(placeholders, ", ")

	queryString := fmt.Sprintf(`INSERT INTO %s (%s) VALUES %s`, tableName, strings.Join(colunmsName, ", "), finalPlaceholder)
	return queryString
}

func PrepareUpdateQuery(tableName string, colunmsName []string) string {

	var placeholders []string
	for i := 0; i < len(colunmsName); i++ {
		placeholders = append(placeholders, fmt.Sprintf("%s = $%d", colunmsName[i], i+1))
	}
	finalPlaceholder := strings.Join(placeholders, ", ")

	queryString := fmt.Sprintf(`UPDATE %s SET %s WHERE uuid = %s AND deleted_at IS NULL `, tableName, finalPlaceholder, fmt.Sprintf("$%d", len(colunmsName)+1))

	return queryString
}

func PrepareSelectQueryForUpdate(tableName string, colunmsName []string) string {
	finalPlaceholder := strings.Join(colunmsName, ", ")

	queryString := fmt.Sprintf(`SELECT %s FROM %s WHERE deleted_at IS NULL AND uuid = $1 FOR UPDATE `, finalPlaceholder, tableName)

	return queryString
}

func PrepareSelectQuery(tableName string, colunmsName []string) string {

	finalPlaceholder := strings.Join(colunmsName, ", ")

	queryString := fmt.Sprintf(`SELECT %s FROM %s WHERE deleted_at IS NULL `, finalPlaceholder, tableName)

	return queryString
}

func PrepareSelectCountQuery(tableName string) string {
	queryString := fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE deleted_at IS NULL `, tableName)

	return queryString
}

func IsValidUUID(value string) bool {
	if value == "" {
		return false
	}
	_, err := uuid.Parse(value)

	return err == nil
}

func StringToInt(value string) int {
	if value == "" {
		return 0
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return intValue
}

func IntToString(value int) string {
	return strconv.Itoa(value)
}

func StringToFloat(value string) float64 {
	if value == "" {
		return 0
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return floatValue
}