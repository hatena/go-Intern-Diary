import React from "react"
import gql from "graphql-tag";
import {RecommendedDiaryList} from "./diaryList"

export const recommendedDiaryFragment = gql`
    fragment RecommendedDiaryFragment on Diary {
        id
        name
        tags {
            tag_name
        }
        user {
            id
            name
        }
    }
`

export const query = gql`
    query ListRecommededDiaries($diaryId: ID!) {
        listRecommededDiaries(diaryId: $diaryId) {
            ...RecommendedDiaryFragment
        }
    }
    ${recommendedDiaryFragment}
`

export type RecommendedDiary = {
    id: string;
    diaryName: string;
    tags: string[];
    userId: string;
    userName: string;
}

interface RecommendedDiaryListContainerProps {
    diaryId: string;
}

export const RecommendedDiaryListContainer: React.StatelessComponent<RecommendedDiaryListContainerProps> = ({diaryId}) => {
    return (
        <div className="RecommendedDiaryList">
            <RecommendedDiaryList diaryId={diaryId}/>
        </div>
    )
}