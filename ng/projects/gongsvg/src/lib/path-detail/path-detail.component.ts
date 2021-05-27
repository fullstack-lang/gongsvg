// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { PathDB } from '../path-db'
import { PathService } from '../path.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

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
			if (this.path.SVG_Paths_reverse != undefined) {
				this.path.SVG_PathsDBID = new NullInt64
				this.path.SVG_PathsDBID.Int64 = this.path.SVG_Paths_reverse.ID
				this.path.SVG_PathsDBID.Valid = true
				this.path.SVG_PathsDBID_Index.Valid = true
				this.path.SVG_Paths_reverse = undefined // very important, otherwise, circular JSON
			}
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
				case "SVG_Paths":
					this.path.SVG_PathsDBID = new NullInt64
					this.path.SVG_PathsDBID.Int64 = id
					this.path.SVG_PathsDBID.Valid = true
					this.path.SVG_PathsDBID_Index.Valid = true
					break
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
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		dialogConfig.data = {
			ID: this.path.ID,
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
			ID: this.path.ID,
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
