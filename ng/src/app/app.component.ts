import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongsvg from 'gongsvg'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {


  GONG__StackPath="github.com/fullstack-lang/gongsvg/go/models"

  view = 'View'
  default = 'View'
  data = 'Data'
  model = 'Model'

  loading = true

  views: string[] = [this.default, this.data, this.model];

  // variable that enables pooling of selected gongstruct
  obsTimer: Observable<number> = timer(1000, 1000)
  lastSelectionDate: string = ''

  constructor(private gongdocGongStructShapeService: gongdoc.GongStructShapeService,
    private gongstructSelectionService: gongsvg.GongstructSelectionService
  ) {

  }

  ngOnInit(): void {
    this.loading = false
  }
}
