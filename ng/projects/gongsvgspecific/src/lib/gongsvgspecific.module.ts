import { NgModule } from '@angular/core';
import { GongsvgspecificComponent } from './gongsvgspecific.component';
import { SvgComponent } from './svg/svg.component';
import { SvgRectComponent } from './svg-rect/svg-rect.component';
import { BrowserModule } from '@angular/platform-browser';
import { TextComponent } from './text/text.component';
import { CircleComponent } from './circle/circle.component';
import { LineComponent } from './line/line.component'



@NgModule({
  declarations: [
    GongsvgspecificComponent,
    SvgComponent,
    SvgRectComponent,
    TextComponent,
    CircleComponent,
    LineComponent
  ],
  imports: [
    BrowserModule
  ],
  exports: [
    GongsvgspecificComponent,
    SvgComponent,
    SvgRectComponent,
    TextComponent
  ]
})
export class GongsvgspecificModule { }
