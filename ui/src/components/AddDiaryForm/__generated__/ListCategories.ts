

/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: ListCategories
// ====================================================

export interface ListCategories_listCategories {
  id: number;
  category_name: string;
}

export interface ListCategories {
  listCategories: ListCategories_listCategories[];
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