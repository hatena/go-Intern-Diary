

/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL fragment: DiaryListFragment
// ====================================================

export interface DiaryListFragment_diaries {
  id: string;
  name: string;
}

export interface DiaryListFragment {
  name: string;
  diaries: DiaryListFragment_diaries[];
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