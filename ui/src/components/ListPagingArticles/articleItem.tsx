import React from "react"
import { Mutation } from "react-apollo";

import {DeleteArticle, DeleteArticleVariables} from "../__generated__/DeleteArticle"
import { Article, PageInfo, deleteArticle, deleteUpdateArticle } from "./container"

interface ArticleItemProps {
    diaryId: string;
    pageInfo: PageInfo;
    article: Article;
}

export const ArticleItem: React.StatelessComponent<ArticleItemProps> = ({diaryId, pageInfo, article}) => (
    <Mutation<DeleteArticle, DeleteArticleVariables> mutation={deleteArticle}>
    {(deleteArticle) => (
        <ArticleItemPresentation
            article={article} 
            deleteArticle={(articleId: string) => deleteArticle({ variables: {articleId}, update: deleteUpdateArticle(diaryId, articleId, pageInfo.currentPage) })}
        />
    )}
    </Mutation>
)

interface ArticleItemPresentationProps {
    article: Article
    deleteArticle?: (articleId: string) => void
}
const ArticleItemPresentation: React.StatelessComponent<ArticleItemPresentationProps> = ({article, deleteArticle}) => {
    return (
        <div className="DiaryArticle">
            <h2>{article.title}</h2>
            <p>{article.content}</p>
            <div>
                {deleteArticle && <button onClick={deleteArticle ? () => {deleteArticle(article.id); }: undefined}>Delete</button>}
            </div>
        </div>  
    )
}