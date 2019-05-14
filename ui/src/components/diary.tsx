import React from "react";
import {RouteComponentProps} from "react-router";
import {Query, Mutation} from "react-apollo";
import gql from "graphql-tag";
import {Link} from "react-router-dom"
import { MutationUpdaterFn } from "apollo-client";

import {DiaryArticleFragment} from "./__generated__/DiaryArticleFragment"
import {GetDiary, GetDiaryVariables} from "./__generated__/GetDiary"
import {DeleteArticle, DeleteArticleVariables} from "./__generated__/DeleteArticle"

export const diaryArticleFragment = gql`
    fragment DiaryArticleFragment on Article {
        id
        diaryId
        title
        content
    }
`

export const diaryFragment = gql`
    fragment DiaryFragment on Diary {
        id
        name
        articles {
            ...DiaryArticleFragment
        }
    }
    ${diaryArticleFragment}
`

export const query = gql`
    query GetDiary($diaryId: ID!) {
        getDiary(diaryId: $diaryId) {
            ...DiaryFragment
        }
    }
    ${diaryFragment}
`

const deleteArticle = gql`
    mutation DeleteArticle($articleId: ID!) {
        deleteArticle(articleId: $articleId)
    }
`

const deleteUpdateArticle: (diaryId: string, articleId: string) => MutationUpdaterFn<DeleteArticle> = (diaryId, articleId) => (cache, result) => {
    const { data } = result
    const diary = cache.readQuery<GetDiary>({ query: query, variables: {diaryId: diaryId}}) 
    if (diary && data) {
        const name = diary.getDiary.name
        const articles = [...diary.getDiary.articles].filter(article => article.id !== articleId);
        // ここがよくわからない
        const newDiary = {
            getDiary: {
                // id: diaryId, name: name, // なぜこれだとダメなのかわからない、このスプレッド記法がなんのためにあるのか
                ...diary.getDiary, 
                articles: articles,
            }
        };
        cache.writeQuery({query, data: newDiary})
    }
}

interface DiaryArticleProps {
    article: DiaryArticleFragment;
    deleteArticle?: (articleId: string) => void;
}

const DiaryArticle: React.StatelessComponent<DiaryArticleProps> = ({ article, deleteArticle }) => (
    <div className="DiaryArticle">
        <h2>{article.title}</h2>
        <p>{article.content}</p>
        <div>
            {deleteArticle && <button onClick={deleteArticle ? () => {deleteArticle(article.id); }: undefined}>Delete</button>}
        </div>
    </div>
);

interface RouteProps {
    diaryId: string
}

interface ToAddArticlePorps {
    diaryId: string
}

const ToAddArtilce: React.StatelessComponent<ToAddArticlePorps> = ({diaryId}) => {
    return (
        <div>
            <Link to={`/diaries/${diaryId}/add`}>新規記事をポスト</Link>
        </div>
    )
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
                    <Mutation<DeleteArticle, DeleteArticleVariables> mutation={deleteArticle}>
                    {(deleteArticle) => {
                        return (
                            <div className="Articles">
                                <h1>{data!.getDiary.name}</h1>
                                <ToAddArtilce diaryId={data!.getDiary.id} />
                                <ul>
                                    {data!.getDiary.articles.map( article =>
                                        <li key={article.id}><DiaryArticle article={article} deleteArticle={(articleId: string) =>
                                            deleteArticle({ variables: {articleId}, update: deleteUpdateArticle(data!.getDiary.id, articleId)})} /></li>)}
                                </ul>
                            </div>
                        )
                    }}
                    </Mutation> 
                )
            }}
        </Query>
    </div>
)