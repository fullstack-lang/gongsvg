import * as gongsvg from 'gongsvg'; // Replace 'gongsvg' with the correct module name

export type ConnectorParams = {
    startRect: gongsvg.RectDB
    endRect: gongsvg.RectDB
    startDirection: gongsvg.DirectionType
    endDirection: gongsvg.DirectionType
    startRatio: number
    endRatio: number
};

function createPoint(x: number, y: number): gongsvg.PointDB {

    let point = new gongsvg.PointDB
    point.X = x
    point.Y = y
    return point
}

export function drawConnector(params: ConnectorParams): gongsvg.PointDB[][] {
    const {
        startRect,
        endRect,
        startDirection: startingDirection,
        endDirection: endDirection,
        startRatio: startRatio,
        endRatio: endRatio } = params;
    const segments: gongsvg.PointDB[][] = []

    if (startingDirection === gongsvg.DirectionType.DIRECTION_HORIZONTAL &&
        endDirection === gongsvg.DirectionType.DIRECTION_VERTICAL) {

        const startX = startRect.X
        const startY = startRect.Y + startRatio * startRect.Height
        const c1_X = endRect.X + endRatio * endRect.Width
        const c1_Y = startY
        const c1 = createPoint(c1_X, c1_Y)

        const firstSegment = drawLine(c1, startRect, startingDirection)

        segments.push(firstSegment)
    }

    return segments;
}

function drawLine(point: gongsvg.PointDB, rect: gongsvg.RectDB, direction: gongsvg.DirectionType): gongsvg.PointDB[] {
    const line: gongsvg.PointDB[] = []

    if (direction === gongsvg.DirectionType.DIRECTION_HORIZONTAL) {
        if (point.X <= rect.X) {
            const endPoint = createPoint(rect.X, point.Y)
            line.push(point, endPoint)
            return line
        }
        if (point.X >= rect.X + rect.Width) {
            const endPoint = createPoint(rect.X + rect.Width, point.Y)
            line.push(point, endPoint)
            return line
        }
        line.push(point, point)
    }

    return line
}