import React from "react"

import { Article } from "./container"

interface ArticleItemProps {
    article: Article
    deleteArticle?: (articleId: string) => void
}

export const ArticleItem: React.StatelessComponent<ArticleItemProps> = ({article, deleteArticle}) => {
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