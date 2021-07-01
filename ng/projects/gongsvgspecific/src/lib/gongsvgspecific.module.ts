import { NgModule } from '@angular/core';
import { GongsvgspecificComponent } from './gongsvgspecific.component';
import { SvgComponent } from './svg/svg.component';
import { BrowserModule } from '@angular/platform-browser';




@NgModule({
  declarations: [
    GongsvgspecificComponent,
    SvgComponent,
  ],
  imports: [
    BrowserModule
  ],
  exports: [
    GongsvgspecificComponent,
    SvgComponent,
  ]
})
export class GongsvgspecificModule { }
