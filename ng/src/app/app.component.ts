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

  default = 'SVG Data/Model'
  svg = 'SVG rendering'
  view = this.svg

  views: string[] = [this.svg, this.default];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackType = "github.com/fullstack-lang/gongsvg/go/models"
  StackName = gongsvg.StackName.StackNameDefault
  loading = true
  constructor(
  ) {

  }

  ngOnInit(): void {
    this.loading = false
  }
}
