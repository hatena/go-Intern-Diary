import React from "react"
import { Query } from "react-apollo";

import { ListArticles, ListArticlesVariables } from "./__generated__/ListArticles";
import { Pagination } from "./pagination"
import {listArticleQuery as query} from "./container"
import {PageInfo, Article} from "./container"
import { ArticleList } from "./articleList";


interface ArticleListWithPaginationProps {
    diaryId: string;
    page: number;
    handlePushPageButton: (page: number) => void;
    canEdit: boolean;
}

export const AritlceListWithPagination: React.StatelessComponent<ArticleListWithPaginationProps> = ({diaryId, page, handlePushPageButton, canEdit}) => {

    return (
        <Query<ListArticles, ListArticlesVariables>
            query={query}
            variables={{
                diaryId: diaryId,
                page: page
            }}
        >
            {result => {
                if (result.error) {
                    return <p className="error">Error: {result.error.message}</p>
                }
                if (result.loading) {
                    return <p className="loading">Loading</p>
                }
                const {data} = result
                const pageInfo: PageInfo = { ...data!.listArticles.pageInfo }
                const articles: Article[] = data!.listArticles.articles.map( article => ({ ...article }))
                
                return (
                    <div>
                        <h2>記事一覧</h2>
                        <ArticleList diaryId={diaryId} pageInfo={pageInfo} articles={articles} canEdit={canEdit}/>
                        <Pagination pageInfo={pageInfo} diaryId={diaryId} handlePushPageButton={handlePushPageButton} />
                    </div>
                )
            }}        
        </Query>
    )
}