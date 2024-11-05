package commands

import (
	"errors"
	"fmt"
	utils2 "github.com/sonhineboy/gsadminCli/pkg/utils"
	"github.com/sonhineboy/gsadminCli/tmp"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strings"
	"text/template"
)

var (
	IgnoreField = map[string]bool{
		"id":         true,
		"created_at": true,
		"updated_at": true,
		"deleted_at": true,
	}

	FieldTypeMapping = map[string]string{
		"varchar":   "string",
		"text":      "string",
		"timestamp": "*LocalTime",
		"bigint":    "int64",
		"int":       "int32",
		"tinyint":   "int8",
		"float":     "float64",
		"decimal":   "string",
		"longtext":  "string",
		"image":     "string",
	}
)

type MakeModelCommand struct {
	MakeBaseCommand
}

func (m *MakeModelCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:          fmt.Sprintf("make:%s", m.makeName),
		Short:        fmt.Sprintf("create a %s;", m.makeName),
		Example:      fmt.Sprintf("make:%s ./models/user.go -t=users -d='root:123@tcp(localhost:3306)/demo'", m.makeName),
		SilenceUsage: true,
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				return err
			}
			if err := utils2.ValidateFormat(args[0]); err != nil {
				return err
			}
			if err := utils2.ValidateIsPackage(args[0], m.pk); err != nil {
				return err
			}
			return nil
		},
		Run: m.Run,
	}

	cmd.GroupID = m.GroupId
	cmd.PersistentFlags().StringVar(&m.pk, "package", "", "如果在根目录创建需要传递 --package=")
	cmd.Flags().StringP("dns", "d", "", "数据库连接dns --dns=")
	cmd.Flags().StringP("table", "t", "", "表 --table=")
	return cmd
}

func NewMakeModelCommand(gId string) *MakeModelCommand {
	m := new(MakeModelCommand)
	m.SetMakeName("model")
	m.SetTmpFunc(tmp.ModelTmp)
	m.SetGroupId(gId)
	return m
}

func (m *MakeModelCommand) Run(cmd *cobra.Command, args []string) {

	dns, _ := cmd.Flags().GetString("dns")
	table, _ := cmd.Flags().GetString("table")

	if len(dns) == 0 || dns == "=" || len(table) == 0 || table == "=" {
		cmd.PrintErr(errors.New("flags dns and table is required"))
		return
	}

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		cmd.PrintErr(err)
		return
	}

	fieldTypes, err := db.Migrator().ColumnTypes(table)
	if err != nil {
		cmd.PrintErr(err)
		return
	}

	fieldMaxLen := 0
	fieldTypeMaxLen := 0

	var fields = make([][]string, 0, len(fieldTypes))

	for _, columnType := range fieldTypes {
		if fieldMaxLen < len(columnType.Name()) {
			fieldMaxLen = len(columnType.Name())
		}

		if _, ok := IgnoreField[columnType.Name()]; ok {
			continue
		}
		fieldType, ok := FieldTypeMapping[columnType.DatabaseTypeName()]

		if !ok {
			fieldType = "string"
		}

		if fieldTypeMaxLen < len(fieldType) {
			fieldTypeMaxLen = len(fieldType)
		}

		fields = append(fields, []string{
			utils2.ToCamelCase(columnType.Name(), true),
			fieldType,
			fmt.Sprintf("`gorm:\"column:%s;json:%s\"`", columnType.Name(), columnType.Name()),
		})
	}

	err = utils2.CreateFileCallBack(args[0], func(operatorFile *utils2.FileOperator, file *os.File) error {
		funcMap := template.FuncMap{
			"toLower": strings.ToLower,
			"transField": func(field []string) string {
				return fmt.Sprintf("%-*s%-*s%s", fieldMaxLen+2, field[0], fieldTypeMaxLen+2, field[1], field[2])
			},
		}
		tem, err := template.New(m.makeName).Funcs(funcMap).Parse(m.tmpFunc())
		if err != nil {
			cmd.PrintErr(err)
		}

		if len(m.pk) == 0 {
			m.pk = utils2.GetPackage(operatorFile.Dir)
		}

		err = tem.Execute(file, map[string]interface{}{
			"Name":    operatorFile.NameToTitle(m.makeName),
			"Package": strings.ToLower(m.pk),
			"Table":   table,
			"Fields":  fields,
		})
		if err != nil {
			return err
		}
		cmd.Printf("This is %s created", operatorFile.AllPath)
		return nil
	})
	if err != nil {
		cmd.PrintErr(err)
	}
}
