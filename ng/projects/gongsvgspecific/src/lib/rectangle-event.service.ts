import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

export type Coordinate = [number, number]

// Define the interface for the event object
interface RectMouseEvent {
  rectangleID: number;
  MousePosRelativeSVG: [number, number]
}

@Injectable({
  providedIn: 'root'
})
export class RectangleEventService {

  //
  // mouse events
  //
  private mouseRectMouseDownEventSource = new Subject<RectMouseEvent>();
  mouseRectMouseDownEvent$ = this.mouseRectMouseDownEventSource.asObservable();
  emitRectMouseDownEvent(rectMouseEvent: RectMouseEvent) {
    // console.log('RectangleEventService, rect mouse down event, rectangle', rectangleID)
    this.mouseRectMouseDownEventSource.next(rectMouseEvent);
  }


  private mouseRectMouseDragEventSource = new Subject<RectMouseEvent>();
  mouseRectMouseDragEvent$ = this.mouseRectMouseDragEventSource.asObservable();
  emitRectMouseDragEvent(rectMouseEvent: RectMouseEvent) {
    // console.log('RectangleEventService, rect mouse drag event, rectangle', rectangleID)
    this.mouseRectMouseDragEventSource.next(rectMouseEvent);
  }

  private mouseRectMouseUpEventSource = new Subject<number>();
  mouseRectMouseUpEvent$ = this.mouseRectMouseUpEventSource.asObservable();
  emitRectMouseUpEvent(rectangleID: number) {
    // console.log('RectangleEventService, rect mouse down event, rectangle', rectangleID)
    this.mouseRectMouseUpEventSource.next(rectangleID);
  }

  //
  // mouse ALT events
  //

  private mouseRectAltKeyMouseDownEventSource = new Subject<RectMouseEvent>();
  mouseRectAltKeyMouseDownEvent$ = this.mouseRectAltKeyMouseDownEventSource.asObservable();
  emitRectAltKeyMouseDownEvent(rectangleID: number, coordinate: [number, number]) {
    this.mouseRectAltKeyMouseDownEventSource.next({ rectangleID, MousePosRelativeSVG: coordinate });
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