import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { LineDB } from '../line-db'
import { LineService } from '../line.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface lineDummyElement {
}

const ELEMENT_DATA: lineDummyElement[] = [
];

@Component({
	selector: 'app-line-presentation',
	templateUrl: './line-presentation.component.html',
	styleUrls: ['./line-presentation.component.css'],
})
export class LinePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	line: LineDB;
 
	constructor(
		private lineService: LineService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getLine();

		// observable for changes in 
		this.lineService.LineServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getLine()
				}
			}
		)
	}

	getLine(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.lineService.getLine(id)
			.subscribe(
				line => {
					this.line = line

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
				editor: ["line-detail", ID]
			}
		}]);
	}
}
