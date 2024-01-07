import { NgModule } from '@angular/core';

import { MatButtonModule } from '@angular/material/button';

import { MatIconModule } from '@angular/material/icon';


import { GongsvgspecificComponent } from './gongsvgspecific.component';


import { CommonModule } from '@angular/common';

import { GongsvgModule } from 'gongsvg';

import { LinkSegmentsPipe } from './link-segments.pipe';
import { GongsvgDiagrammingComponent } from './gongsvg-diagramming/gongsvg-diagramming';



@NgModule({
  declarations: [

    LinkSegmentsPipe,
    GongsvgDiagrammingComponent
  ],
  imports: [
    CommonModule,

    MatIconModule,
    MatButtonModule,

    GongsvgModule
  ],
  exports: [
    GongsvgDiagrammingComponent,
  ]
})
export class GongsvgspecificModule { }
