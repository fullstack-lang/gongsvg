// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { RectDB } from '../rect-db'
import { RectService } from '../rect.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-rect-detail',
	templateUrl: './rect-detail.component.html',
	styleUrls: ['./rect-detail.component.css'],
})
export class RectDetailComponent implements OnInit {

	// insertion point for declarations

	// the RectDB of interest
	rect: RectDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private rectService: RectService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
		// https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
		// this is for routerLink on same component when only queryParameter changes
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getRect()

		// observable for changes in structs
		this.rectService.RectServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getRect()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getRect(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo RectPull returned")

				if (id != 0 && association == undefined) {
					this.rect = frontRepo.Rects.get(id)
				} else {
					this.rect = new (RectDB)
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
			if (this.rect.SVG_Rects_reverse != undefined) {
				this.rect.SVG_RectsDBID = new NullInt64
				this.rect.SVG_RectsDBID.Int64 = this.rect.SVG_Rects_reverse.ID
				this.rect.SVG_RectsDBID.Valid = true
				this.rect.SVG_RectsDBID_Index.Valid = true
				this.rect.SVG_Rects_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.rectService.updateRect(this.rect)
				.subscribe(rect => {
					this.rectService.RectServiceChanged.next("update")

					console.log("rect saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "SVG_Rects":
					this.rect.SVG_RectsDBID = new NullInt64
					this.rect.SVG_RectsDBID.Int64 = id
					this.rect.SVG_RectsDBID.Valid = true
					this.rect.SVG_RectsDBID_Index.Valid = true
					break
			}
			this.rectService.postRect(this.rect).subscribe(rect => {

				this.rectService.RectServiceChanged.next("post")

				this.rect = {} // reset fields
				console.log("rect added")
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
			ID: this.rect.ID,
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
			console.log('The dialog was closed');
		});
	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.rect.ID,
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
			console.log('The dialog was closed');
		});
	}
}
