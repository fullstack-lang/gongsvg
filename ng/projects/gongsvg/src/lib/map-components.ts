// insertion point sub template for components imports 
import { RectsTableComponent } from './rects-table/rects-table.component'
import { SVGsTableComponent } from './svgs-table/svgs-table.component'
import { TextsTableComponent } from './texts-table/texts-table.component'

// insertion point sub template for map of components per struct 
export const MapOfRectsComponents: Map<string, any> = new Map([["RectsTableComponent", RectsTableComponent],])
export const MapOfSVGsComponents: Map<string, any> = new Map([["SVGsTableComponent", SVGsTableComponent],])
export const MapOfTextsComponents: Map<string, any> = new Map([["TextsTableComponent", TextsTableComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Rect", MapOfRectsComponents],
      ["SVG", MapOfSVGsComponents],
      ["Text", MapOfTextsComponents],
    ]
  )