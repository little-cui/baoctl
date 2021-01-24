package address

import (
	"baoctl/pkg/cmd"
	"baoctl/pkg/cmd/command"
	"baoctl/pkg/config"
	"baoctl/pkg/tmpl"
	"context"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strings"
	"time"
)

const (
	SheetName = "Sheet1"
	Layout    = "20060102150405.xlsx"
)

func init() {
	cmd.RegisterCommand(&command.Command{
		Code: 1,
		Desc: `待发货订单xlsx转换`,
		Action: func(ctx context.Context, args ...interface{}) error {
			return Action(ctx)
		},
	})
}

func Action(ctx context.Context) error {
	cf := config.Config().Xlsx

	fmt.Print(tmpl.CmdXlsxPath)
	fmt.Scanln(&cf.FilePath)
	fmt.Print(tmpl.CmdXlsxPWD)
	fmt.Scanln(&cf.Password)
	fmt.Print(tmpl.CmdXlsxSheetN)
	fmt.Scanln(&cf.SheetN)

	opts := excelize.Options{}
	if cf.Password != "" {
		opts.Password = cf.Password
	}
	file, err := excelize.OpenFile(cf.FilePath, opts)
	if err != nil {
		return err
	}

	number := cf.SheetN
	if number <= 0 {
		number = 1
	}
	sheetName := file.GetSheetName(number - 1)
	rows, err := file.GetRows(sheetName)
	if err != nil {
		return err
	}
	return convert(rows)
}

func convert(rows [][]string) error {
	if len(rows) <= 0 {
		return fmt.Errorf("no data")
	}

	headerIndexes, fixedHeaderIndex, err := parseHeaderIndexes(rows[0])
	if err != nil {
		return err
	}

	return export(rows, headerIndexes, fixedHeaderIndex)
}

func parseHeaderIndexes(headers []string) ([]int, int, error) {
	cf := config.Config().Xlsx
	var (
		headerIndexes    []int
		fixedHeaderIndex = -1
	)
	for _, header := range cf.FixedHeaders {
		header = strings.TrimSpace(header)
		for i, h := range headers {
			h = strings.TrimSpace(h)
			if fixedHeaderIndex < 0 && h == cf.FixedHeader {
				fixedHeaderIndex = i
			}
			if h == header {
				headerIndexes = append(headerIndexes, i)
				break
			}
		}
	}
	if fixedHeaderIndex < 0 {
		return nil, -1, fmt.Errorf("no fixed header %s", cf.FixedHeader)
	}
	return headerIndexes, fixedHeaderIndex, nil
}

func export(rows [][]string, headerIndexes []int, fixedHeaderIndex int) error {
	file := excelize.NewFile()
	file.NewSheet(SheetName)
	for r, row := range rows {
		rowN := r + 1
		addCols := 0
		for colN, i := range headerIndexes {
			value := ConvertHeader{Full: strings.TrimSpace(row[i])}
			colN += addCols
			if i == fixedHeaderIndex {
				value = parseFixedHeader(value)
				addCols = len(value.Spilt)
				for step := 0; step < addCols; step++ {
					if err := file.SetCellValue(SheetName, fmt.Sprintf("%c%d", 'A'+colN+step, rowN), value.Spilt[step]); err != nil {
						return err
					}
				}
				colN += addCols
			}
			if err := file.SetCellValue(SheetName, fmt.Sprintf("%c%d", 'A'+colN, rowN), value.Full); err != nil {
				return err
			}
		}
	}
	return file.SaveAs(time.Now().Format(Layout))
}

func parseFixedHeader(value ConvertHeader) ConvertHeader {
	cf := config.Config().Xlsx
	l := len(cf.AddedHeaders)
	arr := strings.Split(value.Full, " ")
	if len(arr) <= l {
		value.Spilt = cf.AddedHeaders
		return value
	}
	return ConvertHeader{
		Spilt: arr[:l],                      // 增加省市区三列
		Full:  strings.Join(arr[l-1:], " "), // 原来的列删除省市信息
	}
}
