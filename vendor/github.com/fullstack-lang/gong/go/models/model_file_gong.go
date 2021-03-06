package models

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const ModelGongFileTemplate = `// generated by ModelGongFileTemplate
package models

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
)

// swagger:ignore
type __void struct{}

// needed for creating set of instances in the stage
var __member __void

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct { // insertion point for definition of arrays registering instances{{` + string(rune(ModelGongStructArrayDefintion)) + `}}
	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures{{` + string(rune(ModelGongInsertionCommitCheckoutSignature)) + `}}
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

// swagger:ignore instructs the gong compiler (gongc) to avoid this particular struct
var Stage StageStruct = StageStruct{ // insertion point for array initiatialisation{{` + string(rune(ModelGongInsertionArrayInitialisation)) + `}}
	// end of insertion point
	Map_GongStructName_InstancesNb: make(map[string]int),
}

func (stage *StageStruct) Commit() {
	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct{{` + string(rune(ModelGongInsertionComputeNbInstances)) + `}}

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls{{` + string(rune(ModelGongInsertionStageFunctions)) + `}}
// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation{{` + string(rune(ModelGongInsertionCreateCallback)) + `}}
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion{{` + string(rune(ModelGongInsertionDeleteCallback)) + `}}
}

func (stage *StageStruct) Reset() { // insertion point for array reset{{` + string(rune(ModelGongInsertionArrayReset)) + `}}
}

func (stage *StageStruct) Nil() { // insertion point for array nil{{` + string(rune(ModelGongInsertionArrayNil)) + `}}
}

const marshallRes = ` + "`" + `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"
)

func init() {
	var __Dummy_time_variable time.Time
	_ = __Dummy_time_variable
	InjectionGateway["{{databaseName}}"] = {{databaseName}}Injection
}

// {{databaseName}}Injection will stage objects of database "{{databaseName}}"
func {{databaseName}}Injection() {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}

` + "`" + `

const IdentifiersDecls = ` + "`" + `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: "{{GeneratedFieldNameValue}}"}).Stage()` + "`" + `

const StringInitStatement = ` + "`" + `
	{{Identifier}}.{{GeneratedFieldName}} = "{{GeneratedFieldNameValue}}"` + "`" + `

const NumberInitStatement = ` + "`" + `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}` + "`" + `

const PointerFieldInitStatement = ` + "`" + `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}` + "`" + `

const SliceOfPointersFieldInitStatement = ` + "`" + `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})` + "`" + `

const TimeInitStatement = ` + "`" + `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")` + "`" + `

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output  is " + name)

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(path.Base(name), ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	decl := ""
	setValueField := ""

	// insertion initialization of objects to stage{{` + string(rune(ModelGongInsertionUnmarshallDeclarations)) + `}}

	// insertion initialization of objects to stage{{` + string(rune(ModelGongInsertionUnmarshallPointersInitializations)) + `}}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	fmt.Fprintln(file, res)
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
`

// insertion points
type ModelGongInsertionPoints int

const (
	ModelGongInsertionCommitCheckoutSignature ModelGongInsertionPoints = iota
	ModelGongInsertionStageFunctions
	ModelGongInsertionCreateCallback
	ModelGongInsertionDeleteCallback
	ModelGongInsertionArrayDefintion
	ModelGongInsertionArrayInitialisation
	ModelGongInsertionArrayReset
	ModelGongInsertionArrayNil
	ModelGongInsertionUnmarshallDeclarations
	ModelGongInsertionUnmarshallPointersInitializations
	ModelGongInsertionComputeNbInstances
	ModelGongInsertionsNb
)

type ModelGongSubTemplate int

const (
	ModelGongCommitCheckout ModelGongSubTemplate = iota
	ModelGongStageFunction
	ModelGongStructCreateCallback
	ModelGongStructDeleteCallback
	ModelGongStructArrayDefintion
	ModelGongStructArrayInitialisation
	ModelGongStructArrayReset
	ModelGongStructArrayNil
	ModelGongStructUnmarshallStatementsStepValuesInit
	ModelGongStructUnmarshallStatementsStepPointersInit
	ModelGongStructComputeNbInstances
)

