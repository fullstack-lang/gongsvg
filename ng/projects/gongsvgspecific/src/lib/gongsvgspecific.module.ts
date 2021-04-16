import { NgModule } from '@angular/core';
import { GongsvgspecificComponent } from './gongsvgspecific.component';
import { SvgComponent } from './svg/svg.component';
import { SvgRectComponent } from './svg-rect/svg-rect.component';
import { BrowserModule } from '@angular/platform-browser'



@NgModule({
  declarations: [
    GongsvgspecificComponent,
    SvgComponent,
    SvgRectComponent
  ],
  imports: [
    BrowserModule
  ],
  exports: [
    GongsvgspecificComponent,
    SvgComponent,
    SvgRectComponent
  ]
})
export class GongsvgspecificModule { }
