import React from "react"
import {Query, Mutation} from "react-apollo"
import gql from "graphql-tag"

import {DiaryList, diaryListFragment} from "./diaries"

import{ GetVisitor } from "./__generated__/GetVisitor"
import { DeleteDiary, DeleteDiaryVariables } from "./__generated__/DeleteDiary"
import { MutationUpdaterFn } from "apollo-client";

import { createUpdateDiary, mutation as createDiary, CreateDiaryForm } from "./addDiary"
import { CreateDiary, CreateDiaryVariables } from "./__generated__/CreateDiary"

export const query = gql`
    query GetVisitor {
        visitor {
            ...DiaryListFragment
        }
    }
${diaryListFragment}
`

export const deleteDiary = gql`
    mutation DeleteDiary($diaryId: ID!) {
        deleteDiary(diaryId: $diaryId)
    }
`;

// キャッシュ？の動きがイメージしづらい
const updateDiary: (diaryId: string) => MutationUpdaterFn<DeleteDiary> = (diaryId) => (cache, result) => {
    const visitor  = cache.readQuery<GetVisitor>({ query });
    const { data } = result
    if (visitor && data) {
        const diaries = [...visitor.visitor.diaries].filter(diary => diary.id !== diaryId);
        const newVisitor = {
            visitor: {
                ...visitor.visitor,
                diaries,
            }
        };
        cache.writeQuery({query, data: newVisitor})
    }
}

export const Index: React.StatelessComponent = () => (
    <div className="Index">
        <h1>Diaries</h1>
        <Query<GetVisitor> query={query}>
            {result => {
                if (result.error) {
                    if (result.error.message == "GraphQL error: please login") {
                        return <a href="/">ログインして下さい</a>
                    }
                    return <p className="error">Error: {result.error.message}</p>                    
                }
                if (result.loading) {
                    return <p className="loading">Loading</p>
                }
                const { data } = result;
                return <>
                    <Mutation<CreateDiary, CreateDiaryVariables> mutation={createDiary}>
                        {(createDiary) => {
                            return <CreateDiaryForm 
                                    create={(name: string) =>
                                    createDiary({ variables: {name}, update: createUpdateDiary})} />
                        }}
                    </Mutation>
                    <Mutation<DeleteDiary, DeleteDiaryVariables> mutation={deleteDiary}>
                        {(deleteDiary) => {
                            return <DiaryList 
                                user={data!.visitor}
                                deleteDiary={(diaryId: string) =>
                                deleteDiary({ variables: {diaryId}, update: updateDiary(diaryId)})} />
                        }}
                    </Mutation>
                </>;
                }}
        </Query>
    </div>
);