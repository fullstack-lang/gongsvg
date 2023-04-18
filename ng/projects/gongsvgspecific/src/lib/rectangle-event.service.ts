import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

export type Coordinate = [number, number]

@Injectable({
  providedIn: 'root'
})
export class RectangleEventService {

  private mouseRectMouseDownEventSource = new Subject<number>();
  mouseRectMouseDownEvent$ = this.mouseRectMouseDownEventSource.asObservable();
  emitRectMouseDownEvent(rectangleID: number) {
    // console.log('RectangleEventService, rect mouse down event, rectangle', rectangleID)
    this.mouseRectMouseDownEventSource.next(rectangleID);
  }

  private mouseRectAltKeyMouseDownEventSource = new Subject<number>();
  mouseRectAltKeyMouseDownEvent$ = this.mouseRectAltKeyMouseDownEventSource.asObservable();
  emitRectAltKeyMouseDownEvent(rectangleID: number) {
    this.mouseRectAltKeyMouseDownEventSource.next(rectangleID);
  }

  private mouseRectAltKeyMouseUpEventSource = new Subject<number>();
  mouseRectAltKeyMouseUpEvent$ = this.mouseRectAltKeyMouseUpEventSource.asObservable();
  emitMouseRectAltKeyMouseUpEvent(rectangleID: number) {
    this.mouseRectAltKeyMouseUpEventSource.next(rectangleID);
  }

  private mouseRectAltKeyMouseDragEventSource = new Subject<Coordinate>()
  mouseRectAltKeyMouseDragEvent$ = this.mouseRectAltKeyMouseDragEventSource.asObservable()
  emitRectAltKeyMouseDragEvent(coordinate: Coordinate) {
    this.mouseRectAltKeyMouseDragEventSource.next(coordinate)
  }
}