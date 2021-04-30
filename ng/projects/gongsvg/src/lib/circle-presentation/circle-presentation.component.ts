import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { CircleDB } from '../circle-db'
import { CircleService } from '../circle.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface circleDummyElement {
}

const ELEMENT_DATA: circleDummyElement[] = [
];

@Component({
	selector: 'app-circle-presentation',
	templateUrl: './circle-presentation.component.html',
	styleUrls: ['./circle-presentation.component.css'],
})
export class CirclePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	circle: CircleDB;
 
	constructor(
		private circleService: CircleService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getCircle();

		// observable for changes in 
		this.circleService.CircleServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getCircle()
				}
			}
		)
	}

	getCircle(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.circleService.getCircle(id)
			.subscribe(
				circle => {
					this.circle = circle

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
				editor: ["circle-detail", ID]
			}
		}]);
	}
}
