import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { SVGsTableComponent } from './svgs-table/svgs-table.component'
import { SVGDetailComponent } from './svg-detail/svg-detail.component'
import { SVGPresentationComponent } from './svg-presentation/svg-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongsvg_go-svgs', component: SVGsTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-svg-adder', component: SVGDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-svg-adder/:id/:originStruct/:originStructFieldName', component: SVGDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-svg-detail/:id', component: SVGDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-svg-presentation/:id', component: SVGPresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-svg-presentation-special/:id', component: SVGPresentationComponent, outlet: 'github_com_fullstack_lang_gongsvg_gosvgpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
