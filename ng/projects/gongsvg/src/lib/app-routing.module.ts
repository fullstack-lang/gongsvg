import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { RectsTableComponent } from './rects-table/rects-table.component'
import { RectDetailComponent } from './rect-detail/rect-detail.component'
import { RectPresentationComponent } from './rect-presentation/rect-presentation.component'

import { SVGsTableComponent } from './svgs-table/svgs-table.component'
import { SVGDetailComponent } from './svg-detail/svg-detail.component'
import { SVGPresentationComponent } from './svg-presentation/svg-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'rects', component: RectsTableComponent, outlet: 'table' },
	{ path: 'rect-adder', component: RectDetailComponent, outlet: 'editor' },
	{ path: 'rect-adder/:id/:association', component: RectDetailComponent, outlet: 'editor' },
	{ path: 'rect-detail/:id', component: RectDetailComponent, outlet: 'editor' },
	{ path: 'rect-presentation/:id', component: RectPresentationComponent, outlet: 'presentation' },
	{ path: 'rect-presentation-special/:id', component: RectPresentationComponent, outlet: 'rectpres' },

	{ path: 'svgs', component: SVGsTableComponent, outlet: 'table' },
	{ path: 'svg-adder', component: SVGDetailComponent, outlet: 'editor' },
	{ path: 'svg-adder/:id/:association', component: SVGDetailComponent, outlet: 'editor' },
	{ path: 'svg-detail/:id', component: SVGDetailComponent, outlet: 'editor' },
	{ path: 'svg-presentation/:id', component: SVGPresentationComponent, outlet: 'presentation' },
	{ path: 'svg-presentation-special/:id', component: SVGPresentationComponent, outlet: 'svgpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
