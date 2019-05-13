import React from "react";
import {RouteComponentProps} from "react-router";
import {Query} from "react-apollo";
import gql from "graphql-tag";

import {DiaryArticleFragment} from "./__generated__/DiaryArticleFragment"
import {GetDiary, GetDiaryVariables} from "./__generated__/GetDiary"

export const diaryArticleFragment = gql`
    fragment DiaryArticleFragment on Article {
        id
        title
        content
    }
`

interface DiaryArticleProps {
    article: DiaryArticleFragment
}

const DiaryArticle: React.StatelessComponent<DiaryArticleProps> = ({ article }) => (
    <div className="DiaryArticle">
        <h2>{article.title}</h2>
        <p>{article.content}</p>
    </div>
);

const diaryFragment = gql`
    fragment DiaryFragment on Diary {
        id
        name
        articles {
            id
            ...DiaryArticleFragment
        }
    }
    ${diaryArticleFragment}
`

export const query = gql`
    query GetDiary($diaryId: ID!) {
        getDiary(diaryId: $diaryId) {
            id
            ...DiaryFragment
        }
    }
    ${diaryFragment}
`

interface RouteProps {
    diaryId: string
}

export const Diary: React.StatelessComponent<RouteComponentProps<RouteProps>> = ({match}) => (
    <div className="Diary">
        <Query<GetDiary, GetDiaryVariables> query={query} variables={{ diaryId: match.params.diaryId}}>
            {result => {
                if (result.error) {
                    return <p className="error">Error: {result.error.message}</p>
                }
                if (result.loading) {
                  return <p className="loading">Loading</p>
                }
                const {data} = result;

                return (
                    <div className="Articles">
                        <h1>{data!.getDiary.name}</h1>
                        <ul>
                            {data!.getDiary.articles.map( article =>
                                <li key={article.id}><DiaryArticle article={article} /></li>)}
                        </ul>
                    </div>
                )
            }}
        </Query>
    </div>
)