// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongdoc/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Button_sql sql.NullBool
var dummy_Button_time time.Duration
var dummy_Button_sort sort.Float64Slice

// ButtonAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model buttonAPI
type ButtonAPI struct {
	gorm.Model

	models.Button

	// encoding of pointers
	ButtonPointersEnconding
}

// ButtonPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ButtonPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field Node{}.Buttons []*Button
	Node_ButtonsDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Node_ButtonsDBID_Index sql.NullInt64
}

// ButtonDB describes a button in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model buttonDB
type ButtonDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field buttonDB.Name
	Name_Data sql.NullString

	// Declation for basic field buttonDB.Icon
	Icon_Data sql.NullString
	// encoding of pointers
	ButtonPointersEnconding
}

// ButtonDBs arrays buttonDBs
// swagger:response buttonDBsResponse
type ButtonDBs []ButtonDB

// ButtonDBResponse provides response
// swagger:response buttonDBResponse
type ButtonDBResponse struct {
	ButtonDB
}

// ButtonWOP is a Button without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ButtonWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Icon string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var Button_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Icon",
}

type BackRepoButtonStruct struct {
	// stores ButtonDB according to their gorm ID
	Map_ButtonDBID_ButtonDB map[uint]*ButtonDB

	// stores ButtonDB ID according to Button address
	Map_ButtonPtr_ButtonDBID map[*models.Button]uint

	// stores Button according to their gorm ID
	Map_ButtonDBID_ButtonPtr map[uint]*models.Button

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoButton *BackRepoButtonStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoButton.stage
	return
}

func (backRepoButton *BackRepoButtonStruct) GetDB() *gorm.DB {
	return backRepoButton.db
}

// GetButtonDBFromButtonPtr is a handy function to access the back repo instance from the stage instance
func (backRepoButton *BackRepoButtonStruct) GetButtonDBFromButtonPtr(button *models.Button) (buttonDB *ButtonDB) {
	id := backRepoButton.Map_ButtonPtr_ButtonDBID[button]
	buttonDB = backRepoButton.Map_ButtonDBID_ButtonDB[id]
	return
}

// BackRepoButton.CommitPhaseOne commits all staged instances of Button to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoButton *BackRepoButtonStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for button := range stage.Buttons {
		backRepoButton.CommitPhaseOneInstance(button)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, button := range backRepoButton.Map_ButtonDBID_ButtonPtr {
		if _, ok := stage.Buttons[button]; !ok {
			backRepoButton.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoButton.CommitDeleteInstance commits deletion of Button to the BackRepo
func (backRepoButton *BackRepoButtonStruct) CommitDeleteInstance(id uint) (Error error) {

	button := backRepoButton.Map_ButtonDBID_ButtonPtr[id]

	// button is not staged anymore, remove buttonDB
	buttonDB := backRepoButton.Map_ButtonDBID_ButtonDB[id]
	query := backRepoButton.db.Unscoped().Delete(&buttonDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoButton.Map_ButtonPtr_ButtonDBID, button)
	delete(backRepoButton.Map_ButtonDBID_ButtonPtr, id)
	delete(backRepoButton.Map_ButtonDBID_ButtonDB, id)

	return
}

// BackRepoButton.CommitPhaseOneInstance commits button staged instances of Button to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoButton *BackRepoButtonStruct) CommitPhaseOneInstance(button *models.Button) (Error error) {

	// check if the button is not commited yet
	if _, ok := backRepoButton.Map_ButtonPtr_ButtonDBID[button]; ok {
		return
	}

	// initiate button
	var buttonDB ButtonDB
	buttonDB.CopyBasicFieldsFromButton(button)

	query := backRepoButton.db.Create(&buttonDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoButton.Map_ButtonPtr_ButtonDBID[button] = buttonDB.ID
	backRepoButton.Map_ButtonDBID_ButtonPtr[buttonDB.ID] = button
	backRepoButton.Map_ButtonDBID_ButtonDB[buttonDB.ID] = &buttonDB

	return
}

// BackRepoButton.CommitPhaseTwo commits all staged instances of Button to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoButton *BackRepoButtonStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, button := range backRepoButton.Map_ButtonDBID_ButtonPtr {
		backRepoButton.CommitPhaseTwoInstance(backRepo, idx, button)
	}

	return
}

// BackRepoButton.CommitPhaseTwoInstance commits {{structname }} of models.Button to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoButton *BackRepoButtonStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, button *models.Button) (Error error) {

	// fetch matching buttonDB
	if buttonDB, ok := backRepoButton.Map_ButtonDBID_ButtonDB[idx]; ok {

		buttonDB.CopyBasicFieldsFromButton(button)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoButton.db.Save(&buttonDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Button intance %s", button.Name))
		return err
	}

	return
}

