import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest } from 'rxjs';

// insertion point sub template for services imports 
import { SVGDB } from './svg-db'
import { SVGService } from './svg.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  SVGs_array = new Array<SVGDB>(); // array of repo instances
  SVGs = new Map<number, SVGDB>(); // map of repo instances
  SVGs_batch = new Map<number, SVGDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

// define the type of nullable Int64 in order to support back pointers IDs
export class NullInt64 {
  Int64: number
  Valid: boolean
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number; // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string; // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean; // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode;

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string;  // The "Aclass"
  SourceField: string; // the "AnarrayofbUse"
  IntermediateStruct: string; // the "AclassBclassUse" 
  IntermediateStructField: string; // the "Bclass" as field
  NextAssociationStruct: string; // the "Bclass"
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template 
    private svgService: SVGService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service"]
    service["post" + structName](instanceToBePosted).subscribe(
      instance => {
        service[structName + "ServiceChanged"].next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service"]
    service["delete" + structName](instanceToBeDeleted).subscribe(
      instance => {
        service[structName + "ServiceChanged"].next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<SVGDB[]>,
  ] = [ // insertion point sub template 
      this.svgService.getSVGs(),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            svgs_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var svgs: SVGDB[]
            svgs = svgs_

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.SVGs_array = svgs

            // clear the map that counts SVG in the GET
            FrontRepoSingloton.SVGs_batch.clear()

            svgs.forEach(
              svg => {
                FrontRepoSingloton.SVGs.set(svg.ID, svg)
                FrontRepoSingloton.SVGs_batch.set(svg.ID, svg)
              }
            )

            // clear svgs that are absent from the batch
            FrontRepoSingloton.SVGs.forEach(
              svg => {
                if (FrontRepoSingloton.SVGs_batch.get(svg.ID) == undefined) {
                  FrontRepoSingloton.SVGs.delete(svg.ID)
                }
              }
            )

            // sort SVGs_array array
            FrontRepoSingloton.SVGs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            svgs.forEach(
              svg => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // SVGPull performs a GET on SVG of the stack and redeem association pointers 
  SVGPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.svgService.getSVGs()
        ]).subscribe(
          ([ // insertion point sub template 
            svgs,
          ]) => {
            // init the array
            FrontRepoSingloton.SVGs_array = svgs

            // clear the map that counts SVG in the GET
            FrontRepoSingloton.SVGs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            svgs.forEach(
              svg => {
                FrontRepoSingloton.SVGs.set(svg.ID, svg)
                FrontRepoSingloton.SVGs_batch.set(svg.ID, svg)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear svgs that are absent from the GET
            FrontRepoSingloton.SVGs.forEach(
              svg => {
                if (FrontRepoSingloton.SVGs_batch.get(svg.ID) == undefined) {
                  FrontRepoSingloton.SVGs.delete(svg.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getSVGUniqueID(id: number): number {
  return 31 * id
}
