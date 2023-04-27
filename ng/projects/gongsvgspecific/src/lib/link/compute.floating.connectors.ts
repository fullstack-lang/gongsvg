
import * as gongsvg from 'gongsvg'

interface Point {
    X: number;
    Y: number;
  }


  interface AnchorPointWithSide {
    point: Point;
    side: gongsvg.SideType;
  }
  
  export function calculateAnchorPoints(rect: gongsvg.RectDB, ratio: number): AnchorPointWithSide[] {
    return [
      { point: { X: rect.X + rect.Width * ratio, Y: rect.Y }, side: gongsvg.SideType.SIDE_TOP }, // Top
      { point: { X: rect.X + rect.Width, Y: rect.Y + rect.Height * ratio }, side: gongsvg.SideType.SIDE_RIGHT }, // Right
      { point: { X: rect.X + rect.Width * ratio, Y: rect.Y + rect.Height }, side: gongsvg.SideType.SIDE_BOTTOM }, // Bottom
      { point: { X: rect.X, Y: rect.Y + rect.Height * ratio }, side: gongsvg.SideType.SIDE_LEFT }, // Left
    ];
  }
function distance(p1: Point, p2: Point): number {
  return Math.abs(p1.X - p2.X) + Math.abs(p1.Y - p2.Y);
}

export function shortestFloatingOrthogonalConnector(rectA: gongsvg.RectDB, ratioA: number, rectB: gongsvg.RectDB, ratioB: number): 
[Point, gongsvg.SideType, Point, gongsvg.SideType] {
  const anchorsA = calculateAnchorPoints(rectA, ratioA);
  const anchorsB = calculateAnchorPoints(rectB, ratioB);

  let minDistance = Infinity;
  let shortestConnection: [Point, gongsvg.SideType, Point, gongsvg.SideType] = 
  [anchorsA[0].point, anchorsA[0].side, anchorsB[0].point, anchorsB[0].side];

  let startHorizonalLink = false
  let canGoHorizontalyStraight = false
  let canGoVerticalyStraight = false

  for (const anchorA of anchorsA) {
    for (const anchorB of anchorsB) {
      const d = distance(anchorA.point, anchorB.point);

      if (d < minDistance) {
        minDistance = d;
        shortestConnection = [anchorA.point, anchorA.side, anchorB.point, anchorB.side];

        startHorizonalLink = 
        (anchorA.side === gongsvg.SideType.SIDE_LEFT && anchorB.side === gongsvg.SideType.SIDE_RIGHT) ||
        (anchorA.side === gongsvg.SideType.SIDE_RIGHT && anchorB.side === gongsvg.SideType.SIDE_LEFT)


        canGoHorizontalyStraight = ( startHorizonalLink && anchorA.point.Y >= rectB.Y && anchorA.point.Y <= rectB.Y+rectB.Height ) ||
        (startHorizonalLink && anchorA.point.Y >= rectB.Y && anchorA.point.Y <= rectB.Y+rectB.Height)

        if (canGoHorizontalyStraight) {
          anchorB.point.Y = anchorA.point.Y
        } 
        else {
          anchorB.point.Y = anchorA.point.Y
          anchorB.point.X = rectB.X + ratioB * rectB.Width
        }

        canGoVerticalyStraight = 
        (anchorA.side === gongsvg.SideType.SIDE_TOP && anchorB.side === gongsvg.SideType.SIDE_BOTTOM && anchorA.point.X >= rectB.X && anchorA.point.X <= rectB.X+rectB.Width) ||
        (anchorA.side === gongsvg.SideType.SIDE_BOTTOM && anchorB.side === gongsvg.SideType.SIDE_TOP && anchorA.point.X >= rectB.X && anchorA.point.X <= rectB.X+rectB.Width);

        if (canGoVerticalyStraight) {
          anchorB.point.X = anchorA.point.X
        }
      }
    }
  }


  console.log( "can go straight horizonaly", canGoHorizontalyStraight, "can go straight verticaly ", canGoVerticalyStraight)
  return shortestConnection;
}