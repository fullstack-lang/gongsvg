import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'lib-ellipse',
  templateUrl: './ellipse.component.svg',
  styleUrls: ['./ellipse.component.css']
})
export class EllipseComponent implements OnInit {

  @Input() CX: number
  @Input() CY: number
  @Input() RX: number
  @Input() RY: number

  @Input() Color: string
  @Input() Stroke: string
  @Input() StrokeWidth: string
  @Input() FillOpacity: number
  @Input() StrokeDashArray: string

  constructor() { }

  ngOnInit(): void {
  }

}
