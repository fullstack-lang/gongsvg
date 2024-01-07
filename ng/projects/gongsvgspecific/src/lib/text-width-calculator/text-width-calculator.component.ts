import { AfterViewInit, Component, ElementRef, ViewChild } from '@angular/core';

@Component({
  selector: 'lib-text-width-calculator',
  templateUrl: './text-width-calculator.component.html',
  styleUrls: ['./text-width-calculator.component.css']
})
export class TextWidthCalculatorComponent implements AfterViewInit {
  @ViewChild('measureElement') measureElement!: ElementRef;

  ngAfterViewInit() {
    console.log("measure text width", "ngAfterViewInit")
    // Use this method to measure the text width
  }

  measureTextWidth(text: string): number {
    // console.log("measure text width")
    const element = this.measureElement.nativeElement;
    element.textContent = text;
    return element.offsetWidth
  }

  measureTextHeight(text: string): number {
    console.log("measure text height")
    const element = this.measureElement.nativeElement;
    element.textContent = text;
    return element.offsetHeight
  }
}
