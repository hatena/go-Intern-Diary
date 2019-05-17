import React from "react"
import {RouteComponentProps} from "react-router";
import gql from "graphql-tag"

import {AritlceListWithPagination} from "./articleListWithPagination"

const listArticlesFragment = gql`
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
    diaryId: string
    page?: string
}

interface ListArticleState {
    page: number
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
                />
            </div>
        )
    }

}