import { Component, OnInit } from '@angular/core';

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';


import * as gongsvg from '../../projects/gongsvg/src/public-api'

import { GongsvgDiagrammingComponent } from '../../projects/gongsvgspecific/src/lib/gongsvg-diagramming/gongsvg-diagramming'

import { ISplitDirection } from 'angular-split';
import { AngularSplitModule } from 'angular-split';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  standalone: true,
  imports: [
    MatRadioModule,
    FormsModule,
    CommonModule,
    MatButtonModule,
    MatIconModule,

    AngularSplitModule,

    GongsvgDiagrammingComponent,

  ],
})
export class AppComponent implements OnInit {

  toogleEditable() {
    this.mySVG.IsEditable = !this.mySVG.IsEditable

    this.svgService.updateFront(this.mySVG, this.StackName).subscribe(
      svg => {
        console.log("SVG Is editale toggled", svg.IsEditable)
      }
    )
  }

  default = 'SVG Data/Model'

  diagramSvgView = 'Diagram SVG'

  view = this.diagramSvgView

  mySVG: gongsvg.SVG = new gongsvg.SVG
  frontRepo: gongsvg.FrontRepo = new gongsvg.FrontRepo

  views: string[] = [this.diagramSvgView, this.default];

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
        for (let svg of gongsvgsFrontRepo.getFrontArray<gongsvg.SVG>(gongsvg.SVG.GONGSTRUCT_NAME)) {
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