import gql from "graphql-tag"
import React from "react"


const listArticlesFragment = gql`
    fragment ListArticlesFragment {
        id
        diaryId
        title
        content
    }
`

const pageInfoFragment = gql `
    fragment PageInfo {
        totalPage
        currentPage
        hasNextPage
        hasPreviousPage
    }
`

const query = gql`
    query ListArticles($diaryId: ID!) {
        listArticles(diaryId: $diaryId) {
            ...PageInfoFragment
            ...ListArticlesFragment
        }
    }
    ${pageInfoFragment}
    ${listArticlesFragment}
`

type Articles = {
    id: string,
    title: string,
    content: string
}

interface ListArticlesProps {

}

interface ListArticleState {

}

export class ListArticlesContainer extends React.PureComponent<ListArticlesProps, ListArticleState> {
    state = {
        
    }

}