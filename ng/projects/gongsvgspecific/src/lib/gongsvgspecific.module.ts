import { NgModule } from '@angular/core';
import { GongsvgspecificComponent } from './gongsvgspecific.component';
import { SvgComponent } from './svg/svg.component';
import { SvgRectComponent } from './svg-rect/svg-rect.component';
import { BrowserModule } from '@angular/platform-browser';
import { TextComponent } from './text/text.component'



@NgModule({
  declarations: [
    GongsvgspecificComponent,
    SvgComponent,
    SvgRectComponent,
    TextComponent
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
