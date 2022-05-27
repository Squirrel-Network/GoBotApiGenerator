package sub_build

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/interfaces"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"fmt"
	"strings"
)

func BuildMarshaller[Scheme interfaces.SchemeInterface](typeScheme Scheme, builder *component.Context, listElements map[string]*types.ApiTypeTL) {
	sendChildTypes := make(map[string]types.FieldTL)
	isMethod := typeScheme.GetType() == "methods"
	foundDefaults := false
	for _, field := range typeScheme.GetFields() {
		genericName := strings.ReplaceAll(utils.FixGeneric(false, field.Name, field.Types, false, false), "[]", "")
		if listElements[genericName] != nil {
			genericCheck := listElements[genericName]
			if len(genericCheck.GetSubTypes()) > 0 && genericCheck.IsSendMethod() {
				sendChildTypes[field.Name] = types.FieldTL{
					Name:     field.Name,
					Types:    genericCheck.GetSubTypes(),
					Optional: field.Optional,
					Default:  field.Default,
				}
			} else if genericCheck.IsSendMethod() && len(genericCheck.GetFields()) == 0 && genericName != "InputFile" {
				sendChildTypes[field.Name] = field
			}
		} else if genericName == "interface{}" {
			sendChildTypes[field.Name] = field
		}
		if len(field.Default) > 0 {
			foundDefaults = true
		}
	}
	if foundDefaults {
		builder.AddLine()
		parentStructName := utils.FixStructName(typeScheme.GetName())
		builder.AddImport("", "encoding/json")
		builder.AddFunc(
			fmt.Sprintf("entity %s", parentStructName),
			"MarshalJSON",
			nil,
			"([]byte, error)",
		)
		builder.InitInlineStruct("alias")
		for _, field := range typeScheme.GetFields() {
		GenerateGeneric:
			isFieldRawImport := utils.IsRawField(field.Types)
			if field.Types[0] == "InputFile" {
				field.Optional = false
			}
			genericName := utils.FixGeneric(
				field.Optional,
				field.Name,
				field.Types,
				isMethod || isFieldRawImport,
				isFieldRawImport,
			)
			jsonName := field.Name
			if field.Optional || field.Types[0] == "InputFile" {
				jsonName += ",omitempty"
			}
			if field.Name == "media" && genericName == "string" {
				field = types.FieldTL{
					Name: field.Name,
					Types: []string{
						"InputFile",
					},
					Optional: field.Optional,
				}
				goto GenerateGeneric
			}
			builder.AddField(
				field.Name,
				genericName,
				jsonName,
			)
		}
		builder.CloseOpenBracket()
		for _, field := range typeScheme.GetFields() {
			prettifiedField := utils.PrettifyField(field.Name)
			var value string
			if len(field.Default) > 0 {
				value = fmt.Sprintf("\"%s\"", field.Default)
			} else {
				value = fmt.Sprintf("entity.%s", prettifiedField)
			}
			builder.FillField(
				prettifiedField,
				value,
			)
		}
		builder.CloseBracket()
		if len(sendChildTypes) > 0 {
			BuildCheck(builder, isMethod, sendChildTypes)
		}
		builder.AddReturn("json.Marshal(alias)").AddLine()
		builder.CloseBracket()
	} else if len(sendChildTypes) > 0 {
		builder.AddLine()
		parentStructName := utils.FixStructName(typeScheme.GetName())
		builder.AddImport("", "encoding/json")
		builder.AddFunc(
			fmt.Sprintf("entity %s", parentStructName),
			"MarshalJSON",
			nil,
			"([]byte, error)",
		)
		BuildCheck(builder, isMethod, sendChildTypes)
		builder.AddType("x0", parentStructName).AddLine()
		builder.AddReturn("json.Marshal((x0)(entity))").AddLine()
		builder.CloseBracket()
	}
}
