// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { FormDivDB } from '../formdiv-db'
import { FormDivService } from '../formdiv.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-formdiv-sorting',
  templateUrl: './formdiv-sorting.component.html',
  styleUrls: ['./formdiv-sorting.component.css']
})
export class FormDivSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of FormDiv instances that are in the association
  associatedFormDivs = new Array<FormDivDB>();

  constructor(
    private formdivService: FormDivService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of formdiv instances
    public dialogRef: MatDialogRef<FormDivSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getFormDivs()
  }

  getFormDivs(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let formdiv of this.frontRepo.FormDivs_array) {
          let ID = this.dialogData.ID
          let revPointerID = formdiv[this.dialogData.ReversePointer as keyof FormDivDB] as unknown as NullInt64
          let revPointerID_Index = formdiv[this.dialogData.ReversePointer + "_Index" as keyof FormDivDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedFormDivs.push(formdiv)
          }
        }

        // sort associated formdiv according to order
        this.associatedFormDivs.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer + "_Index" as keyof typeof t1] as unknown as NullInt64
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer + "_Index" as keyof typeof t2] as unknown as NullInt64
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedFormDivs, event.previousIndex, event.currentIndex);

    // set the order of FormDiv instances
    let index = 0

    for (let formdiv of this.associatedFormDivs) {
      let revPointerID_Index = formdiv[this.dialogData.ReversePointer + "_Index" as keyof FormDivDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedFormDivs.forEach(
      formdiv => {
        this.formdivService.updateFormDiv(formdiv, this.dialogData.GONG__StackPath)
          .subscribe(formdiv => {
            this.formdivService.FormDivServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}