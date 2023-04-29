import * as gongsvg from 'gongsvg'; // Replace 'gongsvg' with the correct module name
import { drawLine } from './draw.line';

export type ConnectorParams = {
    StartRect: gongsvg.RectDB
    EndRect: gongsvg.RectDB
    StartDirection: gongsvg.DirectionType
    EndDirection: gongsvg.DirectionType
    StartRatio: number
    EndRatio: number
};

export function createPoint(x: number, y: number): gongsvg.PointDB {

    let point = new gongsvg.PointDB
    point.X = x
    point.Y = y
    return point
}

export function drawConnector(params: ConnectorParams): gongsvg.PointDB[][] {
    const {
        StartRect: StartRect,
        EndRect: EndRect,
        StartDirection: StartDirection,
        EndDirection: EndDirection,
        StartRatio: StartRatio,
        EndRatio: EndRatio } = params;
    const segments: gongsvg.PointDB[][] = []

    if (StartDirection === gongsvg.DirectionType.DIRECTION_HORIZONTAL &&
        EndDirection === gongsvg.DirectionType.DIRECTION_VERTICAL) {

        const startX = StartRect.X
        const startY = StartRect.Y + StartRatio * StartRect.Height
        const c1_X = EndRect.X + EndRatio * EndRect.Width
        const c1_Y = startY
        const c1 = createPoint(c1_X, c1_Y)

        const firstSegment = drawLine(c1, StartRect, StartDirection)
        const secondSegment = drawLine(c1, EndRect, EndDirection)


        segments.push(firstSegment, secondSegment)
    }

    if (StartDirection === gongsvg.DirectionType.DIRECTION_VERTICAL &&
        EndDirection === gongsvg.DirectionType.DIRECTION_HORIZONTAL) {

        const startX = StartRect.X + StartRatio * StartRect.Width
        const startY = StartRect.Y
        const c1_X = startX
        const c1_Y = EndRect.Y + EndRatio * EndRect.Height
        const c1 = createPoint(c1_X, c1_Y)

        const firstSegment = drawLine(c1, StartRect, StartDirection)
        const secondSegment = drawLine(c1, EndRect, EndDirection)


        segments.push(firstSegment, secondSegment)
    }

    return segments;
}