// BackRepoButton.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoButton *BackRepoButtonStruct) CheckoutPhaseOne() (Error error) {

	buttonDBArray := make([]ButtonDB, 0)
	query := backRepoButton.db.Find(&buttonDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	buttonInstancesToBeRemovedFromTheStage := make(map[*models.Button]any)
	for key, value := range backRepoButton.stage.Buttons {
		buttonInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, buttonDB := range buttonDBArray {
		backRepoButton.CheckoutPhaseOneInstance(&buttonDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		button, ok := backRepoButton.Map_ButtonDBID_ButtonPtr[buttonDB.ID]
		if ok {
			delete(buttonInstancesToBeRemovedFromTheStage, button)
		}
	}

	// remove from stage and back repo's 3 maps all buttons that are not in the checkout
	for button := range buttonInstancesToBeRemovedFromTheStage {
		button.Unstage(backRepoButton.GetStage())

		// remove instance from the back repo 3 maps
		buttonID := backRepoButton.Map_ButtonPtr_ButtonDBID[button]
		delete(backRepoButton.Map_ButtonPtr_ButtonDBID, button)
		delete(backRepoButton.Map_ButtonDBID_ButtonDB, buttonID)
		delete(backRepoButton.Map_ButtonDBID_ButtonPtr, buttonID)
	}

	return
}

// CheckoutPhaseOneInstance takes a buttonDB that has been found in the DB, updates the backRepo and stages the
// models version of the buttonDB
func (backRepoButton *BackRepoButtonStruct) CheckoutPhaseOneInstance(buttonDB *ButtonDB) (Error error) {

	button, ok := backRepoButton.Map_ButtonDBID_ButtonPtr[buttonDB.ID]
	if !ok {
		button = new(models.Button)

		backRepoButton.Map_ButtonDBID_ButtonPtr[buttonDB.ID] = button
		backRepoButton.Map_ButtonPtr_ButtonDBID[button] = buttonDB.ID

		// append model store with the new element
		button.Name = buttonDB.Name_Data.String
		button.Stage(backRepoButton.GetStage())
	}
	buttonDB.CopyBasicFieldsToButton(button)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	button.Stage(backRepoButton.GetStage())

	// preserve pointer to buttonDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ButtonDBID_ButtonDB)[buttonDB hold variable pointers
	buttonDB_Data := *buttonDB
	preservedPtrToButton := &buttonDB_Data
	backRepoButton.Map_ButtonDBID_ButtonDB[buttonDB.ID] = preservedPtrToButton

	return
}

// BackRepoButton.CheckoutPhaseTwo Checkouts all staged instances of Button to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoButton *BackRepoButtonStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, buttonDB := range backRepoButton.Map_ButtonDBID_ButtonDB {
		backRepoButton.CheckoutPhaseTwoInstance(backRepo, buttonDB)
	}
	return
}

// BackRepoButton.CheckoutPhaseTwoInstance Checkouts staged instances of Button to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoButton *BackRepoButtonStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, buttonDB *ButtonDB) (Error error) {

	button := backRepoButton.Map_ButtonDBID_ButtonPtr[buttonDB.ID]
	_ = button // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitButton allows commit of a single button (if already staged)
