import React from "react"
import gql from "graphql-tag"
import { MutationUpdaterFn } from "apollo-client";

import {AritlceListWithPagination} from "./articleListWithPagination"
import {DeleteArticle} from "./__generated__/DeleteArticle"
import { ListArticles} from "./__generated__/ListArticles";

export const listArticlesFragment = gql`
    fragment ListArticlesFragment on Article {
        id
        diaryId
        title
        content
    }
`
const pageInfoFragment = gql`
    fragment PageInfoFragment on PageInfo {
        totalPage
        currentPage
        hasNextPage
        hasPreviousPage
    }
`

export const listArticleQuery = gql`
    query ListArticles($diaryId: ID!, $page: Int!) {
        listArticles(diaryId: $diaryId, page: $page) {
            pageInfo {
                ...PageInfoFragment
            }
            articles {
                ...ListArticlesFragment
            }
        }
    }
    ${pageInfoFragment}
    ${listArticlesFragment}
`

export const deleteArticle = gql`
    mutation DeleteArticle($articleId: ID!) {
        deleteArticle(articleId: $articleId)
    }
`

export const deleteUpdateArticle: (diaryId: string, articleId: string, page: number) => MutationUpdaterFn<DeleteArticle> = (diaryId, articleId, page) => (cache, result) => {
    const { data } = result;
    const listArticles = cache.readQuery<ListArticles>({ query: listArticleQuery, variables: {diaryId: diaryId, page: page}})       
    if (listArticles && data) {
        const articles = [...listArticles.listArticles.articles].filter(article => article.id !== articleId);
        const newDiary = {
            listArticles: {
                ...listArticles.listArticles,
                articles: articles,
            }
        };
        cache.writeQuery({ query: listArticleQuery, variables: {diaryId: diaryId, page: page} , data: newDiary });
    };
}

export type Article = {
    id: string,
    title: string,
    content: string
}

export type PageInfo = {
    totalPage: number
    currentPage: number
    hasNextPage: boolean
    hasPreviousPage: boolean
}

interface ListArticlesProps {
    diaryId: string;
    page?: string;
    canEdit: boolean
}

interface ListArticleState {
    page: number;
}

export class ListArticlesContainer extends React.PureComponent<ListArticlesProps, ListArticleState> {
    state = {
        page: 1
    }

    handlePushPageButton = (page: number) => {
        this.setState({
            page: page
        }) 
    }

    componentWillMount() {
        console.log(this.props.page)
        const newPage = this.props.page
        if (newPage == undefined) {
            this.setState({
                page: 1
            })
        } else {
            this.setState({
                page: Number(newPage)
            })
        }
    }

    componentDidUpdate(){
        console.log(this.props.page)
        const newPage = this.props.page
        if (newPage == undefined) {
            this.setState({
                page: 1
            })
        } else {
            this.setState({
                page: Number(newPage)
            })
        }
    }

    render() {
        return (
            <div>
                <AritlceListWithPagination 
                    diaryId={this.props.diaryId}
                    page={this.state.page}
                    handlePushPageButton={this.handlePushPageButton}
                    canEdit={this.props.canEdit}
                />
            </div>
        )
    }

}