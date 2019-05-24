import React from "react"
import { Query } from "react-apollo";

import {query} from "./container"

interface RecommendedDiaryListProps {
    diaryId: string;
}


export const RecommendedDiaryList: React.StatelessComponent<RecommendedDiaryListProps> = ({diaryId}) => {
    return (
        <div className="RecommendedDiaryList">
            <Query<>
                query={query}
                variables={{
                    diaryId: diaryId
                }}
            >
            </Query>
        </div>
    )
}
