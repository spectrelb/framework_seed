package utils

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"go.uber.org/zap"
	"io"
)

func ReadExcel(file io.Reader) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		zap.L().Error("ReadExcel err:", zap.Error(err))
		return
	}
	f.GetActiveSheetIndex()
}
