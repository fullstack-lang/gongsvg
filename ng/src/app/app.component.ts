import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'SVG Data/Model'
  svg = 'SVG rendering'
  view = this.svg

  views: string[] = [this.svg, this.default];

  GONG__StackPath="github.com/fullstack-lang/gongsvg/go/models"

  loading = true

  constructor(
  ) {

  }

  ngOnInit(): void {
    this.loading = false
  }
}
