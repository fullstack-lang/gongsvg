import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'lib-svg-rect',
  templateUrl: './svg-rect.component.svg',
  styleUrls: ['./svg-rect.component.css']
})
export class SvgRectComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
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
