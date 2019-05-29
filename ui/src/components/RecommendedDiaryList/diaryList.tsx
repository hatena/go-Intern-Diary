import React from "react"
import { Query } from "react-apollo";

import {query, RecommendedDiary} from "./container"
import {DiaryListItem} from "./diaryListItem"
import {ListRecommededDiaries, ListRecommededDiariesVariables} from "./__generated__/ListRecommededDiaries"

interface RecommendedDiaryListProps {
    diaryId: string;
}

export const RecommendedDiaryList: React.StatelessComponent<RecommendedDiaryListProps> = ({diaryId}) => {
    return (
        <div className="RecommendedDiaryList">
            <Query<ListRecommededDiaries, ListRecommededDiariesVariables>
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
                    const recommendedDiaries: RecommendedDiary[] = data!.listRecommededDiaries.map( diary =>
                        { const d: RecommendedDiary =
                            {
                                id: diary.id,
                                diaryName: diary.name,
                                tags: diary.tags.map(tag => tag.tag_name),
                                userId: diary.user.id,
                                userName: diary.user.name
                            }
                            return d
                        }
                    )
                    return (
                        <div>
                            <h2>おすすめの日記</h2>
                            {recommendedDiaries.length == 0 && <span>特にありません</span>}
                            <ul>
                                {recommendedDiaries.map(diary => (
                                    <li key={diary.id}>
                                        <DiaryListItem diary={diary} />
                                    </li>
                                ))}
                            </ul>
                        </div>
                    )
                }
                }
            </Query>
        </div>
    )
}
