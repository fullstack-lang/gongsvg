// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Rect__dummysDeclaration__ models.Rect
var __Rect_time__dummyDeclaration time.Duration

// An RectID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRect updateRect deleteRect
type RectID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RectInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRect updateRect
type RectInput struct {
	// The Rect to submit or modify
	// in: body
	Rect *orm.RectAPI
}

// GetRects
//
// swagger:route GET /rects rects getRects
//
// Get all rects
//
// Responses:
//    default: genericError
//        200: rectDBsResponse
func GetRects(c *gin.Context) {
	db := orm.BackRepo.BackRepoRect.GetDB()

	// source slice
	var rectDBs []orm.RectDB
	query := db.Find(&rectDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	rectAPIs := make([]orm.RectAPI, 0)

	// for each rect, update fields from the database nullable fields
	for idx := range rectDBs {
		rectDB := &rectDBs[idx]
		_ = rectDB
		var rectAPI orm.RectAPI

		// insertion point for updating fields
		rectAPI.ID = rectDB.ID
		rectDB.CopyBasicFieldsToRect(&rectAPI.Rect)
		rectAPI.RectPointersEnconding = rectDB.RectPointersEnconding
		rectAPIs = append(rectAPIs, rectAPI)
	}

	c.JSON(http.StatusOK, rectAPIs)
}

// PostRect
//
// swagger:route POST /rects rects postRect
//
// Creates a rect
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: rectDBResponse
func PostRect(c *gin.Context) {
	db := orm.BackRepo.BackRepoRect.GetDB()

	// Validate input
	var input orm.RectAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create rect
	rectDB := orm.RectDB{}
	rectDB.RectPointersEnconding = input.RectPointersEnconding
	rectDB.CopyBasicFieldsFromRect(&input.Rect)

	query := db.Create(&rectDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, rectDB)
}

// GetRect
//
// swagger:route GET /rects/{ID} rects getRect
//
// Gets the details for a rect.
//
// Responses:
//    default: genericError
//        200: rectDBResponse
func GetRect(c *gin.Context) {
	db := orm.BackRepo.BackRepoRect.GetDB()

	// Get rectDB in DB
	var rectDB orm.RectDB
	if err := db.First(&rectDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var rectAPI orm.RectAPI
	rectAPI.ID = rectDB.ID
	rectAPI.RectPointersEnconding = rectDB.RectPointersEnconding
	rectDB.CopyBasicFieldsToRect(&rectAPI.Rect)

	c.JSON(http.StatusOK, rectAPI)
}

// UpdateRect
//
// swagger:route PATCH /rects/{ID} rects updateRect
//
// Update a rect
//
// Responses:
//    default: genericError
//        200: rectDBResponse
func UpdateRect(c *gin.Context) {
	db := orm.BackRepo.BackRepoRect.GetDB()

	// Get model if exist
	var rectDB orm.RectDB

	// fetch the rect
	query := db.First(&rectDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.RectAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	rectDB.CopyBasicFieldsFromRect(&input.Rect)
	rectDB.RectPointersEnconding = input.RectPointersEnconding

	query = db.Model(&rectDB).Updates(rectDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rectDB
	c.JSON(http.StatusOK, rectDB)
}

// DeleteRect
//
// swagger:route DELETE /rects/{ID} rects deleteRect
//
// Delete a rect
//
// Responses:
//    default: genericError
func DeleteRect(c *gin.Context) {
	db := orm.BackRepo.BackRepoRect.GetDB()

	// Get model if exist
	var rectDB orm.RectDB
	if err := db.First(&rectDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&rectDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
