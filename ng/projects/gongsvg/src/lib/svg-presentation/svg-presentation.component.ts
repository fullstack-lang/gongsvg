import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { SVGDB } from '../svg-db'
import { SVGService } from '../svg.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface svgDummyElement {
}

const ELEMENT_DATA: svgDummyElement[] = [
];

@Component({
	selector: 'app-svg-presentation',
	templateUrl: './svg-presentation.component.html',
	styleUrls: ['./svg-presentation.component.css'],
})
export class SVGPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	svg: SVGDB;
 
	constructor(
		private svgService: SVGService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getSVG();

		// observable for changes in 
		this.svgService.SVGServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getSVG()
				}
			}
		)
	}

	getSVG(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.svgService.getSVG(id)
			.subscribe(
				svg => {
					this.svg = svg

					// insertion point for recovery of durations

				}
			);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				presentation: [structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				editor: ["svg-detail", ID]
			}
		}]);
	}
}
