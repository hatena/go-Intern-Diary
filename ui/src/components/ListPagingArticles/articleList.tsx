import React from "react"

import { ArticleItem } from "./articleItem";
import {PageInfo, Article} from "./container"



interface ArticleListProps {
    diaryId: string;
    pageInfo: PageInfo;
    articles: Article[];
    canEdit: boolean;
}

export const ArticleList: React.StatelessComponent<ArticleListProps> = ({diaryId, pageInfo, articles, canEdit}) => {
    return (
        <div className="ArticleList">
            <ul>
                {articles.map( article => 
                <li key={article.id}>
                    <ArticleItem diaryId={diaryId} pageInfo={pageInfo} article={article} canEdit={canEdit}/>
                </li>
                )}
            </ul>
        </div>
    )
}