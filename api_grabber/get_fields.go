package api_grabber

import (
	"BotApiCompiler/api_grabber/types"
	"github.com/anaskhan96/soup"
	"regexp"
	"strconv"
	"strings"
)

func (ctx *Context) GetFields(currName string, isMethod bool, x soup.Root) {
	body := x.Find("tbody")
	var fields []types.FieldTL
	rgx1, _ := regexp.Compile("must be (.*)")
	rgx2, _ := regexp.Compile("“(.*?)”")
	rgx3, _ := regexp.Compile(", always “(.*?)”")
	rgx4, _ := regexp.Compile("Type of the result, must be (.*)")
	for _, tr := range body.FindAll("tr") {
		children := tr.FindAll("td")
		description := children[2].FullText()
		defaultValue := ""
		if !isMethod && len(children) == 3 {
			fieldName := children[0].FullText()
			if !isMethod {
				var typeIdsReturn []string
				if strings.Contains(description, "Type of the result, must be ") {
					defaultValue = rgx4.FindStringSubmatch(description)[1]
				} else if strings.Contains(description, ", must be ") {
					typeIdentifier := rgx1.FindStringSubmatch(description)[1]
					typeIdsReturn = append(typeIdsReturn, typeIdentifier)
				} else if strings.Contains(description, ", one of “") {
					typeIdentifier := rgx2.FindAllStringSubmatch(description, -1)
					var y []string
					for _, t := range typeIdentifier {
						y = append(y, t[1])
					}
					typeIdsReturn = append(typeIdsReturn, y...)
				} else if strings.Contains(description, ", always ") {
					typeIdentifier := rgx3.FindStringSubmatch(description)[1]
					typeIdsReturn = append(typeIdsReturn, typeIdentifier)
				}
				if len(typeIdsReturn) > 0 {
					ctx.ApiTL.Types[currName].TypeIds = &types.TypeIdsDescriptor{
						CommonName: fieldName,
						TypeIds:    typeIdsReturn,
					}
				}
			}
			fields = append(fields, types.FieldTL{
				Name:     fieldName,
				Types:    CleanTgType(children[1].FullText()),
				Optional: strings.HasPrefix(description, "Optional."),
				Default:  defaultValue,
			})
		} else if isMethod && len(children) == 4 {
			fields = append(fields, types.FieldTL{
				Name:     children[0].FullText(),
				Types:    CleanTgType(children[1].FullText()),
				Optional: description != "Yes",
				Default:  defaultValue,
			})
		} else {
			panic("Invalid table " + strconv.Itoa(len(children)) + " " + currName)
		}
	}
	if isMethod {
		ctx.ApiTL.Methods[currName].Params = fields
	} else {
		ctx.ApiTL.Types[currName].Fields = fields
	}
}
