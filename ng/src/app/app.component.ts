import { Component, OnInit, ViewChild } from '@angular/core';
import { AngularDragEndEventService } from 'projects/gongsvgspecific/src/lib/angular-drag-end-event.service';

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
    private angularDragEndEventService: AngularDragEndEventService,
  ) {

  }

  ngOnInit(): void {
    this.loading = false
  }

  onDragEnd(): void {
    console.log("angular split : on drag end")
    this.angularDragEndEventService.emitMouseUpEvent(0)
  }
}
