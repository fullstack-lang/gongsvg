import { Component, Input, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';
import { Coordinate, RectangleEventService } from '../rectangle-event.service';

import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-layer',
  templateUrl: './layer.component.html',
  styleUrls: ['./layer.component.css']
})
export class LayerComponent {

  @Input() GONG__StackPath: string = ""
  @Input() Layer: gongsvg.LayerDB = new gongsvg.LayerDB()

}
