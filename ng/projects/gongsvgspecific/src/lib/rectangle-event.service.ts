import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

export type Coordinate = [number, number]

// Define the interface for the event object
interface RectMouseEvent {
  rectangleID: number;
  Coordinate: [number, number]
}

@Injectable({
  providedIn: 'root'
})
export class RectangleEventService {

  private mouseRectMouseDragEventSource = new Subject<RectMouseEvent>();
  mouseRectMouseDragEvent$ = this.mouseRectMouseDragEventSource.asObservable();
  emitRectMouseDragEvent(rectMouseEvent: RectMouseEvent) {
    // console.log('RectangleEventService, rect mouse drag event, rectangle', rectangleID)
    this.mouseRectMouseDragEventSource.next(rectMouseEvent);
  }

  private mouseRectMouseUpEventSource = new Subject<number>();
  mouseRectMouseUpEvent$ = this.mouseRectMouseUpEventSource.asObservable();
  emitRectMouseDownEvent(rectangleID: number) {
    // console.log('RectangleEventService, rect mouse down event, rectangle', rectangleID)
    this.mouseRectMouseUpEventSource.next(rectangleID);
  }

  private mouseRectAltKeyMouseDownEventSource = new Subject<RectMouseEvent>();
  mouseRectAltKeyMouseDownEvent$ = this.mouseRectAltKeyMouseDownEventSource.asObservable();
  emitRectAltKeyMouseDownEvent(rectangleID: number, coordinate: [number, number]) {
    this.mouseRectAltKeyMouseDownEventSource.next({ rectangleID, Coordinate: coordinate });
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