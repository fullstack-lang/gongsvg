// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional } from '@angular/core';
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

import { Router, RouterState } from '@angular/router';
import { PolygoneDB } from '../polygone-db'
import { PolygoneService } from '../polygone.service'

// insertion point for additional imports

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-polygonestable',
  templateUrl: './polygones-table.component.html',
  styleUrls: ['./polygones-table.component.css'],
})
export class PolygonesTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Polygone instances
  selection: SelectionModel<PolygoneDB> = new (SelectionModel)
  initialSelection = new Array<PolygoneDB>()

  // the data source for the table
  polygones: PolygoneDB[] = []
  matTableDataSource: MatTableDataSource<PolygoneDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.polygones
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
    this.matTableDataSource.sortingDataAccessor = (polygoneDB: PolygoneDB, property: string) => {
      switch (property) {
        case 'ID':
          return polygoneDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return polygoneDB.Name;

        case 'Points':
          return polygoneDB.Points;

        case 'Color':
          return polygoneDB.Color;

        case 'FillOpacity':
          return polygoneDB.FillOpacity;

        case 'Stroke':
          return polygoneDB.Stroke;

        case 'StrokeWidth':
          return polygoneDB.StrokeWidth;

        case 'StrokeDashArray':
          return polygoneDB.StrokeDashArray;

        case 'Transform':
          return polygoneDB.Transform;

        case 'SVG_Polygones':
          if (this.frontRepo.SVGs.get(polygoneDB.SVG_PolygonesDBID.Int64) != undefined) {
            return this.frontRepo.SVGs.get(polygoneDB.SVG_PolygonesDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (polygoneDB: PolygoneDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the polygoneDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += polygoneDB.Name.toLowerCase()
      mergedContent += polygoneDB.Points.toLowerCase()
      mergedContent += polygoneDB.Color.toLowerCase()
      mergedContent += polygoneDB.FillOpacity.toString()
      mergedContent += polygoneDB.Stroke.toLowerCase()
      mergedContent += polygoneDB.StrokeWidth.toString()
      mergedContent += polygoneDB.StrokeDashArray.toLowerCase()
      mergedContent += polygoneDB.Transform.toLowerCase()
      if (polygoneDB.SVG_PolygonesDBID.Int64 != 0) {
        mergedContent += this.frontRepo.SVGs.get(polygoneDB.SVG_PolygonesDBID.Int64)!.Name.toLowerCase()
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
    private polygoneService: PolygoneService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of polygone instances
    public dialogRef: MatDialogRef<PolygonesTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
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
    this.polygoneService.PolygoneServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getPolygones()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Points",
        "Color",
        "FillOpacity",
        "Stroke",
        "StrokeWidth",
        "StrokeDashArray",
        "Transform",
        "SVG_Polygones",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Points",
        "Color",
        "FillOpacity",
        "Stroke",
        "StrokeWidth",
        "StrokeDashArray",
        "Transform",
        "SVG_Polygones",
      ]
      this.selection = new SelectionModel<PolygoneDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getPolygones()
    this.matTableDataSource = new MatTableDataSource(this.polygones)
  }

  getPolygones(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.polygones = this.frontRepo.Polygones_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let polygone of this.polygones) {
            let ID = this.dialogData.ID
            let revPointer = polygone[this.dialogData.ReversePointer as keyof PolygoneDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(polygone)
            }
            this.selection = new SelectionModel<PolygoneDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, PolygoneDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as PolygoneDB[]
          for (let associationInstance of sourceField) {
            let polygone = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as PolygoneDB
            this.initialSelection.push(polygone)
          }

          this.selection = new SelectionModel<PolygoneDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.polygones
      }
    )
  }

  // newPolygone initiate a new polygone
  // create a new Polygone objet
  newPolygone() {
  }

  deletePolygone(polygoneID: number, polygone: PolygoneDB) {
    // list of polygones is truncated of polygone before the delete
    this.polygones = this.polygones.filter(h => h !== polygone);

    this.polygoneService.deletePolygone(polygoneID).subscribe(
      polygone => {
        this.polygoneService.PolygoneServiceChanged.next("delete")
      }
    );
  }

  editPolygone(polygoneID: number, polygone: PolygoneDB) {

  }

  // display polygone in router
  displayPolygoneInRouter(polygoneID: number) {
    this.router.navigate(["github_com_fullstack_lang_gongsvg_go-" + "polygone-display", polygoneID])
  }

  // set editor outlet
  setEditorRouterOutlet(polygoneID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongsvg_go_editor: ["github_com_fullstack_lang_gongsvg_go-" + "polygone-detail", polygoneID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(polygoneID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongsvg_go_presentation: ["github_com_fullstack_lang_gongsvg_go-" + "polygone-presentation", polygoneID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.polygones.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.polygones.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<PolygoneDB>()

      // reset all initial selection of polygone that belong to polygone
      for (let polygone of this.initialSelection) {
        let index = polygone[this.dialogData.ReversePointer as keyof PolygoneDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(polygone)

      }

      // from selection, set polygone that belong to polygone
      for (let polygone of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = polygone[this.dialogData.ReversePointer as keyof PolygoneDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(polygone)
      }


      // update all polygone (only update selection & initial selection)
      for (let polygone of toUpdate) {
        this.polygoneService.updatePolygone(polygone)
          .subscribe(polygone => {
            this.polygoneService.PolygoneServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, PolygoneDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedPolygone = new Set<number>()
      for (let polygone of this.initialSelection) {
        if (this.selection.selected.includes(polygone)) {
          // console.log("polygone " + polygone.Name + " is still selected")
        } else {
          console.log("polygone " + polygone.Name + " has been unselected")
          unselectedPolygone.add(polygone.ID)
          console.log("is unselected " + unselectedPolygone.has(polygone.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let polygone = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as PolygoneDB
      if (unselectedPolygone.has(polygone.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<PolygoneDB>) = new Array<PolygoneDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          polygone => {
            if (!this.initialSelection.includes(polygone)) {
              // console.log("polygone " + polygone.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + polygone.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = polygone.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = polygone.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("polygone " + polygone.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<PolygoneDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
