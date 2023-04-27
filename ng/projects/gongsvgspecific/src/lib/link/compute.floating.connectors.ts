
import * as gongsvg from 'gongsvg'

interface Point {
    X: number;
    Y: number;
  }


export function calculateAnchorPoints(rect: gongsvg.RectDB, ratio: number): Point[] {
  return [
    { X: rect.X + rect.Width * ratio, Y: rect.Y }, // Top
    { X: rect.X + rect.Width, Y: rect.Y + rect.Height * ratio }, // Right
    { X: rect.X + rect.Width * ratio, Y: rect.Y + rect.Height }, // Bottom
    { X: rect.X, Y: rect.Y + rect.Height * ratio }, // Left
  ];
}

function distance(p1: Point, p2: Point): number {
  return Math.abs(p1.X - p2.X) + Math.abs(p1.Y - p2.Y);
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