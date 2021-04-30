import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'lib-svg-rect',
  templateUrl: './svg-rect.component.svg',
  styleUrls: ['./svg-rect.component.css']
})
export class SvgRectComponent implements OnInit {

  @Input() X: number
  @Input() Y: number
  @Input() Width: number
  @Input() Height: number
  @Input() Color: string

  constructor() { }

  ngOnInit(): void {
    // console.log("X " + this.X)
    // console.log("Y " + this.Y)
    // console.log("Height " + this.Height)
    // console.log("Width " + this.Width)
    // console.log("Color " + this.Color)
  }

  fillColor = 'rgb(255, 0, 0)';
  Text = "Toto"
  width = 150
  height = 200
  changeColor() {
    const r = Math.floor(Math.random() * 256);
    const g = Math.floor(Math.random() * 256);
    const b = Math.floor(Math.random() * 256);
    this.fillColor = `rgb(${r}, ${g}, ${b})`;
  }
}
