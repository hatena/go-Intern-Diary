

/* tslint:disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: ListArticles
// ====================================================

export interface ListArticles_listArticles_pageInfo {
  totalPage: number;
  currentPage: number;
  hasNextPage: boolean;
  hasPreviousPage: boolean;
}

export interface ListArticles_listArticles_articles {
  id: string;
  diaryId: string;
  title: string;
  content: string;
}

export interface ListArticles_listArticles {
  pageInfo: ListArticles_listArticles_pageInfo;
  articles: ListArticles_listArticles_articles[];
}

export interface ListArticles {
  listArticles: ListArticles_listArticles;
}

export interface ListArticlesVariables {
  diaryId: string;
  page: number;
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