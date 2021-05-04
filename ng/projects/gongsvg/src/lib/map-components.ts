// insertion point sub template for components imports 
import { CirclesTableComponent } from './circles-table/circles-table.component'
import { EllipsesTableComponent } from './ellipses-table/ellipses-table.component'
import { LinesTableComponent } from './lines-table/lines-table.component'
import { PathsTableComponent } from './paths-table/paths-table.component'
import { PolygonesTableComponent } from './polygones-table/polygones-table.component'
import { PolylinesTableComponent } from './polylines-table/polylines-table.component'
import { RectsTableComponent } from './rects-table/rects-table.component'
import { SVGsTableComponent } from './svgs-table/svgs-table.component'
import { TextsTableComponent } from './texts-table/texts-table.component'

// insertion point sub template for map of components per struct 
export const MapOfCirclesComponents: Map<string, any> = new Map([["CirclesTableComponent", CirclesTableComponent],])
export const MapOfEllipsesComponents: Map<string, any> = new Map([["EllipsesTableComponent", EllipsesTableComponent],])
export const MapOfLinesComponents: Map<string, any> = new Map([["LinesTableComponent", LinesTableComponent],])
export const MapOfPathsComponents: Map<string, any> = new Map([["PathsTableComponent", PathsTableComponent],])
export const MapOfPolygonesComponents: Map<string, any> = new Map([["PolygonesTableComponent", PolygonesTableComponent],])
export const MapOfPolylinesComponents: Map<string, any> = new Map([["PolylinesTableComponent", PolylinesTableComponent],])
export const MapOfRectsComponents: Map<string, any> = new Map([["RectsTableComponent", RectsTableComponent],])
export const MapOfSVGsComponents: Map<string, any> = new Map([["SVGsTableComponent", SVGsTableComponent],])
export const MapOfTextsComponents: Map<string, any> = new Map([["TextsTableComponent", TextsTableComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Circle", MapOfCirclesComponents],
      ["Ellipse", MapOfEllipsesComponents],
      ["Line", MapOfLinesComponents],
      ["Path", MapOfPathsComponents],
      ["Polygone", MapOfPolygonesComponents],
      ["Polyline", MapOfPolylinesComponents],
      ["Rect", MapOfRectsComponents],
      ["SVG", MapOfSVGsComponents],
      ["Text", MapOfTextsComponents],
    ]
  )