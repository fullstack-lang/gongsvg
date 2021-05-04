import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'lib-polygone',
  templateUrl: './polygone.component.svg',
  styleUrls: ['./polygone.component.css']
})
export class PolygoneComponent implements OnInit {

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
