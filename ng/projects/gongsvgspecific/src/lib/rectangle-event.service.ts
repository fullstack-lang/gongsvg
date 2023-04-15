import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

export type Coordinate = [number, number]

@Injectable({
  providedIn: 'root'
})
export class RectangleEventService {
  private mouseRectShiftKeyDownEventSource = new Subject<number>();

  mouseRectShiftKeyDownEvent$ = this.mouseRectShiftKeyDownEventSource.asObservable();

  emitRectShiftKeyMouseDownEvent(rectangleID: number) {
    this.mouseRectShiftKeyDownEventSource.next(rectangleID);
  }

  private mouseRectShiftKeyMouseUpEventSource = new Subject<number>();

  mouseRectShiftKeyMouseUpEvent$ = this.mouseRectShiftKeyMouseUpEventSource.asObservable();

  emitMouseRectShiftKeyMouseUpEvent(rectangleID: number) {
    this.mouseRectShiftKeyMouseUpEventSource.next(rectangleID);
  }

  private mouseRectShiftKeyDragEventSource = new Subject<Coordinate>()

  mouseRectShiftKeyDragEvent$ = this.mouseRectShiftKeyDragEventSource.asObservable()

  emitRectShiftKeyMouseDragEvent(coordinate: Coordinate) {
    this.mouseRectShiftKeyDragEventSource.next(coordinate)
  }
}