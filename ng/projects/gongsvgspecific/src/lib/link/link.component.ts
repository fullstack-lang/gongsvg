import { Component, Input } from '@angular/core';
import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-link',
  templateUrl: './link.component.svg',
  styleUrls: ['./link.component.css']
})
export class LinkComponent {

  @Input() Link?: gongsvg.LinkDB
  @Input() GONG__StackPath: string = ""

}
