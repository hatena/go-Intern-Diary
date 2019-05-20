import React from "react";
import {Query} from "react-apollo";
import {Link} from "react-router-dom"

import {GetDiary, GetDiaryVariables} from "../__generated__/GetDiary"
import {query} from "./container"

interface RouteProps {
    diaryId: string;
}

export const Diary: React.StatelessComponent<RouteProps> = ({diaryId}) => (
    <div className="Diary">
        <Query<GetDiary, GetDiaryVariables> query={query} variables={{ diaryId: diaryId}}>
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
    </div>
)

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
