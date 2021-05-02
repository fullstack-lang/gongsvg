import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbService } from '../commitnb.service'

// insertion point for per struct import code
import { CircleService } from '../circle.service'
import { getCircleUniqueID } from '../front-repo.service'
import { LineService } from '../line.service'
import { getLineUniqueID } from '../front-repo.service'
import { RectService } from '../rect.service'
import { getRectUniqueID } from '../front-repo.service'
import { SVGService } from '../svg.service'
import { getSVGUniqueID } from '../front-repo.service'
import { TextService } from '../text.service'
import { getTextUniqueID } from '../front-repo.service'

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ONE__ZERO_ONE_ASSOCIATION = 'ONE__ZERO_ONE_ASSOCIATION',
  ONE__ZERO_MANY_ASSOCIATION = 'ONE__ZERO_MANY_ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children?: GongNode[];
  type: GongNodeType;
  structName: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
interface GongFlatNode {
  expandable: boolean;
  name: string;
  level: number;
  type: GongNodeType;
  structName: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-gongsvg-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
})
export class SidebarComponent implements OnInit {

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      associatedStructName: node.associatedStructName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  frontRepo: FrontRepo
  commitNb: number

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private commitNbService: CommitNbService,

    // insertion point for per struct service declaration
    private circleService: CircleService,
    private lineService: LineService,
    private rectService: RectService,
    private svgService: SVGService,
    private textService: TextService,
  ) { }

  ngOnInit(): void {
    this.refresh()

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.circleService.CircleServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.lineService.LineServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.rectService.RectServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.svgService.SVGServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.textService.TextServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull().subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      if (this.treeControl.dataNodes != undefined) {
        this.treeControl.dataNodes.forEach(
          node => {
            if (this.treeControl.isExpanded(node)) {
              memoryOfExpandedNodes[node.uniqueIdPerStack] = true
            } else {
              memoryOfExpandedNodes[node.uniqueIdPerStack] = false
            }
          }
        )
      }

      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the Circle part of the mat tree
      */
      let circleGongNodeStruct: GongNode = {
        name: "Circle",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Circle",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(circleGongNodeStruct)

      this.frontRepo.Circles_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Circles_array.forEach(
        circleDB => {
          let circleGongNodeInstance: GongNode = {
            name: circleDB.Name,
            type: GongNodeType.INSTANCE,
            id: circleDB.ID,
            uniqueIdPerStack: getCircleUniqueID(circleDB.ID),
            structName: "Circle",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          circleGongNodeStruct.children.push(circleGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Line part of the mat tree
      */
      let lineGongNodeStruct: GongNode = {
        name: "Line",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Line",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(lineGongNodeStruct)

      this.frontRepo.Lines_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Lines_array.forEach(
        lineDB => {
          let lineGongNodeInstance: GongNode = {
            name: lineDB.Name,
            type: GongNodeType.INSTANCE,
            id: lineDB.ID,
            uniqueIdPerStack: getLineUniqueID(lineDB.ID),
            structName: "Line",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          lineGongNodeStruct.children.push(lineGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Rect part of the mat tree
      */
      let rectGongNodeStruct: GongNode = {
        name: "Rect",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Rect",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(rectGongNodeStruct)

      this.frontRepo.Rects_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Rects_array.forEach(
        rectDB => {
          let rectGongNodeInstance: GongNode = {
            name: rectDB.Name,
            type: GongNodeType.INSTANCE,
            id: rectDB.ID,
            uniqueIdPerStack: getRectUniqueID(rectDB.ID),
            structName: "Rect",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          rectGongNodeStruct.children.push(rectGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the SVG part of the mat tree
      */
      let svgGongNodeStruct: GongNode = {
        name: "SVG",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "SVG",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(svgGongNodeStruct)

      this.frontRepo.SVGs_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.SVGs_array.forEach(
        svgDB => {
          let svgGongNodeInstance: GongNode = {
            name: svgDB.Name,
            type: GongNodeType.INSTANCE,
            id: svgDB.ID,
            uniqueIdPerStack: getSVGUniqueID(svgDB.ID),
            structName: "SVG",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          svgGongNodeStruct.children.push(svgGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Rects
          */
          let RectsGongNodeAssociation: GongNode = {
            name: "(Rect) Rects",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: svgDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "SVG",
            associatedStructName: "Rect",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          svgGongNodeInstance.children.push(RectsGongNodeAssociation)

          svgDB.Rects?.forEach(rectDB => {
            let rectNode: GongNode = {
              name: rectDB.Name,
              type: GongNodeType.INSTANCE,
              id: rectDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getSVGUniqueID(svgDB.ID)
                + 11 * getRectUniqueID(rectDB.ID),
              structName: "Rect",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            RectsGongNodeAssociation.children.push(rectNode)
          })

          /**
          * let append a node for the slide of pointer Texts
          */
          let TextsGongNodeAssociation: GongNode = {
            name: "(Text) Texts",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: svgDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "SVG",
            associatedStructName: "Text",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          svgGongNodeInstance.children.push(TextsGongNodeAssociation)

          svgDB.Texts?.forEach(textDB => {
            let textNode: GongNode = {
              name: textDB.Name,
              type: GongNodeType.INSTANCE,
              id: textDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getSVGUniqueID(svgDB.ID)
                + 11 * getTextUniqueID(textDB.ID),
              structName: "Text",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            TextsGongNodeAssociation.children.push(textNode)
          })

          /**
          * let append a node for the slide of pointer Circles
          */
          let CirclesGongNodeAssociation: GongNode = {
            name: "(Circle) Circles",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: svgDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "SVG",
            associatedStructName: "Circle",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          svgGongNodeInstance.children.push(CirclesGongNodeAssociation)

          svgDB.Circles?.forEach(circleDB => {
            let circleNode: GongNode = {
              name: circleDB.Name,
              type: GongNodeType.INSTANCE,
              id: circleDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getSVGUniqueID(svgDB.ID)
                + 11 * getCircleUniqueID(circleDB.ID),
              structName: "Circle",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            CirclesGongNodeAssociation.children.push(circleNode)
          })

          /**
          * let append a node for the slide of pointer Lines
          */
          let LinesGongNodeAssociation: GongNode = {
            name: "(Line) Lines",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: svgDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "SVG",
            associatedStructName: "Line",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          svgGongNodeInstance.children.push(LinesGongNodeAssociation)

          svgDB.Lines?.forEach(lineDB => {
            let lineNode: GongNode = {
              name: lineDB.Name,
              type: GongNodeType.INSTANCE,
              id: lineDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getSVGUniqueID(svgDB.ID)
                + 11 * getLineUniqueID(lineDB.ID),
              structName: "Line",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            LinesGongNodeAssociation.children.push(lineNode)
          })

        }
      )

      /**
      * fill up the Text part of the mat tree
      */
      let textGongNodeStruct: GongNode = {
        name: "Text",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Text",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(textGongNodeStruct)

      this.frontRepo.Texts_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Texts_array.forEach(
        textDB => {
          let textGongNodeInstance: GongNode = {
            name: textDB.Name,
            type: GongNodeType.INSTANCE,
            id: textDB.ID,
            uniqueIdPerStack: getTextUniqueID(textDB.ID),
            structName: "Text",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          textGongNodeStruct.children.push(textGongNodeInstance)

          // insertion point for per field code
        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      if (this.treeControl.dataNodes != undefined) {
        this.treeControl.dataNodes.forEach(
          node => {
            if (memoryOfExpandedNodes[node.uniqueIdPerStack] != undefined) {
              if (memoryOfExpandedNodes[node.uniqueIdPerStack]) {
                this.treeControl.expand(node)
              }
            }
          }
        )
      }
    });

    // fetch the number of commits
    this.commitNbService.getCommitNb().subscribe(
      commitNb => {
        this.commitNb = commitNb
      }
    )
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        table: [path]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      this.router.navigate([{
        outlets: {
          table: [path.toLowerCase()]
        }
      }]);
    }

    if (type == GongNodeType.INSTANCE) {
      this.router.navigate([{
        outlets: {
          presentation: [structName.toLowerCase() + "-presentation", id]
        }
      }]);
    }
  }

  setEditorRouterOutlet(path) {
    this.router.navigate([{
      outlets: {
        editor: [path.toLowerCase()]
      }
    }]);
  }

  setEditorSpecialRouterOutlet( node: GongFlatNode) {
    console.log("setEditorSpecialRouterOutlet " + node)
    this.router.navigate([{
      outlets: {
        editor: [node.associatedStructName.toLowerCase() + "-adder", node.id, node.structName + "_" + node.name]
      }
    }]);
  }
}