func (backRepo *BackRepoStruct) CommitButton(button *models.Button) {
	backRepo.BackRepoButton.CommitPhaseOneInstance(button)
	if id, ok := backRepo.BackRepoButton.Map_ButtonPtr_ButtonDBID[button]; ok {
		backRepo.BackRepoButton.CommitPhaseTwoInstance(backRepo, id, button)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitButton allows checkout of a single button (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutButton(button *models.Button) {
	// check if the button is staged
	if _, ok := backRepo.BackRepoButton.Map_ButtonPtr_ButtonDBID[button]; ok {

		if id, ok := backRepo.BackRepoButton.Map_ButtonPtr_ButtonDBID[button]; ok {
			var buttonDB ButtonDB
			buttonDB.ID = id

			if err := backRepo.BackRepoButton.db.First(&buttonDB, id).Error; err != nil {
				log.Panicln("CheckoutButton : Problem with getting object with id:", id)
			}
			backRepo.BackRepoButton.CheckoutPhaseOneInstance(&buttonDB)
			backRepo.BackRepoButton.CheckoutPhaseTwoInstance(backRepo, &buttonDB)
		}
	}
}

// CopyBasicFieldsFromButton
func (buttonDB *ButtonDB) CopyBasicFieldsFromButton(button *models.Button) {
	// insertion point for fields commit

	buttonDB.Name_Data.String = button.Name
	buttonDB.Name_Data.Valid = true

	buttonDB.Icon_Data.String = button.Icon
	buttonDB.Icon_Data.Valid = true
}

// CopyBasicFieldsFromButtonWOP
func (buttonDB *ButtonDB) CopyBasicFieldsFromButtonWOP(button *ButtonWOP) {
	// insertion point for fields commit

	buttonDB.Name_Data.String = button.Name
	buttonDB.Name_Data.Valid = true

	buttonDB.Icon_Data.String = button.Icon
	buttonDB.Icon_Data.Valid = true
}

// CopyBasicFieldsToButton
func (buttonDB *ButtonDB) CopyBasicFieldsToButton(button *models.Button) {
	// insertion point for checkout of basic fields (back repo to stage)
	button.Name = buttonDB.Name_Data.String
	button.Icon = buttonDB.Icon_Data.String
}

// CopyBasicFieldsToButtonWOP
func (buttonDB *ButtonDB) CopyBasicFieldsToButtonWOP(button *ButtonWOP) {
	button.ID = int(buttonDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	button.Name = buttonDB.Name_Data.String
	button.Icon = buttonDB.Icon_Data.String
}

// Backup generates a json file from a slice of all ButtonDB instances in the backrepo
func (backRepoButton *BackRepoButtonStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ButtonDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ButtonDB, 0)
	for _, buttonDB := range backRepoButton.Map_ButtonDBID_ButtonDB {
		forBackup = append(forBackup, buttonDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Button ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Button file", err.Error())
	}
}

// Backup generates a json file from a slice of all ButtonDB instances in the backrepo
func (backRepoButton *BackRepoButtonStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ButtonDB, 0)
	for _, buttonDB := range backRepoButton.Map_ButtonDBID_ButtonDB {
		forBackup = append(forBackup, buttonDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Button")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Button_Fields, -1)
	for _, buttonDB := range forBackup {

		var buttonWOP ButtonWOP
		buttonDB.CopyBasicFieldsToButtonWOP(&buttonWOP)

		row := sh.AddRow()
		row.WriteStruct(&buttonWOP, -1)
	}
}

// RestoreXL from the "Button" sheet all ButtonDB instances
func (backRepoButton *BackRepoButtonStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoButtonid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Button"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoButton.rowVisitorButton)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoButton *BackRepoButtonStruct) rowVisitorButton(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var buttonWOP ButtonWOP
		row.ReadStruct(&buttonWOP)

		// add the unmarshalled struct to the stage
		buttonDB := new(ButtonDB)
		buttonDB.CopyBasicFieldsFromButtonWOP(&buttonWOP)

		buttonDB_ID_atBackupTime := buttonDB.ID
		buttonDB.ID = 0
		query := backRepoButton.db.Create(buttonDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoButton.Map_ButtonDBID_ButtonDB[buttonDB.ID] = buttonDB
		BackRepoButtonid_atBckpTime_newID[buttonDB_ID_atBackupTime] = buttonDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ButtonDB.json" in dirPath that stores an array
// of ButtonDB and stores it in the database
// the map BackRepoButtonid_atBckpTime_newID is updated accordingly
func (backRepoButton *BackRepoButtonStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoButtonid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ButtonDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Button file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ButtonDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ButtonDBID_ButtonDB
	for _, buttonDB := range forRestore {

		buttonDB_ID_atBackupTime := buttonDB.ID
		buttonDB.ID = 0
		query := backRepoButton.db.Create(buttonDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoButton.Map_ButtonDBID_ButtonDB[buttonDB.ID] = buttonDB
		BackRepoButtonid_atBckpTime_newID[buttonDB_ID_atBackupTime] = buttonDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Button file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Button>id_atBckpTime_newID
// to compute new index
func (backRepoButton *BackRepoButtonStruct) RestorePhaseTwo() {

	for _, buttonDB := range backRepoButton.Map_ButtonDBID_ButtonDB {

		// next line of code is to avert unused variable compilation error
		_ = buttonDB

		// insertion point for reindexing pointers encoding
		// This reindex button.Buttons
		if buttonDB.Node_ButtonsDBID.Int64 != 0 {
			buttonDB.Node_ButtonsDBID.Int64 =
				int64(BackRepoNodeid_atBckpTime_newID[uint(buttonDB.Node_ButtonsDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoButton.db.Model(buttonDB).Updates(*buttonDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoButtonid_atBckpTime_newID map[uint]uint