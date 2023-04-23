import { Component, Input, OnInit } from '@angular/core';
import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-link',
  templateUrl: './link.component.svg',
  styleUrls: ['./link.component.css']
})
export class LinkComponent implements OnInit {

  @Input() Link?: gongsvg.LinkDB
  @Input() GONG__StackPath: string = ""

  ngOnInit(): void {
    console.log("LinkComponent init: ", this.Link?.Name)

  }

}
