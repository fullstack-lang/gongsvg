import * as gongsvg from 'gongsvg'; // Replace 'gongsvg' with the correct module name
import { drawLinePointRect } from './draw.line';

export type ConnectorParams = {
    StartRect: gongsvg.RectDB
    EndRect: gongsvg.RectDB
    StartDirection: gongsvg.OrientationType
    EndDirection: gongsvg.OrientationType
    StartRatio: number
    EndRatio: number
    CornerOffsetRatio: number
}

export type Segment = {
    StartPoint: gongsvg.PointDB,
    EndPoint: gongsvg.PointDB
    Orientation: gongsvg.OrientationType
    Number: number
}
export function createSegment(
    start: gongsvg.PointDB,
    end: gongsvg.PointDB,
    orientaion: gongsvg.OrientationType,
    number: number): Segment {
    return { StartPoint: start, EndPoint: end, Orientation: orientaion, Number: number }
}

export function createPoint(x: number, y: number): gongsvg.PointDB {

    let point = new gongsvg.PointDB
    point.X = x
    point.Y = y
    return point
}

export function drawSegments(params: ConnectorParams): Segment[] {
    const {
        StartRect: StartRect,
        EndRect: EndRect,
        StartDirection: StartDirection,
        EndDirection: EndDirection,
        StartRatio: StartRatio,
        EndRatio: EndRatio,
        CornerOffsetRatio: CornerOffsetRatio } = params;
    const segments: Segment[] = []

    if (StartDirection === gongsvg.OrientationType.ORIENTATION_HORIZONTAL &&
        EndDirection === gongsvg.OrientationType.ORIENTATION_VERTICAL) {

        const startY = StartRect.Y + StartRatio * StartRect.Height
        const c1_X = EndRect.X + EndRatio * EndRect.Width
        const c1_Y = startY
        const c1 = createPoint(c1_X, c1_Y)

        const firstSegment = drawLinePointRect(c1, StartRect, StartDirection, 0)
        const secondSegment = drawLinePointRect(c1, EndRect, EndDirection, 1)



        segments.push(firstSegment, secondSegment)
    }

    if (StartDirection === gongsvg.OrientationType.ORIENTATION_VERTICAL &&
        EndDirection === gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {

        const startX = StartRect.X + StartRatio * StartRect.Width
        const c1_X = startX
        const c1_Y = EndRect.Y + EndRatio * EndRect.Height
        const c1 = createPoint(c1_X, c1_Y)

        const firstSegment = drawLinePointRect(c1, StartRect, StartDirection, 0)
        const secondSegment = drawLinePointRect(c1, EndRect, EndDirection, 1)


        segments.push(firstSegment, secondSegment)
    }

    if (StartDirection === gongsvg.OrientationType.ORIENTATION_HORIZONTAL &&
        EndDirection === gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {

        const c1_X = StartRect.X + CornerOffsetRatio * StartRect.Width
        const c1_Y = StartRect.Y + StartRatio * StartRect.Height

        const c1 = createPoint(c1_X, c1_Y)

        const c2_X = c1_X
        const c2_Y = EndRect.Y + EndRatio * EndRect.Height

        const c2 = createPoint(c2_X, c2_Y)

        const firstSegment = drawLinePointRect(c1, StartRect, StartDirection, 0)
        const secondSegment = createSegment(c1, c2, gongsvg.OrientationType.ORIENTATION_VERTICAL, 1)
        const thirdSegment = drawLinePointRect(c2, EndRect, EndDirection, 2)


        segments.push(firstSegment, secondSegment, thirdSegment)
    }

    if (StartDirection === gongsvg.OrientationType.ORIENTATION_VERTICAL &&
        EndDirection === gongsvg.OrientationType.ORIENTATION_VERTICAL) {

        const c1_X = StartRect.X + StartRatio * StartRect.Width
        const c1_Y = StartRect.Y + CornerOffsetRatio * StartRect.Height

        const c1 = createPoint(c1_X, c1_Y)

        const c2_X = EndRect.X + EndRatio * EndRect.Width
        const c2_Y = c1_Y

        const c2 = createPoint(c2_X, c2_Y)

        const firstSegment = drawLinePointRect(c1, StartRect, StartDirection, 0)
        const secondSegment = createSegment(c1, c2, gongsvg.OrientationType.ORIENTATION_HORIZONTAL, 1)
        const thirdSegment = drawLinePointRect(c2, EndRect, EndDirection, 2)


        segments.push(firstSegment, secondSegment, thirdSegment)
    }

    return segments;
}