// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { GongEnumShapeDB } from '../gongenumshape-db'
import { GongEnumShapeService } from '../gongenumshape.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { ClassdiagramDB } from '../classdiagram-db'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// GongEnumShapeDetailComponent is initilizaed from different routes
// GongEnumShapeDetailComponentState detail different cases 
enum GongEnumShapeDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_GongEnumShapes_SET,
}

@Component({
	selector: 'app-gongenumshape-detail',
	templateUrl: './gongenumshape-detail.component.html',
	styleUrls: ['./gongenumshape-detail.component.css'],
})
export class GongEnumShapeDetailComponent implements OnInit {

	// insertion point for declarations

	// the GongEnumShapeDB of interest
	gongenumshape: GongEnumShapeDB = new GongEnumShapeDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: GongEnumShapeDetailComponentState = GongEnumShapeDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private gongenumshapeService: GongEnumShapeService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {

		// compute state
		this.id = +this.route.snapshot.paramMap.get('id')!;
		this.originStruct = this.route.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.route.snapshot.paramMap.get('originStructFieldName')!;

		const association = this.route.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = GongEnumShapeDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = GongEnumShapeDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "GongEnumShapes":
						// console.log("GongEnumShape" + " is instanciated with back pointer to instance " + this.id + " Classdiagram association GongEnumShapes")
						this.state = GongEnumShapeDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_GongEnumShapes_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getGongEnumShape()

		// observable for changes in structs
		this.gongenumshapeService.GongEnumShapeServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGongEnumShape()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getGongEnumShape(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case GongEnumShapeDetailComponentState.CREATE_INSTANCE:
						this.gongenumshape = new (GongEnumShapeDB)
						break;
					case GongEnumShapeDetailComponentState.UPDATE_INSTANCE:
						let gongenumshape = frontRepo.GongEnumShapes.get(this.id)
						console.assert(gongenumshape != undefined, "missing gongenumshape with id:" + this.id)
						this.gongenumshape = gongenumshape!
						break;
					// insertion point for init of association field
					case GongEnumShapeDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Classdiagram_GongEnumShapes_SET:
						this.gongenumshape = new (GongEnumShapeDB)
						this.gongenumshape.Classdiagram_GongEnumShapes_reverse = frontRepo.Classdiagrams.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		if (this.gongenumshape.PositionID == undefined) {
			this.gongenumshape.PositionID = new NullInt64
		}
		if (this.gongenumshape.Position != undefined) {
			this.gongenumshape.PositionID.Int64 = this.gongenumshape.Position.ID
			this.gongenumshape.PositionID.Valid = true
		} else {
			this.gongenumshape.PositionID.Int64 = 0
			this.gongenumshape.PositionID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.gongenumshape.Classdiagram_GongEnumShapes_reverse != undefined) {
			if (this.gongenumshape.Classdiagram_GongEnumShapesDBID == undefined) {
				this.gongenumshape.Classdiagram_GongEnumShapesDBID = new NullInt64
			}
			this.gongenumshape.Classdiagram_GongEnumShapesDBID.Int64 = this.gongenumshape.Classdiagram_GongEnumShapes_reverse.ID
			this.gongenumshape.Classdiagram_GongEnumShapesDBID.Valid = true
			if (this.gongenumshape.Classdiagram_GongEnumShapesDBID_Index == undefined) {
				this.gongenumshape.Classdiagram_GongEnumShapesDBID_Index = new NullInt64
			}
			this.gongenumshape.Classdiagram_GongEnumShapesDBID_Index.Valid = true
			this.gongenumshape.Classdiagram_GongEnumShapes_reverse = new ClassdiagramDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case GongEnumShapeDetailComponentState.UPDATE_INSTANCE:
				this.gongenumshapeService.updateGongEnumShape(this.gongenumshape)
					.subscribe(gongenumshape => {
						this.gongenumshapeService.GongEnumShapeServiceChanged.next("update")
					});
				break;
			default:
				this.gongenumshapeService.postGongEnumShape(this.gongenumshape).subscribe(gongenumshape => {
					this.gongenumshapeService.GongEnumShapeServiceChanged.next("post")
					this.gongenumshape = new (GongEnumShapeDB) // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.gongenumshape.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(AssociatedStruct).get(
					AssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}
		if (selectionMode == SelectionMode.MANY_MANY_ASSOCIATION_MODE) {
			dialogData.ID = this.gongenumshape.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "GongEnumShape"
			dialogData.SourceField = sourceField

			// set up the intermediate struct
			dialogData.IntermediateStruct = AssociatedStruct
			dialogData.IntermediateStructField = intermediateStructField

			// set up the end struct
			dialogData.NextAssociationStruct = nextAssociatedStruct

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(nextAssociatedStruct).get(
					nextAssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}

	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.gongenumshape.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	fillUpNameIfEmpty(event: { value: { Name: string; }; }) {
		if (this.gongenumshape.Name == "") {
			this.gongenumshape.Name = event.value.Name
		}
	}

	toggleTextArea(fieldName: string) {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			let displayAsTextArea = this.mapFields_displayAsTextArea.get(fieldName)
			this.mapFields_displayAsTextArea.set(fieldName, !displayAsTextArea)
		} else {
			this.mapFields_displayAsTextArea.set(fieldName, true)
		}
	}

	isATextArea(fieldName: string): boolean {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			return this.mapFields_displayAsTextArea.get(fieldName)!
		} else {
			return false
		}
	}

	compareObjects(o1: any, o2: any) {
		if (o1?.ID == o2?.ID) {
			return true;
		}
		else {
			return false
		}
	}
}