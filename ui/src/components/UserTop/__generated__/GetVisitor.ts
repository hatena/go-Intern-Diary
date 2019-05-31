

/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: GetVisitor
// ====================================================

export interface GetVisitor_visitor_diaries {
  id: string;
  name: string;
}

export interface GetVisitor_visitor {
  name: string;
  diaries: GetVisitor_visitor_diaries[];
}

export interface GetVisitor {
  visitor: GetVisitor_visitor;
}

/* tslint:disable */
// This file was automatically generated and should not be edited.

//==============================================================
// START Enums and Input Objects
//==============================================================

/**
 * 
 */
export interface TagWithCategoryInput {
  tag_name: string;
  category_id: number;
}

//==============================================================
// END Enums and Input Objects
//==============================================================