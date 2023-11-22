import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongsvg from 'gongsvg'

import { GongdocModule } from 'gongdoc'
import { GongdocspecificModule } from 'gongdocspecific'

import { GongtreeModule } from 'gongtree'
import { GongtreespecificModule } from 'gongtreespecific'

import { GongtableModule } from 'gongtable'
import { GongtablespecificModule } from 'gongtablespecific'
import { ISplitDirection } from 'angular-split';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  toogleEditable() {
    this.mySVG.IsEditable = !this.mySVG.IsEditable
    this.svgService.updateSVG(this.mySVG, this.StackName, this.frontRepo).subscribe(
      () => {
        console.log("SVG Is editale toggled")
      }
    )
  }

  default = 'SVG Data/Model'
  svgView = 'SVG rendering'
  view = this.svgView

  mySVG: gongsvg.SVGDB = new gongsvg.SVGDB
  frontRepo: gongsvg.FrontRepo = new gongsvg.FrontRepo

  views: string[] = [this.svgView, this.default];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackType = "github.com/fullstack-lang/gongsvg/go/models"
  StackName = gongsvg.StackName.StackNameDefault
  loading = true
  constructor(
    private gongsvgFrontRepoService: gongsvg.FrontRepoService,
    private svgService: gongsvg.SVGService,
  ) {

  }

  direction: ISplitDirection = 'horizontal'
  sizes = {
    percentWithoutWildcards: {
      area1: 30,
      area2: 70,
    },
    percentWithWildcards: {
      area1: '*',
      area2: 20,
      area3: 10,
    },
    pixel: {
      area1: 120,
      area2: '*',
      area3: 160,
    },
  }

  dragEndPixel(event: any) {

    // const dragEvent = event as unknown as DragEvent
    console.log(event)
  }

  ngOnInit(): void {
    this.loading = false

    // get all svgs and assigns the first one to mysvg
    this.gongsvgFrontRepoService.pull(this.StackName).subscribe(
      gongsvgsFrontRepo => {
        this.frontRepo = gongsvgsFrontRepo
        for (let svg of gongsvgsFrontRepo.SVGs_array) {
          this.mySVG = svg
        }
      }
    )
  }
}

export type DragEvent = {
  gutterNum: number;
  sizes: Array<number>;
}