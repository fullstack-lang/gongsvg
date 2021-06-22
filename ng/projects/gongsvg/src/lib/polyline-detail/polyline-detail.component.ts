// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { PolylineDB } from '../polyline-db'
import { PolylineService } from '../polyline.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-polyline-detail',
	templateUrl: './polyline-detail.component.html',
	styleUrls: ['./polyline-detail.component.css'],
})
export class PolylineDetailComponent implements OnInit {

	// insertion point for declarations

	// the PolylineDB of interest
	polyline: PolylineDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private polylineService: PolylineService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.getPolyline()

		// observable for changes in structs
		this.polylineService.PolylineServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getPolyline()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getPolyline(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.polyline = frontRepo.Polylines.get(id)
				} else {
					this.polyline = new (PolylineDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization
		
		// insertion point for translation/nullation of each field
		
		// save from the front pointer space to the non pointer space for serialization
		if (association == undefined) {
			// insertion point for translation/nullation of each pointers
			if (this.polyline.SVG_Polylines_reverse != undefined) {
				if (this.polyline.SVG_PolylinesDBID == undefined) {
					this.polyline.SVG_PolylinesDBID = new NullInt64
				}
				this.polyline.SVG_PolylinesDBID.Int64 = this.polyline.SVG_Polylines_reverse.ID
				this.polyline.SVG_PolylinesDBID.Valid = true
				if (this.polyline.SVG_PolylinesDBID_Index == undefined) {
					this.polyline.SVG_PolylinesDBID_Index = new NullInt64
				}
				this.polyline.SVG_PolylinesDBID_Index.Valid = true
				this.polyline.SVG_Polylines_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.polylineService.updatePolyline(this.polyline)
				.subscribe(polyline => {
					this.polylineService.PolylineServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "SVG_Polylines":
					this.polyline.SVG_PolylinesDBID = new NullInt64
					this.polyline.SVG_PolylinesDBID.Int64 = id
					this.polyline.SVG_PolylinesDBID.Valid = true
					this.polyline.SVG_PolylinesDBID_Index = new NullInt64
					this.polyline.SVG_PolylinesDBID_Index.Valid = true
					break
			}
			this.polylineService.postPolyline(this.polyline).subscribe(polyline => {

				this.polylineService.PolylineServiceChanged.next("post")

				this.polyline = {} // reset fields
			});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		dialogConfig.data = {
			ID: this.polyline.ID,
			ReversePointer: reverseField,
			OrderingMode: false,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'sTableComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.polyline.ID,
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

	fillUpNameIfEmpty(event) {
		if (this.polyline.Name == undefined) {
			this.polyline.Name = event.value.Name		
		}
	}
}
