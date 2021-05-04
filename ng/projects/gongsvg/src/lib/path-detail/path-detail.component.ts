// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { PathDB } from '../path-db'
import { PathService } from '../path.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-path-detail',
	templateUrl: './path-detail.component.html',
	styleUrls: ['./path-detail.component.css'],
})
export class PathDetailComponent implements OnInit {

	// insertion point for declarations

	// the PathDB of interest
	path: PathDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private pathService: PathService,
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
		this.getPath()

		// observable for changes in structs
		this.pathService.PathServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getPath()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getPath(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo PathPull returned")

				if (id != 0 && association == undefined) {
					this.path = frontRepo.Paths.get(id)
				} else {
					this.path = new (PathDB)
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
		}

		if (id != 0 && association == undefined) {

			this.pathService.updatePath(this.path)
				.subscribe(path => {
					this.pathService.PathServiceChanged.next("update")

					console.log("path saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.pathService.postPath(this.path).subscribe(path => {

				this.pathService.PathServiceChanged.next("post")

				this.path = {} // reset fields
				console.log("path added")
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
			ID: this.path.ID,
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
