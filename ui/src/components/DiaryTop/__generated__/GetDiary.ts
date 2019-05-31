

/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: GetDiary
// ====================================================

export interface GetDiary_getDiary_tags {
  tag_name: string;
}

export interface GetDiary_getDiary {
  id: string;
  name: string;
  tags: GetDiary_getDiary_tags[];
  canEdit: boolean;
}

export interface GetDiary {
  getDiary: GetDiary_getDiary;
}

export interface GetDiaryVariables {
  diaryId: string;
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