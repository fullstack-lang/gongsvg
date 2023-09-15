import { Component, OnInit, ViewChild } from '@angular/core';

import * as gongsvg from 'gongsvg'
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'SVG Data/Model'
  svg = 'SVG rendering'
  view = this.svg

  views: string[] = [this.svg, this.default];

  GONG__MODEL__StacksPath = "github.com/fullstack-lang/gongsvg/go/models"
  GONG__DATA__StackPath = gongsvg.StackName.StackNameDefault

  loading = true

  constructor(
  ) {

  }

  ngOnInit(): void {
    this.loading = false
  }
}
