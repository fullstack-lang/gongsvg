
import * as gongsvg from 'gongsvg'; // Replace 'gongsvg' with the correct module name
import { Segment, createPoint, createSegment } from './draw.segments';

export function drawLinePointRect(
    point: gongsvg.PointDB,
    rect: gongsvg.RectDB,
    direction: gongsvg.OrientationType,
    number: number):
    Segment {
    let segment = createSegment(point, createPoint(0, 0), direction, number)

    if (direction === gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {
        if (point.X <= rect.X + rect.Width / 2) {
            const endPoint = createPoint(rect.X, point.Y)
            segment.EndPoint = endPoint
        }
        else {
            const endPoint = createPoint(rect.X + rect.Width, point.Y)
            segment.EndPoint = endPoint
        }
    }
    if (direction === gongsvg.OrientationType.ORIENTATION_VERTICAL) {
        if (point.Y <= rect.Y + rect.Height / 2) {
            const endPoint = createPoint(point.X, rect.Y)
            segment.EndPoint = endPoint
            return segment
        }
        else {
            const endPoint = createPoint(point.X, rect.Y + rect.Height)
            segment.EndPoint = endPoint
            return segment
        }
    }
    return segment
}