// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { PolygoneDB } from '../polygone-db'
import { PolygoneService } from '../polygone.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-polygone-detail',
	templateUrl: './polygone-detail.component.html',
	styleUrls: ['./polygone-detail.component.css'],
})
export class PolygoneDetailComponent implements OnInit {

	// insertion point for declarations

	// the PolygoneDB of interest
	polygone: PolygoneDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private polygoneService: PolygoneService,
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
		this.getPolygone()

		// observable for changes in structs
		this.polygoneService.PolygoneServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getPolygone()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getPolygone(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo PolygonePull returned")

				if (id != 0 && association == undefined) {
					this.polygone = frontRepo.Polygones.get(id)
				} else {
					this.polygone = new (PolygoneDB)
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
			if (this.polygone.SVG_Polygones_reverse != undefined) {
				this.polygone.SVG_PolygonesDBID = new NullInt64
				this.polygone.SVG_PolygonesDBID.Int64 = this.polygone.SVG_Polygones_reverse.ID
				this.polygone.SVG_PolygonesDBID.Valid = true
				this.polygone.SVG_Polygones_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.polygoneService.updatePolygone(this.polygone)
				.subscribe(polygone => {
					this.polygoneService.PolygoneServiceChanged.next("update")

					console.log("polygone saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "SVG_Polygones":
					this.polygone.SVG_PolygonesDBID = new NullInt64
					this.polygone.SVG_PolygonesDBID.Int64 = id
					this.polygone.SVG_PolygonesDBID.Valid = true
					break
			}
			this.polygoneService.postPolygone(this.polygone).subscribe(polygone => {

				this.polygoneService.PolygoneServiceChanged.next("post")

				this.polygone = {} // reset fields
				console.log("polygone added")
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
			ID: this.polygone.ID,
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
