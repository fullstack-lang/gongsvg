import { Component, OnInit, ViewChild } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'SVG Data/Model'
  svg = 'SVG rendering'
  view = this.svg

  views: string[] = [this.svg, this.default];

  GONG__DATAMODEL_StackPath = "github.com/fullstack-lang/gongsvg/go/models"
  GONG__StackPath = "svg"

  loading = true

  constructor(
  ) {

  }

  ngOnInit(): void {
    this.loading = false
  }
}
