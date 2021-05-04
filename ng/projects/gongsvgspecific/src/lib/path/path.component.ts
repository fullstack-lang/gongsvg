import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'lib-path',
  templateUrl: './path.component.svg',
  styleUrls: ['./path.component.css']
})
export class PathComponent implements OnInit {

  @Input() Definition: string

  @Input() Color: string
  @Input() Stroke: string
  @Input() StrokeWidth: string
  @Input() FillOpacity: number
  @Input() StrokeDashArray: string
  constructor() { }

  ngOnInit(): void {
  }

}
