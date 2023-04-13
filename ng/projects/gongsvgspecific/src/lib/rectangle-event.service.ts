import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

export type Coordinate = [number, number]

@Injectable({
  providedIn: 'root'
})
export class RectangleEventService {
  private mouseDownEventSource = new Subject<number>();

  mouseDownEvent$ = this.mouseDownEventSource.asObservable();

  emitMouseDownEvent(rectangleID: number) {
    this.mouseDownEventSource.next(rectangleID);
  }

  private mouseUpEventSource = new Subject<number>();

  mouseUpEvent$ = this.mouseUpEventSource.asObservable();

  emitMouseUpEvent(rectangleID: number) {
    this.mouseUpEventSource.next(rectangleID);
  }

  private mouseDragEventSource = new Subject<Coordinate>()

  mouseDragEvent$ = this.mouseDragEventSource.asObservable()

  emitMouseDragEvent(coordinate: Coordinate) {
    this.mouseDragEventSource.next(coordinate)
  }
}