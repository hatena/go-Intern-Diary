

/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: ListRecommededDiaries
// ====================================================

export interface ListRecommededDiaries_listRecommededDiaries_tags {
  tag_name: string;
}

export interface ListRecommededDiaries_listRecommededDiaries_user {
  id: string;
  name: string;
}

export interface ListRecommededDiaries_listRecommededDiaries {
  id: string;
  name: string;
  tags: ListRecommededDiaries_listRecommededDiaries_tags[];
  user: ListRecommededDiaries_listRecommededDiaries_user;
}

export interface ListRecommededDiaries {
  listRecommededDiaries: ListRecommededDiaries_listRecommededDiaries[];
}

export interface ListRecommededDiariesVariables {
  diaryId: string;
}

/* tslint:disable */
// This file was automatically generated and should not be edited.

//==============================================================
// START Enums and Input Objects
//==============================================================

//==============================================================
// END Enums and Input Objects
//==============================================================