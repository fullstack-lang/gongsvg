import { Component, ElementRef, Input, ViewChild } from '@angular/core';

@Component({
  selector: 'lib-material-svg',
  templateUrl: './material-svg.component.html',
  styleUrls: ['./material-svg.component.css']
})
export class MaterialSvgComponent {

  @Input() GONG__StackPath: string = ""
  @ViewChild('drawingArea') drawingArea: ElementRef<HTMLDivElement> | undefined

}
