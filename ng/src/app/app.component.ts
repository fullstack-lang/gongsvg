import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongsvg from 'gongsvg'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  view = 'Default view'
  default = 'Default view'
  data = 'Data view'
  diagrams = 'Diagrams view'
  meta = 'Meta view'

  views: string[] = [this.default, this.data, this.diagrams, this.meta];

  // variable that enables pooling of selected gongstruct
  obsTimer: Observable<number> = timer(1000, 1000)
  lastSelectionDate: string = ''

  constructor(private gongdocGongStructShapeService: gongdoc.GongStructShapeService,
    private gongstructSelectionService: gongsvg.GongstructSelectionService
  ) {

  }

  ngOnInit(): void {

    // pool the gongdoc command and check wether a gongstruct has been selected
    this.obsTimer.subscribe(
      currTime => {
        // pool all gongstructshapes and find which one is selected
        this.gongdocGongStructShapeService.getGongStructShapes().subscribe(
          gongstructshapes => {
            for (let gongstructshape of gongstructshapes) {
              if (gongstructshape.IsSelected) {
                gongstructshape.IsSelected = false
                // console.log("gongstructshape " + gongstructshape.ReferenceName + " is selected")
                this.gongdocGongStructShapeService.updateGongStructShape(gongstructshape, "").subscribe(
                  gongstructshape2 => {
                    // console.log("gongstructshape has been unselected")
                  }
                )
                this.gongstructSelectionService.gongstructSelected(gongstructshape.Identifier)
              }
            }
          }
        )
      }
    )
  }
}
