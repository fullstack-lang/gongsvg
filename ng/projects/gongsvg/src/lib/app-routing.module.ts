import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { CirclesTableComponent } from './circles-table/circles-table.component'
import { CircleDetailComponent } from './circle-detail/circle-detail.component'
import { CirclePresentationComponent } from './circle-presentation/circle-presentation.component'

import { EllipsesTableComponent } from './ellipses-table/ellipses-table.component'
import { EllipseDetailComponent } from './ellipse-detail/ellipse-detail.component'
import { EllipsePresentationComponent } from './ellipse-presentation/ellipse-presentation.component'

import { LinesTableComponent } from './lines-table/lines-table.component'
import { LineDetailComponent } from './line-detail/line-detail.component'
import { LinePresentationComponent } from './line-presentation/line-presentation.component'

import { PathsTableComponent } from './paths-table/paths-table.component'
import { PathDetailComponent } from './path-detail/path-detail.component'
import { PathPresentationComponent } from './path-presentation/path-presentation.component'

import { PolygonesTableComponent } from './polygones-table/polygones-table.component'
import { PolygoneDetailComponent } from './polygone-detail/polygone-detail.component'
import { PolygonePresentationComponent } from './polygone-presentation/polygone-presentation.component'

import { PolylinesTableComponent } from './polylines-table/polylines-table.component'
import { PolylineDetailComponent } from './polyline-detail/polyline-detail.component'
import { PolylinePresentationComponent } from './polyline-presentation/polyline-presentation.component'

import { RectsTableComponent } from './rects-table/rects-table.component'
import { RectDetailComponent } from './rect-detail/rect-detail.component'
import { RectPresentationComponent } from './rect-presentation/rect-presentation.component'

import { SVGsTableComponent } from './svgs-table/svgs-table.component'
import { SVGDetailComponent } from './svg-detail/svg-detail.component'
import { SVGPresentationComponent } from './svg-presentation/svg-presentation.component'

import { TextsTableComponent } from './texts-table/texts-table.component'
import { TextDetailComponent } from './text-detail/text-detail.component'
import { TextPresentationComponent } from './text-presentation/text-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'circles', component: CirclesTableComponent, outlet: 'table' },
	{ path: 'circle-adder', component: CircleDetailComponent, outlet: 'editor' },
	{ path: 'circle-adder/:id/:association', component: CircleDetailComponent, outlet: 'editor' },
	{ path: 'circle-detail/:id', component: CircleDetailComponent, outlet: 'editor' },
	{ path: 'circle-presentation/:id', component: CirclePresentationComponent, outlet: 'presentation' },
	{ path: 'circle-presentation-special/:id', component: CirclePresentationComponent, outlet: 'circlepres' },

	{ path: 'ellipses', component: EllipsesTableComponent, outlet: 'table' },
	{ path: 'ellipse-adder', component: EllipseDetailComponent, outlet: 'editor' },
	{ path: 'ellipse-adder/:id/:association', component: EllipseDetailComponent, outlet: 'editor' },
	{ path: 'ellipse-detail/:id', component: EllipseDetailComponent, outlet: 'editor' },
	{ path: 'ellipse-presentation/:id', component: EllipsePresentationComponent, outlet: 'presentation' },
	{ path: 'ellipse-presentation-special/:id', component: EllipsePresentationComponent, outlet: 'ellipsepres' },

	{ path: 'lines', component: LinesTableComponent, outlet: 'table' },
	{ path: 'line-adder', component: LineDetailComponent, outlet: 'editor' },
	{ path: 'line-adder/:id/:association', component: LineDetailComponent, outlet: 'editor' },
	{ path: 'line-detail/:id', component: LineDetailComponent, outlet: 'editor' },
	{ path: 'line-presentation/:id', component: LinePresentationComponent, outlet: 'presentation' },
	{ path: 'line-presentation-special/:id', component: LinePresentationComponent, outlet: 'linepres' },

	{ path: 'paths', component: PathsTableComponent, outlet: 'table' },
	{ path: 'path-adder', component: PathDetailComponent, outlet: 'editor' },
	{ path: 'path-adder/:id/:association', component: PathDetailComponent, outlet: 'editor' },
	{ path: 'path-detail/:id', component: PathDetailComponent, outlet: 'editor' },
	{ path: 'path-presentation/:id', component: PathPresentationComponent, outlet: 'presentation' },
	{ path: 'path-presentation-special/:id', component: PathPresentationComponent, outlet: 'pathpres' },

	{ path: 'polygones', component: PolygonesTableComponent, outlet: 'table' },
	{ path: 'polygone-adder', component: PolygoneDetailComponent, outlet: 'editor' },
	{ path: 'polygone-adder/:id/:association', component: PolygoneDetailComponent, outlet: 'editor' },
	{ path: 'polygone-detail/:id', component: PolygoneDetailComponent, outlet: 'editor' },
	{ path: 'polygone-presentation/:id', component: PolygonePresentationComponent, outlet: 'presentation' },
	{ path: 'polygone-presentation-special/:id', component: PolygonePresentationComponent, outlet: 'polygonepres' },

	{ path: 'polylines', component: PolylinesTableComponent, outlet: 'table' },
	{ path: 'polyline-adder', component: PolylineDetailComponent, outlet: 'editor' },
	{ path: 'polyline-adder/:id/:association', component: PolylineDetailComponent, outlet: 'editor' },
	{ path: 'polyline-detail/:id', component: PolylineDetailComponent, outlet: 'editor' },
	{ path: 'polyline-presentation/:id', component: PolylinePresentationComponent, outlet: 'presentation' },
	{ path: 'polyline-presentation-special/:id', component: PolylinePresentationComponent, outlet: 'polylinepres' },

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

	{ path: 'texts', component: TextsTableComponent, outlet: 'table' },
	{ path: 'text-adder', component: TextDetailComponent, outlet: 'editor' },
	{ path: 'text-adder/:id/:association', component: TextDetailComponent, outlet: 'editor' },
	{ path: 'text-detail/:id', component: TextDetailComponent, outlet: 'editor' },
	{ path: 'text-presentation/:id', component: TextPresentationComponent, outlet: 'presentation' },
	{ path: 'text-presentation-special/:id', component: TextPresentationComponent, outlet: 'textpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
