import * as gongsvg from 'gongsvg'; // Replace 'gongsvg' with the correct module name
import { drawLinePointPoint, drawLinePointRect } from './draw.line';

export type ConnectorParams = {
    StartRect: gongsvg.RectDB
    EndRect: gongsvg.RectDB
    StartDirection: gongsvg.DirectionType
    EndDirection: gongsvg.DirectionType
    StartRatio: number
    EndRatio: number
    CornerOffsetRatio: number
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
        EndRatio: EndRatio,
        CornerOffsetRatio: CornerOffsetRatio } = params;
    const segments: gongsvg.PointDB[][] = []

    if (StartDirection === gongsvg.DirectionType.DIRECTION_HORIZONTAL &&
        EndDirection === gongsvg.DirectionType.DIRECTION_VERTICAL) {

        const startY = StartRect.Y + StartRatio * StartRect.Height
        const c1_X = EndRect.X + EndRatio * EndRect.Width
        const c1_Y = startY
        const c1 = createPoint(c1_X, c1_Y)

        const firstSegment = drawLinePointRect(c1, StartRect, StartDirection)
        const secondSegment = drawLinePointRect(c1, EndRect, EndDirection)


        segments.push(firstSegment, secondSegment)
    }

    if (StartDirection === gongsvg.DirectionType.DIRECTION_VERTICAL &&
        EndDirection === gongsvg.DirectionType.DIRECTION_HORIZONTAL) {

        const startX = StartRect.X + StartRatio * StartRect.Width
        const c1_X = startX
        const c1_Y = EndRect.Y + EndRatio * EndRect.Height
        const c1 = createPoint(c1_X, c1_Y)

        const firstSegment = drawLinePointRect(c1, StartRect, StartDirection)
        const secondSegment = drawLinePointRect(c1, EndRect, EndDirection)


        segments.push(firstSegment, secondSegment)
    }

    if (StartDirection === gongsvg.DirectionType.DIRECTION_HORIZONTAL &&
        EndDirection === gongsvg.DirectionType.DIRECTION_HORIZONTAL) {

        const c1_X = StartRect.X + CornerOffsetRatio * StartRect.Width
        const c1_Y = StartRect.Y + StartRatio * StartRect.Height

        const c1 = createPoint(c1_X, c1_Y)

        const c2_X = c1_X
        const c2_Y = EndRect.Y + EndRatio * EndRect.Height

        const c2 = createPoint(c2_X, c2_Y)

        const firstSegment = drawLinePointRect(c1, StartRect, StartDirection)
        const secondSegment = drawLinePointPoint(c1, c2)
        const thirdSegment = drawLinePointRect(c2, EndRect, EndDirection)


        segments.push(firstSegment, secondSegment, thirdSegment)
    }

    if (StartDirection === gongsvg.DirectionType.DIRECTION_VERTICAL &&
        EndDirection === gongsvg.DirectionType.DIRECTION_VERTICAL) {

        const c1_X = StartRect.X + StartRatio * StartRect.Width
        const c1_Y = StartRect.Y + CornerOffsetRatio * StartRect.Height

        const c1 = createPoint(c1_X, c1_Y)

        const c2_X = EndRect.X + EndRatio * EndRect.Width
        const c2_Y = c1_Y

        const c2 = createPoint(c2_X, c2_Y)

        const firstSegment = drawLinePointRect(c1, StartRect, StartDirection)
        const secondSegment = drawLinePointPoint(c1, c2)
        const thirdSegment = drawLinePointRect(c2, EndRect, EndDirection)


        segments.push(firstSegment, secondSegment, thirdSegment)
    }

    return segments;
}