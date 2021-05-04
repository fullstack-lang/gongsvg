import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';

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


@NgModule({
	declarations: [
		// insertion point for declarations 
		CirclesTableComponent,
		CircleDetailComponent,
		CirclePresentationComponent,

		EllipsesTableComponent,
		EllipseDetailComponent,
		EllipsePresentationComponent,

		LinesTableComponent,
		LineDetailComponent,
		LinePresentationComponent,

		PathsTableComponent,
		PathDetailComponent,
		PathPresentationComponent,

		PolygonesTableComponent,
		PolygoneDetailComponent,
		PolygonePresentationComponent,

		PolylinesTableComponent,
		PolylineDetailComponent,
		PolylinePresentationComponent,

		RectsTableComponent,
		RectDetailComponent,
		RectPresentationComponent,

		SVGsTableComponent,
		SVGDetailComponent,
		SVGPresentationComponent,

		TextsTableComponent,
		TextDetailComponent,
		TextPresentationComponent,


		SplitterComponent,
		SidebarComponent
	],
	imports: [
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
		RouterModule,

		AppRoutingModule,

		MatSliderModule,
		MatSelectModule,
		MatFormFieldModule,
		MatInputModule,
		MatDatepickerModule,
		MatTableModule,
		MatSortModule,
		MatPaginatorModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		CirclesTableComponent,
		CircleDetailComponent,
		CirclePresentationComponent,

		EllipsesTableComponent,
		EllipseDetailComponent,
		EllipsePresentationComponent,

		LinesTableComponent,
		LineDetailComponent,
		LinePresentationComponent,

		PathsTableComponent,
		PathDetailComponent,
		PathPresentationComponent,

		PolygonesTableComponent,
		PolygoneDetailComponent,
		PolygonePresentationComponent,

		PolylinesTableComponent,
		PolylineDetailComponent,
		PolylinePresentationComponent,

		RectsTableComponent,
		RectDetailComponent,
		RectPresentationComponent,

		SVGsTableComponent,
		SVGDetailComponent,
		SVGPresentationComponent,

		TextsTableComponent,
		TextDetailComponent,
		TextPresentationComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class GongsvgModule { }
