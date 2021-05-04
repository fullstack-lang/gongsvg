import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'lib-polyline',
  templateUrl: './polyline.component.svg',
  styleUrls: ['./polyline.component.css']
})
export class PolylineComponent implements OnInit {

  @Input() Points: string

  @Input() Color: string
  @Input() Stroke: string
  @Input() StrokeWidth: string
  @Input() FillOpacity: number
  @Input() StrokeDashArray: string

  constructor() { }

  ngOnInit(): void {
  }

}
