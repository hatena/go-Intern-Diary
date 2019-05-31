

/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL fragment: RecommendedDiaryFragment
// ====================================================

export interface RecommendedDiaryFragment_tags {
  tag_name: string;
}

export interface RecommendedDiaryFragment_user {
  id: string;
  name: string;
}

export interface RecommendedDiaryFragment {
  id: string;
  name: string;
  tags: RecommendedDiaryFragment_tags[];
  user: RecommendedDiaryFragment_user;
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