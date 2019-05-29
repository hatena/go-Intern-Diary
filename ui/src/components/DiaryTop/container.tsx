import React from "react";
import {RouteComponentProps} from "react-router";
import gql from "graphql-tag";

import {Diary} from "./diary"

export const diaryFragment = gql`
    fragment DiaryFragment on Diary {
        id
        name
        tags {
            tag_name
        }
        canEdit
    }
`

export const query = gql`
    query GetDiary($diaryId: ID!) {
        getDiary(diaryId: $diaryId) {
            ...DiaryFragment
        }
    }
    ${diaryFragment}
`

export type Tag = {
    name: string
}

interface RouteProps {
    diaryId: string;
    page?: string
}

export const DiaryTopContainer: React.StatelessComponent<RouteComponentProps<RouteProps>> = ({match}) => (
    <div className="DiaryTop">
        <Diary diaryId={match.params.diaryId} page={match.params.page}/>
    </div>
)