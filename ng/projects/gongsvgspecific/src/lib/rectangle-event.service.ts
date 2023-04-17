import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

export type Coordinate = [number, number]

@Injectable({
  providedIn: 'root'
})
export class RectangleEventService {

  private mouseRectAltKeyDownEventSource = new Subject<number>();
  mouseRectAltKeyDownEvent$ = this.mouseRectAltKeyDownEventSource.asObservable();
  emitRectAltKeyMouseDownEvent(rectangleID: number) {
    this.mouseRectAltKeyDownEventSource.next(rectangleID);
  }

  private mouseRectAltKeyMouseUpEventSource = new Subject<number>();
  mouseRectAltKeyMouseUpEvent$ = this.mouseRectAltKeyMouseUpEventSource.asObservable();
  emitMouseRectAltKeyMouseUpEvent(rectangleID: number) {
    this.mouseRectAltKeyMouseUpEventSource.next(rectangleID);
  }

  private mouseRectAltKeyDragEventSource = new Subject<Coordinate>()
  mouseRectAltKeyDragEvent$ = this.mouseRectAltKeyDragEventSource.asObservable()
  emitRectAltKeyMouseDragEvent(coordinate: Coordinate) {
    this.mouseRectAltKeyDragEventSource.next(coordinate)
  }
}