var ModelGongSubTemplateCode map[ModelGongSubTemplate]string = // new line
map[ModelGongSubTemplate]string{
	ModelGongCommitCheckout: `
	Commit{{Structname}}({{structname}} *{{Structname}})
	Checkout{{Structname}}({{structname}} *{{Structname}})`,

	ModelGongStageFunction: `
func (stage *StageStruct) get{{Structname}}OrderedStructWithNameField() []*{{Structname}} {
	// have alphabetical order generation
	{{structname}}Ordered := []*{{Structname}}{}
	for {{structname}} := range stage.{{Structname}}s {
		{{structname}}Ordered = append({{structname}}Ordered, {{structname}})
	}
	sort.Slice({{structname}}Ordered[:], func(i, j int) bool {
		return {{structname}}Ordered[i].Name < {{structname}}Ordered[j].Name
	})
	return {{structname}}Ordered
}

// Stage puts {{structname}} to the model stage
func ({{structname}} *{{Structname}}) Stage() *{{Structname}} {
	Stage.{{Structname}}s[{{structname}}] = __member
	Stage.{{Structname}}s_mapString[{{structname}}.Name] = {{structname}}

	return {{structname}}
}

// Unstage removes {{structname}} off the model stage
func ({{structname}} *{{Structname}}) Unstage() *{{Structname}} {
	delete(Stage.{{Structname}}s, {{structname}})
	delete(Stage.{{Structname}}s_mapString, {{structname}}.Name)
	return {{structname}}
}

// commit {{structname}} to the back repo (if it is already staged)
func ({{structname}} *{{Structname}}) Commit() *{{Structname}} {
	if _, ok := Stage.{{Structname}}s[{{structname}}]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.Commit{{Structname}}({{structname}})
		}
	}
	return {{structname}}
}

// Checkout {{structname}} to the back repo (if it is already staged)
func ({{structname}} *{{Structname}}) Checkout() *{{Structname}} {
	if _, ok := Stage.{{Structname}}s[{{structname}}]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.Checkout{{Structname}}({{structname}})
		}
	}
	return {{structname}}
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of {{structname}} to the model stage
func ({{structname}} *{{Structname}}) StageCopy() *{{Structname}} {
	_{{structname}} := new({{Structname}})
	*_{{structname}} = *{{structname}}
	_{{structname}}.Stage()
	return _{{structname}}
}

// StageAndCommit appends {{structname}} to the model stage and commit to the orm repo
func ({{structname}} *{{Structname}}) StageAndCommit() *{{Structname}} {
	{{structname}}.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORM{{Structname}}({{structname}})
	}
	return {{structname}}
}

// DeleteStageAndCommit appends {{structname}} to the model stage and commit to the orm repo
func ({{structname}} *{{Structname}}) DeleteStageAndCommit() *{{Structname}} {
	{{structname}}.Unstage()
	DeleteORM{{Structname}}({{structname}})
	return {{structname}}
}

// StageCopyAndCommit appends a copy of {{structname}} to the model stage and commit to the orm repo
func ({{structname}} *{{Structname}}) StageCopyAndCommit() *{{Structname}} {
	_{{structname}} := new({{Structname}})
	*_{{structname}} = *{{structname}}
	_{{structname}}.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORM{{Structname}}({{structname}})
	}
	return _{{structname}}
}

// CreateORM{{Structname}} enables dynamic staging of a {{Structname}} instance
func CreateORM{{Structname}}({{structname}} *{{Structname}}) {
	{{structname}}.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORM{{Structname}}({{structname}})
	}
}

// DeleteORM{{Structname}} enables dynamic staging of a {{Structname}} instance
func DeleteORM{{Structname}}({{structname}} *{{Structname}}) {
	{{structname}}.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORM{{Structname}}({{structname}})
	}
}
`,

	ModelGongStructCreateCallback: `
	CreateORM{{Structname}}({{Structname}} *{{Structname}})`,

	ModelGongStructDeleteCallback: `
	DeleteORM{{Structname}}({{Structname}} *{{Structname}})`,

	ModelGongStructArrayDefintion: `
	{{Structname}}s           map[*{{Structname}}]struct{}
	{{Structname}}s_mapString map[string]*{{Structname}}
`,

	ModelGongStructArrayInitialisation: `
	{{Structname}}s:           make(map[*{{Structname}}]struct{}),
	{{Structname}}s_mapString: make(map[string]*{{Structname}}),
`,

	ModelGongStructArrayReset: `
	stage.{{Structname}}s = make(map[*{{Structname}}]struct{})
	stage.{{Structname}}s_mapString = make(map[string]*{{Structname}})
`,

	ModelGongStructArrayNil: `
	stage.{{Structname}}s = nil
	stage.{{Structname}}s_mapString = nil
`,

	ModelGongStructUnmarshallStatementsStepValuesInit: `
	map_{{Structname}}_Identifiers := make(map[*{{Structname}}]string)
	_ = map_{{Structname}}_Identifiers

	{{structname}}Ordered := []*{{Structname}}{}
	for {{structname}} := range stage.{{Structname}}s {
		{{structname}}Ordered = append({{structname}}Ordered, {{structname}})
	}
	sort.Slice({{structname}}Ordered[:], func(i, j int) bool {
		return {{structname}}Ordered[i].Name < {{structname}}Ordered[j].Name
	})
	identifiersDecl += fmt.Sprintf("\n\n	// Declarations of staged instances of {{Structname}}")
	for idx, {{structname}} := range {{structname}}Ordered {

		id = generatesIdentifier("{{Structname}}", idx, {{structname}}.Name)
		map_{{Structname}}_Identifiers[{{structname}}] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "{{Structname}}")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", {{structname}}.Name)
		identifiersDecl += decl

		initializerStatements += fmt.Sprintf("\n\n	// {{Structname}} %s values setup", {{structname}}.Name)
		// Initialisation of values{{ValuesInitialization}}
	}
`,

	ModelGongStructUnmarshallStatementsStepPointersInit: `
	for idx, {{structname}} := range {{structname}}Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("{{Structname}}", idx, {{structname}}.Name)
		map_{{Structname}}_Identifiers[{{structname}}] = id

		// Initialisation of values{{PointersInitialization}}
	}
`,

	ModelGongStructComputeNbInstances: `
	stage.Map_GongStructName_InstancesNb["{{Structname}}"] = len(stage.{{Structname}}s)`,
}

