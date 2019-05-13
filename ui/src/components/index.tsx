import React from "react"
import {Query} from "react-apollo"
import gql from "graphql-tag"

import {DiaryList, diaryListFragment} from "./diaries"
import {UserType, User} from "./user"

interface Visitor {
    visitor: UserType
}

interface ListDiaries_listDiaries {
    id: string;
    name: string;
}
interface ListDiaries {
    user: Visitor;
    listDiaries: ListDiaries_listDiaries[];
}

const query = gql`
    query Visitor {
        visitor {
            id, name
        }
    }
`

export const Index: React.StatelessComponent = () => (
    <div className="Index">
        <h1>Diaries</h1>
        <Query<Visitor> query={query}>
            {result => {
                if (result.error) {
                    return <p className="error">Error: {result.error.message}</p>                    
                }
                if (result.loading) {
                    return <p className="loading">Loading</p>
                }
                const { data } = result;
                if (data == undefined) {
                    return <p>ログインして下さい</p>
                }
                const user = { 
                    id: data.visitor.id,
                    name: data.visitor.name
                 }
                return <User user={user} />;
            }}
        </Query>
    </div>
);