

/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL fragment: DiaryFragment
// ====================================================

export interface DiaryFragment_tags {
  tag_name: string;
}

export interface DiaryFragment {
  id: string;
  name: string;
  tags: DiaryFragment_tags[];
  canEdit: boolean;
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