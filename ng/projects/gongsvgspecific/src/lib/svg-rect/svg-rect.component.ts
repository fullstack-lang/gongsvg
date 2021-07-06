import { Component, OnInit, Input } from '@angular/core';

import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-svg-rect',
  templateUrl: './svg-rect.component.svg',
  styleUrls: ['./svg-rect.component.css']
})
export class SvgRectComponent implements OnInit {

  @Input() Rect: gongsvg.RectDB

  AttributeName: string
  Values: string
  Dur: string
  RepeatCount: string

  Animates = new Array<gongsvg.AnimateDB>()

  constructor() { }

  ngOnInit(): void {

    this.AttributeName = "rx"
    this.Values = "0;100;0"
    this.Dur = "3s"
    this.RepeatCount = "indefinite"

    let animate = new gongsvg.AnimateDB()
    animate.AttributeName ="x"
    animate.Dur = "3s"
    animate.Values = "0;100;0"
    animate.RepeatCount= "indefinite"

    this.Animates.push(animate)

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
