import { Component, OnInit, Input  } from '@angular/core';

@Component({
  selector: 'lib-text',
  templateUrl: './text.component.svg',
  styleUrls: ['./text.component.css']
})
export class TextComponent implements OnInit {

  @Input() X: number
  @Input() Y: number
  @Input() Color: string
  @Input() Content: string

  constructor() { }

  ngOnInit(): void {
  }

}
