// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { CircleDB } from '../circle-db'
import { CircleService } from '../circle.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

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
		// https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
		// this is for routerLink on same component when only queryParameter changes
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
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
				console.log("front repo CirclePull returned")

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
				this.circle.SVG_CirclesDBID = new NullInt64
				this.circle.SVG_CirclesDBID.Int64 = this.circle.SVG_Circles_reverse.ID
				this.circle.SVG_CirclesDBID.Valid = true
				this.circle.SVG_Circles_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.circleService.updateCircle(this.circle)
				.subscribe(circle => {
					this.circleService.CircleServiceChanged.next("update")

					console.log("circle saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "SVG_Circles":
					this.circle.SVG_CirclesDBID = new NullInt64
					this.circle.SVG_CirclesDBID.Int64 = id
					this.circle.SVG_CirclesDBID.Valid = true
					break
			}
			this.circleService.postCircle(this.circle).subscribe(circle => {

				this.circleService.CircleServiceChanged.next("post")

				this.circle = {} // reset fields
				console.log("circle added")
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
		dialogConfig.data = {
			ID: this.circle.ID,
			ReversePointer: reverseField,
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
}
