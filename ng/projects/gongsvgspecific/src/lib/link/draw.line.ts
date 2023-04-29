
import * as gongsvg from 'gongsvg'; // Replace 'gongsvg' with the correct module name
import { createPoint } from './compute.floating.connectors.2';

export function drawLine(point: gongsvg.PointDB, rect: gongsvg.RectDB, direction: gongsvg.DirectionType): gongsvg.PointDB[] {
    const line: gongsvg.PointDB[] = []

    if (direction === gongsvg.DirectionType.DIRECTION_HORIZONTAL) {
        if (point.X <= rect.X + rect.Width / 2) {
            const endPoint = createPoint(rect.X, point.Y)
            line.push(point, endPoint)
        }
        else {
            const endPoint = createPoint(rect.X + rect.Width, point.Y)
            line.push(point, endPoint)
        }
        line.push(point, point)

    }
    if (direction === gongsvg.DirectionType.DIRECTION_VERTICAL) {
        if (point.Y <= rect.Y + rect.Height / 2) {
            const endPoint = createPoint(point.X, rect.Y)
            line.push(point, endPoint)
            return line
        }
        else {
            const endPoint = createPoint(point.X, rect.Y + rect.Height)
            line.push(point, endPoint)
            return line
        }
    }


    return line
}