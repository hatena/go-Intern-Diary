import React from "react"
import { Query } from "react-apollo";

import { ListArticles, ListArticlesVariables } from "./__generated__/ListArticles";
import { Pagination } from "./pagination"
import {listArticleQuery as query} from "./container"
import {PageInfo, Article} from "./container"
import { ArticleList } from "./articleList";


interface ArticleListWithPaginationProps {
    diaryId: string
    page: number
    handlePushPageButton: (page: number) => void
}

export const AritlceListWithPagination: React.StatelessComponent<ArticleListWithPaginationProps> = ({diaryId, page, handlePushPageButton}) => {

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
                        <ArticleList diaryId={diaryId} pageInfo={pageInfo} articles={articles} />
                        <Pagination pageInfo={pageInfo} diaryId={diaryId} handlePushPageButton={handlePushPageButton} />
                    </div>
                )
            }}        
        </Query>
    )
}