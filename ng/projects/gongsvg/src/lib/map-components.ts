// insertion point sub template for components imports 
  import { SVGsTableComponent } from './svgs-table/svgs-table.component'
  import { SVGSortingComponent } from './svg-sorting/svg-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfSVGsComponents: Map<string, any> = new Map([["SVGsTableComponent", SVGsTableComponent],])
  export const MapOfSVGSortingComponents: Map<string, any> = new Map([["SVGSortingComponent", SVGSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["SVG", MapOfSVGsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["SVG", MapOfSVGSortingComponents],
    ]
  )
