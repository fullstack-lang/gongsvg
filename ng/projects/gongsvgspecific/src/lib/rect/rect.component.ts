import { Component, OnInit, Input } from '@angular/core';

import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-rect',
  templateUrl: './rect.component.svg',
  styleUrls: ['./rect.component.css']
})
export class RectComponent implements OnInit {

  @Input() Rect?: gongsvg.RectDB

  // In your component
  anchorRadius = 5; // Adjust this value according to your desired anchor size
  anchorFillColor = 'blue'; // Choose your desired anchor fill color


  constructor() { }

  ngOnInit(): void {
  }
}
