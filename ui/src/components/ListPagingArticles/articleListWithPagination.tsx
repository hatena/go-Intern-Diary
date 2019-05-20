import gql from "graphql-tag"
import React from "react"
import {RouteComponentProps} from "react-router";
import { Query, Mutation } from "react-apollo";

import {DeleteArticle, DeleteArticleVariables} from "../__generated__/DeleteArticle"
import {deleteArticle, deleteUpdateArticle} from "../diary"
import { ListArticles, ListArticlesVariables } from "./__generated__/ListArticles";
import { ArticleItem } from "./articleItem";
import { Pagination } from "./pagination"

import {PageInfo, Article, listArticleQuery as query } from "./container"

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
                        <div className="ArticleList">
                            <ul>
                                {articles.map( article => 
                                <li key={article.id}>
                                    <Mutation<DeleteArticle, DeleteArticleVariables> mutation={deleteArticle}>
                                    {(deleteArticle) => (
                                        <ArticleItem 
                                            article={article} 
                                            deleteArticle={(articleId: string) => deleteArticle({ variables: {articleId}, update: deleteUpdateArticle(diaryId, articleId, pageInfo.currentPage) })}
                                        />
                                    )}
                                    </Mutation>
                                </li>
                                )}
                            </ul>
                        </div>
                        <div className="Pagination">
                            <Pagination pageInfo={pageInfo} diaryId={diaryId} handlePushPageButton={handlePushPageButton} />
                        </div>
                    </div>
                )
            }}        
        </Query>
    )
}