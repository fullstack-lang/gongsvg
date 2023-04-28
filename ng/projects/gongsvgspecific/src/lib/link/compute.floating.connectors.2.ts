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
        startDirection: startDirection,
        endDirection: endDirection,
        startRatio: startRatio,
        endRatio: endRatio } = params;
    const segments: gongsvg.PointDB[][] = []

    if (startDirection === gongsvg.DirectionType.DIRECTION_HORIZONTAL &&
        endDirection === gongsvg.DirectionType.DIRECTION_VERTICAL) {

        const startX = startRect.X
        const startY = startRect.Y + startRatio * startRect.Height
        const c1_X = endRect.X + endRatio * endRect.Width
        const c1_Y = startY
        const c1 = createPoint(c1_X, c1_Y)

        const firstSegment = drawLine(c1, startRect, startDirection)
        const secondSegment = drawLine(c1, endRect, endDirection)


        segments.push(firstSegment, secondSegment)
    }

    if (startDirection === gongsvg.DirectionType.DIRECTION_VERTICAL &&
        endDirection === gongsvg.DirectionType.DIRECTION_HORIZONTAL) {

        const startX = startRect.X + startRatio * startRect.Width
        const startY = startRect.Y 
        const c1_X = startX
        const c1_Y = endRect.Y + endRatio * endRect.Height
        const c1 = createPoint(c1_X, c1_Y)

        const firstSegment = drawLine(c1, startRect, startDirection)
        const secondSegment = drawLine(c1, endRect, endDirection)


        segments.push(firstSegment, secondSegment)
    }

    return segments;
}

function drawLine(point: gongsvg.PointDB, rect: gongsvg.RectDB, direction: gongsvg.DirectionType): gongsvg.PointDB[] {
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