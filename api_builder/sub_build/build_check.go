package sub_build

import (
	"BotApiCompiler/api_builder/component"
	"BotApiCompiler/api_builder/utils"
	"BotApiCompiler/api_grabber/types"
	"BotApiCompiler/consts"
	"fmt"
	"strings"
)

func BuildCheck(builder *component.Context, isMethod bool, sendChildTypes map[string]types.FieldTL) {
	builder.AddImport("", "fmt")
	for fieldName, fieldTypes := range sendChildTypes {
		genericNameTmp := utils.FixGeneric(false, fieldTypes.Name, fieldTypes.Types, false, false)
		arrayCounts := strings.Count(genericNameTmp, "[]")
		entityName := fmt.Sprintf("entity.%s", utils.PrettifyField(fieldName))
		for i := 0; i < arrayCounts; i++ {
			builder.AddFor(fmt.Sprintf("_, x%d := range %s", i, entityName))
			entityName = fmt.Sprintf("x%d", i)
		}
		originalEntityName := entityName
		if fieldTypes.Optional {
			builder.AddIf(fmt.Sprintf("%s != nil", entityName))
			if !strings.Contains(genericNameTmp, "interface") {
				entityName = fmt.Sprintf("(*%s)", entityName)
			}
		}
		builder.InitSwitch(fmt.Sprintf("%s.(type)", entityName))
		var fixedCases []string
		for _, field := range fieldTypes.Types {
			genericName := strings.ReplaceAll(utils.FixGeneric(false, "", []string{field}, false, false), "[]", "")
			if isMethod {
				builder.AddImport("", fmt.Sprintf("%s/types", consts.PackageName))
				fixedCases = append(fixedCases, fmt.Sprintf("*types.%s", genericName))
			} else {
				fixedCases = append(fixedCases, genericName)
			}
		}
		builder.AddCase(false, fixedCases)
		builder.AddBreak().CloseCase()
		builder.AddDefault()
		builder.AddReturn(fmt.Sprintf("nil, fmt.Errorf(\"%s: unknown type: %%T\", %s)", fieldName, originalEntityName))
		builder.CloseCase().AddLine()
		for i := 0; i < arrayCounts; i++ {
			builder.CloseBracket()
		}
		if fieldTypes.Optional {
			builder.CloseBracket()
		}
		builder.CloseBracket()
	}
}
