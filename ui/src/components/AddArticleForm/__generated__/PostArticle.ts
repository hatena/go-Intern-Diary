

/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: PostArticle
// ====================================================

export interface PostArticle_postArticle {
  id: string;
  diaryId: string;
  title: string;
  content: string;
}

export interface PostArticle {
  postArticle: PostArticle_postArticle;
}

export interface PostArticleVariables {
  diaryId: string;
  title: string;
  content: string;
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