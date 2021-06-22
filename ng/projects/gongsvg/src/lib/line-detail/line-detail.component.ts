// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { LineDB } from '../line-db'
import { LineService } from '../line.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-line-detail',
	templateUrl: './line-detail.component.html',
	styleUrls: ['./line-detail.component.css'],
})
export class LineDetailComponent implements OnInit {

	// insertion point for declarations

	// the LineDB of interest
	line: LineDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private lineService: LineService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.getLine()

		// observable for changes in structs
		this.lineService.LineServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getLine()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getLine(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.line = frontRepo.Lines.get(id)
				} else {
					this.line = new (LineDB)
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
			if (this.line.SVG_Lines_reverse != undefined) {
				if (this.line.SVG_LinesDBID == undefined) {
					this.line.SVG_LinesDBID = new NullInt64
				}
				this.line.SVG_LinesDBID.Int64 = this.line.SVG_Lines_reverse.ID
				this.line.SVG_LinesDBID.Valid = true
				if (this.line.SVG_LinesDBID_Index == undefined) {
					this.line.SVG_LinesDBID_Index = new NullInt64
				}
				this.line.SVG_LinesDBID_Index.Valid = true
				this.line.SVG_Lines_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.lineService.updateLine(this.line)
				.subscribe(line => {
					this.lineService.LineServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "SVG_Lines":
					this.line.SVG_LinesDBID = new NullInt64
					this.line.SVG_LinesDBID.Int64 = id
					this.line.SVG_LinesDBID.Valid = true
					this.line.SVG_LinesDBID_Index = new NullInt64
					this.line.SVG_LinesDBID_Index.Valid = true
					break
			}
			this.lineService.postLine(this.line).subscribe(line => {

				this.lineService.LineServiceChanged.next("post")

				this.line = {} // reset fields
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
			ID: this.line.ID,
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
			ID: this.line.ID,
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
		if (this.line.Name == undefined) {
			this.line.Name = event.value.Name		
		}
	}
}
