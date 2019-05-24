import React from "react"
import { Query } from "react-apollo";

import {query, RecommendedDiary} from "./container"
import {DiaryListItem} from "./diaryListItem"

interface RecommendedDiaryListProps {
    diaryId: string;
}

export const RecommendedDiaryList: React.StatelessComponent<RecommendedDiaryListProps> = ({diaryId}) => {
    return (
        <div className="RecommendedDiaryList">
            <Query<ListRecommendDiary, ListRecommendDiaryVariables>
                query={query}
                variables={{
                    diaryId: diaryId
                }}
            >
                {result =>{
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
                    const recommendedDiaries = data!.
                    {recommendedDiaries.map(diary => (
                        <div>
                            <DiaryListItem diary={diary} />
                        </div>
                    ))}
                }
                }}
            </Query>
        </div>
    )
}
