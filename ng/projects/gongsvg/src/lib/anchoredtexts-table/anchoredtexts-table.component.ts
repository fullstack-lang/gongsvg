// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional, Input } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { ActivatedRoute, Router, RouterState } from '@angular/router';
import { AnchoredTextDB } from '../anchoredtext-db'
import { AnchoredTextService } from '../anchoredtext.service'

// insertion point for additional imports

import { RouteService } from '../route-service';

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-anchoredtextstable',
  templateUrl: './anchoredtexts-table.component.html',
  styleUrls: ['./anchoredtexts-table.component.css'],
})
export class AnchoredTextsTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of AnchoredText instances
  selection: SelectionModel<AnchoredTextDB> = new (SelectionModel)
  initialSelection = new Array<AnchoredTextDB>()

  // the data source for the table
  anchoredtexts: AnchoredTextDB[] = []
  matTableDataSource: MatTableDataSource<AnchoredTextDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.anchoredtexts
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (anchoredtextDB: AnchoredTextDB, property: string) => {
      switch (property) {
        case 'ID':
          return anchoredtextDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return anchoredtextDB.Name;

        case 'Content':
          return anchoredtextDB.Content;

        case 'X_Offset':
          return anchoredtextDB.X_Offset;

        case 'Y_Offset':
          return anchoredtextDB.Y_Offset;

        case 'Color':
          return anchoredtextDB.Color;

        case 'FillOpacity':
          return anchoredtextDB.FillOpacity;

        case 'Stroke':
          return anchoredtextDB.Stroke;

        case 'StrokeWidth':
          return anchoredtextDB.StrokeWidth;

        case 'StrokeDashArray':
          return anchoredtextDB.StrokeDashArray;

        case 'StrokeDashArrayWhenSelected':
          return anchoredtextDB.StrokeDashArrayWhenSelected;

        case 'Transform':
          return anchoredtextDB.Transform;

        case 'Link_TextAtArrowEnd':
          if (this.frontRepo.Links.get(anchoredtextDB.Link_TextAtArrowEndDBID.Int64) != undefined) {
            return this.frontRepo.Links.get(anchoredtextDB.Link_TextAtArrowEndDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (anchoredtextDB: AnchoredTextDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the anchoredtextDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += anchoredtextDB.Name.toLowerCase()
      mergedContent += anchoredtextDB.Content.toLowerCase()
      mergedContent += anchoredtextDB.X_Offset.toString()
      mergedContent += anchoredtextDB.Y_Offset.toString()
      mergedContent += anchoredtextDB.Color.toLowerCase()
      mergedContent += anchoredtextDB.FillOpacity.toString()
      mergedContent += anchoredtextDB.Stroke.toLowerCase()
      mergedContent += anchoredtextDB.StrokeWidth.toString()
      mergedContent += anchoredtextDB.StrokeDashArray.toLowerCase()
      mergedContent += anchoredtextDB.StrokeDashArrayWhenSelected.toLowerCase()
      mergedContent += anchoredtextDB.Transform.toLowerCase()
      if (anchoredtextDB.Link_TextAtArrowEndDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Links.get(anchoredtextDB.Link_TextAtArrowEndDBID.Int64)!.Name.toLowerCase()
      }


      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private anchoredtextService: AnchoredTextService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of anchoredtext instances
    public dialogRef: MatDialogRef<AnchoredTextsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
    private activatedRoute: ActivatedRoute,

    private routeService: RouteService,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      this.GONG__StackPath = dialogData.GONG__StackPath
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.anchoredtextService.AnchoredTextServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getAnchoredTexts()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Content",
        "X_Offset",
        "Y_Offset",
        "Color",
        "FillOpacity",
        "Stroke",
        "StrokeWidth",
        "StrokeDashArray",
        "StrokeDashArrayWhenSelected",
        "Transform",
        "Link_TextAtArrowEnd",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Content",
        "X_Offset",
        "Y_Offset",
        "Color",
        "FillOpacity",
        "Stroke",
        "StrokeWidth",
        "StrokeDashArray",
        "StrokeDashArrayWhenSelected",
        "Transform",
        "Link_TextAtArrowEnd",
      ]
      this.selection = new SelectionModel<AnchoredTextDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getAnchoredTexts()

    this.matTableDataSource = new MatTableDataSource(this.anchoredtexts)
  }

  getAnchoredTexts(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.anchoredtexts = this.frontRepo.AnchoredTexts_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let anchoredtext of this.anchoredtexts) {
            let ID = this.dialogData.ID
            let revPointer = anchoredtext[this.dialogData.ReversePointer as keyof AnchoredTextDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(anchoredtext)
            }
            this.selection = new SelectionModel<AnchoredTextDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, AnchoredTextDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to AnchoredTextDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as AnchoredTextDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let anchoredtext = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as AnchoredTextDB
              this.initialSelection.push(anchoredtext)
            }
          }

          this.selection = new SelectionModel<AnchoredTextDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.anchoredtexts
      }
    )
  }

  // newAnchoredText initiate a new anchoredtext
  // create a new AnchoredText objet
  newAnchoredText() {
  }

  deleteAnchoredText(anchoredtextID: number, anchoredtext: AnchoredTextDB) {
    // list of anchoredtexts is truncated of anchoredtext before the delete
    this.anchoredtexts = this.anchoredtexts.filter(h => h !== anchoredtext);

    this.anchoredtextService.deleteAnchoredText(anchoredtextID, this.GONG__StackPath).subscribe(
      anchoredtext => {
        this.anchoredtextService.AnchoredTextServiceChanged.next("delete")
      }
    );
  }

  editAnchoredText(anchoredtextID: number, anchoredtext: AnchoredTextDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(anchoredtextID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "anchoredtext" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, anchoredtextID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.anchoredtexts.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.anchoredtexts.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<AnchoredTextDB>()

      // reset all initial selection of anchoredtext that belong to anchoredtext
      for (let anchoredtext of this.initialSelection) {
        let index = anchoredtext[this.dialogData.ReversePointer as keyof AnchoredTextDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(anchoredtext)

      }

      // from selection, set anchoredtext that belong to anchoredtext
      for (let anchoredtext of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = anchoredtext[this.dialogData.ReversePointer as keyof AnchoredTextDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(anchoredtext)
      }


      // update all anchoredtext (only update selection & initial selection)
      for (let anchoredtext of toUpdate) {
        this.anchoredtextService.updateAnchoredText(anchoredtext, this.GONG__StackPath)
          .subscribe(anchoredtext => {
            this.anchoredtextService.AnchoredTextServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, AnchoredTextDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedAnchoredText = new Set<number>()
      for (let anchoredtext of this.initialSelection) {
        if (this.selection.selected.includes(anchoredtext)) {
          // console.log("anchoredtext " + anchoredtext.Name + " is still selected")
        } else {
          console.log("anchoredtext " + anchoredtext.Name + " has been unselected")
          unselectedAnchoredText.add(anchoredtext.ID)
          console.log("is unselected " + unselectedAnchoredText.has(anchoredtext.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let anchoredtext = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as AnchoredTextDB
      if (unselectedAnchoredText.has(anchoredtext.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<AnchoredTextDB>) = new Array<AnchoredTextDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          anchoredtext => {
            if (!this.initialSelection.includes(anchoredtext)) {
              // console.log("anchoredtext " + anchoredtext.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + anchoredtext.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = anchoredtext.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = anchoredtext.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("anchoredtext " + anchoredtext.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<AnchoredTextDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}