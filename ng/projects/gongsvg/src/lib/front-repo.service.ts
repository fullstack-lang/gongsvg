import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest } from 'rxjs';

// insertion point sub template for services imports 
import { RectDB } from './rect-db'
import { RectService } from './rect.service'

import { SVGDB } from './svg-db'
import { SVGService } from './svg.service'

import { TextDB } from './text-db'
import { TextService } from './text.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Rects_array = new Array<RectDB>(); // array of repo instances
  Rects = new Map<number, RectDB>(); // map of repo instances
  Rects_batch = new Map<number, RectDB>(); // same but only in last GET (for finding repo instances to delete)
  SVGs_array = new Array<SVGDB>(); // array of repo instances
  SVGs = new Map<number, SVGDB>(); // map of repo instances
  SVGs_batch = new Map<number, SVGDB>(); // same but only in last GET (for finding repo instances to delete)
  Texts_array = new Array<TextDB>(); // array of repo instances
  Texts = new Map<number, TextDB>(); // map of repo instances
  Texts_batch = new Map<number, TextDB>(); // same but only in last GET (for finding repo instances to delete)
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

// define the interface for information that is forwarded from the calling instance to 
// the select table
export interface DialogData {
  ID: number; // ID of the calling instance
  ReversePointer: string; // field of {{Structname}} that serve as reverse pointer
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
    private rectService: RectService,
    private svgService: SVGService,
    private textService: TextService,
  ) { }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<RectDB[]>,
    Observable<SVGDB[]>,
    Observable<TextDB[]>,
  ] = [ // insertion point sub template 
      this.rectService.getRects(),
      this.svgService.getSVGs(),
      this.textService.getTexts(),
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
            rects_,
            svgs_,
            texts_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var rects: RectDB[]
            rects = rects_
            var svgs: SVGDB[]
            svgs = svgs_
            var texts: TextDB[]
            texts = texts_

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.Rects_array = rects

            // clear the map that counts Rect in the GET
            FrontRepoSingloton.Rects_batch.clear()
            
            rects.forEach(
              rect => {
                FrontRepoSingloton.Rects.set(rect.ID, rect)
                FrontRepoSingloton.Rects_batch.set(rect.ID, rect)
              }
            )
            
            // clear rects that are absent from the batch
            FrontRepoSingloton.Rects.forEach(
              rect => {
                if (FrontRepoSingloton.Rects_batch.get(rect.ID) == undefined) {
                  FrontRepoSingloton.Rects.delete(rect.ID)
                }
              }
            )
            
            // sort Rects_array array
            FrontRepoSingloton.Rects_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });
            
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
            
            // init the array
            FrontRepoSingloton.Texts_array = texts

            // clear the map that counts Text in the GET
            FrontRepoSingloton.Texts_batch.clear()
            
            texts.forEach(
              text => {
                FrontRepoSingloton.Texts.set(text.ID, text)
                FrontRepoSingloton.Texts_batch.set(text.ID, text)
              }
            )
            
            // clear texts that are absent from the batch
            FrontRepoSingloton.Texts.forEach(
              text => {
                if (FrontRepoSingloton.Texts_batch.get(text.ID) == undefined) {
                  FrontRepoSingloton.Texts.delete(text.ID)
                }
              }
            )
            
            // sort Texts_array array
            FrontRepoSingloton.Texts_array.sort((t1, t2) => {
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
            rects.forEach(
              rect => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Rects redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(rect.SVG_RectsDBID.Int64)
                  if (_svg) {
                    if (_svg.Rects == undefined) {
                      _svg.Rects = new Array<SVGDB>()
                    }
                    _svg.Rects.push(rect)
                    if (rect.SVG_Rects_reverse == undefined) {
                      rect.SVG_Rects_reverse = _svg
                    }
                  }
                }
              }
            )
            svgs.forEach(
              svg => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            texts.forEach(
              text => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field SVG.Texts redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(text.SVG_TextsDBID.Int64)
                  if (_svg) {
                    if (_svg.Texts == undefined) {
                      _svg.Texts = new Array<SVGDB>()
                    }
                    _svg.Texts.push(text)
                    if (text.SVG_Texts_reverse == undefined) {
                      text.SVG_Texts_reverse = _svg
                    }
                  }
                }
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

  // RectPull performs a GET on Rect of the stack and redeem association pointers 
  RectPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.rectService.getRects()
        ]).subscribe(
          ([ // insertion point sub template 
            rects,
          ]) => {
            // init the array
            FrontRepoSingloton.Rects_array = rects

            // clear the map that counts Rect in the GET
            FrontRepoSingloton.Rects_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            rects.forEach(
              rect => {
                FrontRepoSingloton.Rects.set(rect.ID, rect)
                FrontRepoSingloton.Rects_batch.set(rect.ID, rect)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field SVG.Rects redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(rect.SVG_RectsDBID.Int64)
                  if (_svg) {
                    if (_svg.Rects == undefined) {
                      _svg.Rects = new Array<SVGDB>()
                    }
                    _svg.Rects.push(rect)
                    if (rect.SVG_Rects_reverse == undefined) {
                      rect.SVG_Rects_reverse = _svg
                    }
                  }
                }
              }
            )

            // clear rects that are absent from the GET
            FrontRepoSingloton.Rects.forEach(
              rect => {
                if (FrontRepoSingloton.Rects_batch.get(rect.ID) == undefined) {
                  FrontRepoSingloton.Rects.delete(rect.ID)
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

  // TextPull performs a GET on Text of the stack and redeem association pointers 
  TextPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.textService.getTexts()
        ]).subscribe(
          ([ // insertion point sub template 
            texts,
          ]) => {
            // init the array
            FrontRepoSingloton.Texts_array = texts

            // clear the map that counts Text in the GET
            FrontRepoSingloton.Texts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            texts.forEach(
              text => {
                FrontRepoSingloton.Texts.set(text.ID, text)
                FrontRepoSingloton.Texts_batch.set(text.ID, text)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field SVG.Texts redeeming
                {
                  let _svg = FrontRepoSingloton.SVGs.get(text.SVG_TextsDBID.Int64)
                  if (_svg) {
                    if (_svg.Texts == undefined) {
                      _svg.Texts = new Array<SVGDB>()
                    }
                    _svg.Texts.push(text)
                    if (text.SVG_Texts_reverse == undefined) {
                      text.SVG_Texts_reverse = _svg
                    }
                  }
                }
              }
            )

            // clear texts that are absent from the GET
            FrontRepoSingloton.Texts.forEach(
              text => {
                if (FrontRepoSingloton.Texts_batch.get(text.ID) == undefined) {
                  FrontRepoSingloton.Texts.delete(text.ID)
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
export function getRectUniqueID(id: number): number {
  return 31 * id
}
export function getSVGUniqueID(id: number): number {
  return 37 * id
}
export function getTextUniqueID(id: number): number {
  return 41 * id
}
