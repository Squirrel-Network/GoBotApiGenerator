package sub_build

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/interfaces"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"BotApiCompiler/consts"
	"fmt"
	"sort"
	"strings"
)

func BuildType[Scheme interfaces.SchemeInterface](typeScheme Scheme, builder *component.Context) []types.FieldTL {
	var filesInput []types.FieldTL
	var structName string
	isMethod := typeScheme.GetType() == "methods"
	structName = utils.FixStructName(typeScheme.GetName())
	if len(typeScheme.GetFields()) == 0 {
		if !isMethod {
			builder.AddType(utils.FixStructName(structName), "interface{}").AddLine()
		} else {
			builder.AddType(utils.FixStructName(structName), "struct{}").AddLine()
		}
	} else {
		typesOrdered := make([]types.FieldTL, 0, len(typeScheme.GetFields()))
		for _, field := range typeScheme.GetFields() {
			typesOrdered = append(typesOrdered, field)
		}
		sort.Slice(typesOrdered, func(i, j int) bool {
			return typesOrdered[i].Name < typesOrdered[j].Name
		})
		builder.InitStruct(structName)
		for _, field := range typesOrdered {
			jsonName := field.Name
			if field.Types[0] == "InputFile" || field.Types[0] == "InputMedia" {
				field.Optional = false
			}
		CheckGeneric:
			isFieldRawImport := utils.IsRawField(field.Types)
			if isFieldRawImport && isMethod {
				builder.AddImport("rawTypes", fmt.Sprintf("%s/types/raw", consts.PackageName))
			} else if !utils.IsSimpleGeneric(field.Types) && isMethod {
				builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
			}
			genericName := utils.FixGeneric(
				field.Optional,
				field.Name,
				field.Types,
				isMethod || isFieldRawImport,
				isFieldRawImport,
			)
			if len(field.Default) > 0 {
				if isMethod {
					panic("Default value for methods is not supported")
				}
				continue
			}
			if field.Optional || field.Types[0] == "InputFile" || strings.Contains(field.Types[0], "Array of") {
				jsonName += ",omitempty"
			}
			if field.Types[0] == "InputFile" {
				filesInput = append(filesInput, field)
			} else if field.Name == "media" {
				if genericName == "string" {
					field = types.FieldTL{
						Name: field.Name,
						Types: []string{
							"InputFile",
						},
						Optional: field.Optional,
					}
					goto CheckGeneric
				}
				filesInput = append(filesInput, field)
			}
			builder.AddField(
				field.Name,
				genericName,
				jsonName,
			)
		}
		builder.CloseBracket()
	}
	return filesInput
}
