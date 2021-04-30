// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { SVGDB } from '../svg-db'
import { SVGService } from '../svg.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-svg-detail',
	templateUrl: './svg-detail.component.html',
	styleUrls: ['./svg-detail.component.css'],
})
export class SVGDetailComponent implements OnInit {

	// insertion point for declarations
	DisplayFormControl = new FormControl(false);

	// the SVGDB of interest
	svg: SVGDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private svgService: SVGService,
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
		this.getSVG()

		// observable for changes in structs
		this.svgService.SVGServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getSVG()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getSVG(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo SVGPull returned")

				if (id != 0 && association == undefined) {
					this.svg = frontRepo.SVGs.get(id)
				} else {
					this.svg = new (SVGDB)
				}

				// insertion point for recovery of form controls value for bool fields
				this.DisplayFormControl.setValue(this.svg.Display)
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization
		
		// insertion point for translation/nullation of each field
		this.svg.Display = this.DisplayFormControl.value
		
		// save from the front pointer space to the non pointer space for serialization
		if (association == undefined) {
			// insertion point for translation/nullation of each pointers
		}

		if (id != 0 && association == undefined) {

			this.svgService.updateSVG(this.svg)
				.subscribe(svg => {
					this.svgService.SVGServiceChanged.next("update")

					console.log("svg saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.svgService.postSVG(this.svg).subscribe(svg => {

				this.svgService.SVGServiceChanged.next("post")

				this.svg = {} // reset fields
				console.log("svg added")
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
			ID: this.svg.ID,
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
