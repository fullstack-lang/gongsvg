// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { RectDB } from '../rect-db'
import { RectService } from '../rect.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-rects-table',
  templateUrl: './rects-table.component.html',
  styleUrls: ['./rects-table.component.css'],
})
export class RectsTableComponent implements OnInit {

  // used if the component is called as a selection component of Rect instances
  selection: SelectionModel<RectDB>;
  initialSelection = new Array<RectDB>();

  // the data source for the table
  rects: RectDB[];

  // front repo, that will be referenced by this.rects
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private rectService: RectService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of rect instances
    public dialogRef: MatDialogRef<RectsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.rectService.RectServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getRects()
        }
      }
    )
    if (dialogData == undefined) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "X",
        "Y",
        "Width",
        "Height",
        "Rects",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "X",
        "Y",
        "Width",
        "Height",
        "Rects",
      ]
      this.selection = new SelectionModel<RectDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getRects()
  }

  getRects(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.rects = this.frontRepo.Rects_array;

        // insertion point for variables Recoveries

        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.rects.forEach(
            rect => {
              let ID = this.dialogData.ID
              let revPointer = rect[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push(rect)
              }
            }
          )
          this.selection = new SelectionModel<RectDB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // newRect initiate a new rect
  // create a new Rect objet
  newRect() {
  }

  deleteRect(rectID: number, rect: RectDB) {
    // list of rects is truncated of rect before the delete
    this.rects = this.rects.filter(h => h !== rect);

    this.rectService.deleteRect(rectID).subscribe(
      rect => {
        this.rectService.RectServiceChanged.next("delete")

        console.log("rect deleted")
      }
    );
  }

  editRect(rectID: number, rect: RectDB) {

  }

  // display rect in router
  displayRectInRouter(rectID: number) {
    this.router.navigate( ["rect-display", rectID])
  }

  // set editor outlet
  setEditorRouterOutlet(rectID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["rect-detail", rectID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(rectID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["rect-presentation", rectID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.rects.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.rects.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<RectDB>()

    // reset all initial selection of rect that belong to rect through Anarrayofb
    this.initialSelection.forEach(
      rect => {
        rect[this.dialogData.ReversePointer].Int64 = 0
        rect[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(rect)
      }
    )

    // from selection, set rect that belong to rect through Anarrayofb
    this.selection.selected.forEach(
      rect => {
        console.log("selection ID " + rect.ID)
        let ID = +this.dialogData.ID
        rect[this.dialogData.ReversePointer].Int64 = ID
        rect[this.dialogData.ReversePointer].Valid = true
        toUpdate.add(rect)
      }
    )

    // update all rect (only update selection & initial selection)
    toUpdate.forEach(
      rect => {
        this.rectService.updateRect(rect)
          .subscribe(rect => {
            this.rectService.RectServiceChanged.next("update")
            console.log("rect saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
