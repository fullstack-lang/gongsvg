import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { EllipseDB } from '../ellipse-db'
import { EllipseService } from '../ellipse.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface ellipseDummyElement {
}

const ELEMENT_DATA: ellipseDummyElement[] = [
];

@Component({
	selector: 'app-ellipse-presentation',
	templateUrl: './ellipse-presentation.component.html',
	styleUrls: ['./ellipse-presentation.component.css'],
})
export class EllipsePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	ellipse: EllipseDB;
 
	constructor(
		private ellipseService: EllipseService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getEllipse();

		// observable for changes in 
		this.ellipseService.EllipseServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getEllipse()
				}
			}
		)
	}

	getEllipse(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.ellipseService.getEllipse(id)
			.subscribe(
				ellipse => {
					this.ellipse = ellipse

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
				editor: ["ellipse-detail", ID]
			}
		}]);
	}
}
