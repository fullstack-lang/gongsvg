Repo archived, function is no as a lib in gong

# gongsvg

*gongsvg* stack provides a way to render basic SVG elements and also a more abstract level of representation for drawing
rects and links between rects.

A typical gongsvg data set is made of one "svg" element that is contrued with layers. Those layers

Each layer can hold any kind of SVG elements (rect, link, line, ...). The purpose of layers is to allow for specification
of display order and wether to displaya layer or not.

The *gongsvg* stack goal is to be reused for different use cases. One use case is the drawing of boxes and arrows between boxes for architecting models. Therefore, the stack allows the user to move/resize boxes (called "rect" in SVG term) and to link (a 
set of "line" in SVG term) box between them and to move the links between boxes.