var ModelGongSubSubTemplateCode map[string]string = // new line
map[string]string{}

var ModelGongSubSubToSubMap map[string]string = //
map[string]string{}

//
// Sub Templates
//
type GongFilePerStructSubTemplate int

const (
	GongFileFieldSubTmplSetBasicFieldBool GongFilePerStructSubTemplate = iota
	GongFileFieldSubTmplSetBasicFieldInt
	GongFileFieldSubTmplSetBasicFieldFloat64
	GongFileFieldSubTmplSetBasicFieldString
	GongFileFieldSubTmplSetTimeField
	GongFileFieldSubTmplSetBasicFieldStringEnum
	GongFileFieldSubTmplSetPointerField
	GongFileFieldSubTmplSetSliceOfPointersField
)

var GongFileFieldFieldSubTemplateCode map[GongFilePerStructSubTemplate]string = // declaration of the sub templates
map[GongFilePerStructSubTemplate]string{

	GongFileFieldSubTmplSetBasicFieldBool: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", {{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	GongFileFieldSubTmplSetTimeField: `
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", {{structname}}.{{FieldName}}.String())
		initializerStatements += setValueField

`,
	GongFileFieldSubTmplSetBasicFieldInt: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", {{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	GongFileFieldSubTmplSetBasicFieldFloat64: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", {{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	GongFileFieldSubTmplSetBasicFieldString: `
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string({{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	GongFileFieldSubTmplSetPointerField: `
		if {{structname}}.{{FieldName}} != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "{{FieldName}}")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_{{AssocStructName}}_Identifiers[{{structname}}.{{FieldName}}])
			pointersInitializesStatements += setPointerField
		}
`,
	GongFileFieldSubTmplSetSliceOfPointersField: `
		for _, _{{assocstructname}} := range {{structname}}.{{FieldName}} {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "{{FieldName}}")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_{{AssocStructName}}_Identifiers[_{{assocstructname}}])
			pointersInitializesStatements += setPointerField
		}
`,
}

func CodeGeneratorModelGong(
	mdlPkg *ModelPkg,
	pkgName string,
	pkgPath string) {

	// generate the typescript file
	codeGO := ModelGongFileTemplate

	insertions := make(map[ModelGongInsertionPoints]string)
	for insertion := ModelGongInsertionPoints(0); insertion < ModelGongInsertionsNb; insertion++ {
		insertions[insertion] = ""
	}

	subCodes := make(map[ModelGongSubTemplate]string)
	for subTemplate := range ModelGongSubTemplateCode {
		subCodes[subTemplate] = ""
	}

	gongStructs := []*GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		gongStructs = append(gongStructs, _struct)
	}
	sort.Slice(gongStructs[:], func(i, j int) bool {
		return gongStructs[i].Name < gongStructs[j].Name
	})

	for _, gongStruct := range gongStructs {

		if !gongStruct.HasNameField() {
			continue
		}

		for subTemplate := range ModelGongSubTemplateCode {

			// replace {{ValuesInitialization}}
			valInitCode := ""
			pointerInitCode := ""
			for _, field := range gongStruct.Fields {
				switch field := field.(type) {
				case *GongBasicField:

					switch field.basicKind {
					case types.String:
						valInitCode += Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldString],
							"{{FieldName}}", field.Name)
					case types.Bool:
						valInitCode += Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldBool],
							"{{FieldName}}", field.Name)
					case types.Float64:
						valInitCode += Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldFloat64],
							"{{FieldName}}", field.Name)
					case types.Int, types.Int64:
						valInitCode += Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldInt],
							"{{FieldName}}", field.Name)
					default:
					}
				case *GongTimeField:
					valInitCode += Replace1(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetTimeField],
						"{{FieldName}}", field.Name)
				case *PointerToGongStructField:
					pointerInitCode += Replace2(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetPointerField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
				case *SliceOfPointerToGongStructField:
					pointerInitCode += Replace3(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetSliceOfPointersField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
				default:
				}
			}

			valInitCode = Replace2(valInitCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			pointerInitCode = Replace2(pointerInitCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := Replace4(ModelGongSubTemplateCode[subTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{ValuesInitialization}}", valInitCode,
				"{{PointersInitialization}}", pointerInitCode,
			)

			subCodes[subTemplate] += generatedCodeFromSubTemplate
		}

	}

	// substitutes {{<<insertion points>>}} stuff with generated code
	for insertion := ModelGongInsertionPoints(0); insertion < ModelGongInsertionsNb; insertion++ {

		// compute insertion
		insertions[insertion] = subCodes[ModelGongSubTemplate(insertion)]

		toReplace := "{{" + string(rune(insertion)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, insertions[insertion])
	}

	codeGO = Replace3(codeGO,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", strings.Title(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName))

	file, err := os.Create(filepath.Join(pkgPath, "gong.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
