import React from "react"
import {Query} from "react-apollo"
import gql from "graphql-tag"

import {DiaryList, diaryListFragment} from "./diaries"

import{ GetVisitor } from "./__generated__/GetVisitor"
import { CreateDiaryFormContainer, mutation as createDiary } from "./AddDiaryForm/container"

export const query = gql`
    query GetVisitor {
        visitor {
            ...DiaryListFragment
        }
    }
${diaryListFragment}
`

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
                    <CreateDiaryFormContainer />
                    <DiaryList user={data!.visitor} />
                </>;
                }}
        </Query>
    </div>
);