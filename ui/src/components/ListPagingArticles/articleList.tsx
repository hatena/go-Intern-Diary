import React from "react"

import { ArticleItem } from "./articleItem";
import {PageInfo, Article} from "./container"



interface ArticleListProps {
    diaryId: string;
    pageInfo: PageInfo;
    articles: Article[];
}

export const ArticleList: React.StatelessComponent<ArticleListProps> = ({diaryId, pageInfo, articles}) => {
    return (
        <div className="ArticleList">
            <ul>
                {articles.map( article => 
                <li key={article.id}>
                    <ArticleItem diaryId={diaryId} pageInfo={pageInfo} article={article}/>
                </li>
                )}
            </ul>
        </div>
    )
}