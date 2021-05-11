// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { EllipseDB } from '../ellipse-db'
import { EllipseService } from '../ellipse.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-ellipse-detail',
	templateUrl: './ellipse-detail.component.html',
	styleUrls: ['./ellipse-detail.component.css'],
})
export class EllipseDetailComponent implements OnInit {

	// insertion point for declarations

	// the EllipseDB of interest
	ellipse: EllipseDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private ellipseService: EllipseService,
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
		this.getEllipse()

		// observable for changes in structs
		this.ellipseService.EllipseServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getEllipse()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getEllipse(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo EllipsePull returned")

				if (id != 0 && association == undefined) {
					this.ellipse = frontRepo.Ellipses.get(id)
				} else {
					this.ellipse = new (EllipseDB)
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
			if (this.ellipse.SVG_Ellipses_reverse != undefined) {
				this.ellipse.SVG_EllipsesDBID = new NullInt64
				this.ellipse.SVG_EllipsesDBID.Int64 = this.ellipse.SVG_Ellipses_reverse.ID
				this.ellipse.SVG_EllipsesDBID.Valid = true
				this.ellipse.SVG_EllipsesDBID_Index.Valid = true
				this.ellipse.SVG_Ellipses_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.ellipseService.updateEllipse(this.ellipse)
				.subscribe(ellipse => {
					this.ellipseService.EllipseServiceChanged.next("update")

					console.log("ellipse saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "SVG_Ellipses":
					this.ellipse.SVG_EllipsesDBID = new NullInt64
					this.ellipse.SVG_EllipsesDBID.Int64 = id
					this.ellipse.SVG_EllipsesDBID.Valid = true
					this.ellipse.SVG_EllipsesDBID_Index.Valid = true
					break
			}
			this.ellipseService.postEllipse(this.ellipse).subscribe(ellipse => {

				this.ellipseService.EllipseServiceChanged.next("post")

				this.ellipse = {} // reset fields
				console.log("ellipse added")
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
			ID: this.ellipse.ID,
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
			ID: this.ellipse.ID,
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
