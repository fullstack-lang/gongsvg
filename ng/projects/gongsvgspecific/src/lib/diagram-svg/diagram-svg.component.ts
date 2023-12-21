import { Component, Input } from '@angular/core';

@Component({
  selector: 'lib-diagram-svg',
  templateUrl: './diagram-svg.component.html',
  styleUrls: ['./diagram-svg.component.css']
})
export class DiagramSvgComponent {

  @Input() GONG__StackPath: string = ""

}
