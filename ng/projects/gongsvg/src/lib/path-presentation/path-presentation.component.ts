import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { PathDB } from '../path-db'
import { PathService } from '../path.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface pathDummyElement {
}

const ELEMENT_DATA: pathDummyElement[] = [
];

@Component({
	selector: 'app-path-presentation',
	templateUrl: './path-presentation.component.html',
	styleUrls: ['./path-presentation.component.css'],
})
export class PathPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	path: PathDB;
 
	constructor(
		private pathService: PathService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getPath();

		// observable for changes in 
		this.pathService.PathServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getPath()
				}
			}
		)
	}

	getPath(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.pathService.getPath(id)
			.subscribe(
				path => {
					this.path = path

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
				editor: ["path-detail", ID]
			}
		}]);
	}
}
