import React from "react"

import {RecommendedDiary} from "./container"

interface DiaryListItemProps {
    diary: RecommendedDiary;
}

export const DiaryListItem: React.StatelessComponent<DiaryListItemProps> = ({diary}) => (
    <div key={diary.id}>
        <h2>{diary.diaryName}</h2><span>{diary.userName}</span>
    </div>
)