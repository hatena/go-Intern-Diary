import React from "react";
import {RouteComponentProps} from "react-router";
import {Query, Mutation} from "react-apollo";
import gql from "graphql-tag";
import {Link} from "react-router-dom"
import { MutationUpdaterFn } from "apollo-client";

import {DiaryArticleFragment} from "./__generated__/DiaryArticleFragment"
import {GetDiary, GetDiaryVariables} from "./__generated__/GetDiary"
import {DeleteArticle, DeleteArticleVariables} from "./__generated__/DeleteArticle"
import {ListArticlesContainer} from "./ListPagingArticles/container"
import {ListArticles} from "./ListPagingArticles/__generated__/ListArticles"
import {listArticleQuery} from "./ListPagingArticles/container";

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
    // window.location.reload();
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
    page?: string
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

// export const Diary: React.StatelessComponent<RouteComponentProps<RouteProps>> = ({match}) => (
//     <div className="Diary">
//         <Query<GetDiary, GetDiaryVariables> query={query} variables={{ diaryId: match.params.diaryId}}>
//             {result => {
//                 if (result.error) {
//                     return <p className="error">Error: {result.error.message}</p>
//                 }
//                 if (result.loading) {
//                     return <p className="loading">Loading</p>
//                 }
//                 const {data} = result;

//                 return (        
//                     <Mutation<DeleteArticle, DeleteArticleVariables> mutation={deleteArticle}>
//                     {(deleteArticle) => {
//                         return (
//                             <div className="Articles">
//                                 <h1>{data!.getDiary.name}</h1>
//                                 <ToAddArtilce diaryId={data!.getDiary.id} />
//                                 <ul>
//                                     {data!.getDiary.articles.map( article =>
//                                         <li key={article.id}><DiaryArticle article={article} deleteArticle={(articleId: string) =>
//                                             deleteArticle({ variables: {articleId}, update: deleteUpdateArticle(data!.getDiary.id, articleId)})} /></li>)}
//                                 </ul>
//                             </div>
//                         )
//                     }}
//                     </Mutation> 
//                 )
//             }}
//         </Query>
//     </div>
// )

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

                return ( <div>
                            <h1>{data!.getDiary.name}</h1>
                            <ToAddArtilce diaryId={data!.getDiary.id} />
                        </div>
                )
            }}
        </Query>
        <ListArticlesContainer diaryId={match.params.diaryId} page={match.params.page} />
    </div>
)