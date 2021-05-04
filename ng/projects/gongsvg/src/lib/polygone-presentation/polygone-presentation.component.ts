import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { PolygoneDB } from '../polygone-db'
import { PolygoneService } from '../polygone.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface polygoneDummyElement {
}

const ELEMENT_DATA: polygoneDummyElement[] = [
];

@Component({
	selector: 'app-polygone-presentation',
	templateUrl: './polygone-presentation.component.html',
	styleUrls: ['./polygone-presentation.component.css'],
})
export class PolygonePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	polygone: PolygoneDB;
 
	constructor(
		private polygoneService: PolygoneService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getPolygone();

		// observable for changes in 
		this.polygoneService.PolygoneServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getPolygone()
				}
			}
		)
	}

	getPolygone(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.polygoneService.getPolygone(id)
			.subscribe(
				polygone => {
					this.polygone = polygone

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
				editor: ["polygone-detail", ID]
			}
		}]);
	}
}
