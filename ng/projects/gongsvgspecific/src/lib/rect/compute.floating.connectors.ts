
import * as gongsvg from 'gongsvg'

interface Point {
    x: number;
    y: number;
  }


export function calculateAnchorPoints(rect: gongsvg.RectDB, ratio: number): Point[] {
  return [
    { x: rect.X + rect.Width * ratio, y: rect.Y }, // Top
    { x: rect.X + rect.Width, y: rect.Y + rect.Height * ratio }, // Right
    { x: rect.X + rect.Width * ratio, y: rect.Y + rect.Height }, // Bottom
    { x: rect.X, y: rect.Y + rect.Height * ratio }, // Left
  ];
}

function distance(p1: Point, p2: Point): number {
  return Math.abs(p1.x - p2.x) + Math.abs(p1.y - p2.y);
}

export function shortestFloatingOrthogonalConnector(rectA: gongsvg.RectDB, ratioA: number, rectB: gongsvg.RectDB, ratioB: number): [Point, Point] {
  const anchorsA = calculateAnchorPoints(rectA, ratioA);
  const anchorsB = calculateAnchorPoints(rectB, ratioB);

  let minDistance = Infinity;
  let shortestConnection: [Point, Point] = [anchorsA[0], anchorsB[0]];

  for (const anchorA of anchorsA) {
    for (const anchorB of anchorsB) {
      const d = distance(anchorA, anchorB);

      if (d < minDistance) {
        minDistance = d;
        shortestConnection = [anchorA, anchorB];
      }
    }
  }

  return shortestConnection;
}