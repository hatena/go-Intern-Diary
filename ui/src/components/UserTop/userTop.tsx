import React from "react"
import {Query} from "react-apollo"

import {DiaryList} from "../DiaryList/diaryList"
import{ GetVisitor } from "./__generated__/GetVisitor"
import { CreateDiaryFormContainer} from "../AddDiaryForm/container"
import { query } from "./container"


export const UserTop: React.StatelessComponent = () => (
    <div className="UserTop">
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