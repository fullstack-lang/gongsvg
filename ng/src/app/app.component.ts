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
