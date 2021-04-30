import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'lib-circle',
  templateUrl: './circle.component.svg',
  styleUrls: ['./circle.component.css']
})
export class CircleComponent implements OnInit {

  @Input() CX: number
  @Input() CY: number
  @Input() Radius: number

  @Input() Color: string
  @Input() Stroke: string
  @Input() StrokeWidth: string
  @Input() FillOpacity: number
  @Input() StrokeDashArray: string

  constructor() { }

  ngOnInit(): void {
  }

}
