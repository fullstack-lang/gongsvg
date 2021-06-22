// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { CircleDB } from '../circle-db'
import { CircleService } from '../circle.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-circle-detail',
	templateUrl: './circle-detail.component.html',
	styleUrls: ['./circle-detail.component.css'],
})
export class CircleDetailComponent implements OnInit {

	// insertion point for declarations

	// the CircleDB of interest
	circle: CircleDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private circleService: CircleService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.getCircle()

		// observable for changes in structs
		this.circleService.CircleServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getCircle()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getCircle(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.circle = frontRepo.Circles.get(id)
				} else {
					this.circle = new (CircleDB)
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
			if (this.circle.SVG_Circles_reverse != undefined) {
				if (this.circle.SVG_CirclesDBID == undefined) {
					this.circle.SVG_CirclesDBID = new NullInt64
				}
				this.circle.SVG_CirclesDBID.Int64 = this.circle.SVG_Circles_reverse.ID
				this.circle.SVG_CirclesDBID.Valid = true
				if (this.circle.SVG_CirclesDBID_Index == undefined) {
					this.circle.SVG_CirclesDBID_Index = new NullInt64
				}
				this.circle.SVG_CirclesDBID_Index.Valid = true
				this.circle.SVG_Circles_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.circleService.updateCircle(this.circle)
				.subscribe(circle => {
					this.circleService.CircleServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "SVG_Circles":
					this.circle.SVG_CirclesDBID = new NullInt64
					this.circle.SVG_CirclesDBID.Int64 = id
					this.circle.SVG_CirclesDBID.Valid = true
					this.circle.SVG_CirclesDBID_Index = new NullInt64
					this.circle.SVG_CirclesDBID_Index.Valid = true
					break
			}
			this.circleService.postCircle(this.circle).subscribe(circle => {

				this.circleService.CircleServiceChanged.next("post")

				this.circle = {} // reset fields
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
			ID: this.circle.ID,
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
			ID: this.circle.ID,
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
		if (this.circle.Name == undefined) {
			this.circle.Name = event.value.Name		
		}
	}
}
